package cache

import (
	"fmt"
	"strings"

	"github.com/Kindred87/Spoke/internal/regex"
	"golang.org/x/sync/errgroup"
)

const (
	conditionTag = "if_condition ="
)

// BuildFrom caches the given tree.
func BuildFrom(tree map[string]any) error {
	if tree == nil {
		return fmt.Errorf("received nil tree")
	}

	if db == nil {
		err := createDB()
		if err != nil {
			return fmt.Errorf("error while creating cache: %w", err)
		}
	}

	err := cacheScenes(tree)
	if err != nil {
		return fmt.Errorf("error while caching scenes: %w", err)
	}

	return nil
}

// cacheScenes caches the scenes in the given tree.
func cacheScenes(tree map[string]any) error {
	var eg errgroup.Group

	for key := range tree {
		k := key
		eg.Go(func() error { return cacheComponents(tree[k], []string{bucketScenes, k}, &eg) })
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

// cacheComponents recursively caches the given component and all of its sub-components.
func cacheComponents(component any, path []string, eg *errgroup.Group) error {
	switch comp := component.(type) {
	case string:
		// If text matches the branch's name
		if comp == path[len(path)-1] {
			return nil
		}
		p := path
		eg.Go(func() error { return db.InsertValue(comp, p) })
	case map[string]any:
		for key, value := range comp {

			// If we have a choice, we want to assign the first sub-component as its identifier.
			if key == "choose" {
				k, err := parseChoiceKey(key, value)
				if err != nil {
					return fmt.Errorf("error while parsing choice in %v: %w", removeFrom(bucketScenes, path), err)
				}
				key = k
			}

			switch value.(type) {
			case nil:
				k, p := key, path
				eg.Go(func() error { return db.InsertValue("goto "+k, p) })
				continue
			}

			k, v, p := key, value, path
			eg.Go(func() error { return cacheComponents(v, append(p, k), eg) })
		}
	case []any:
		for _, value := range comp {
			v, p := value, path
			eg.Go(func() error { return cacheComponents(v, p, eg) })
		}
	default:
		return fmt.Errorf("value %v of type %T within %v is unsupported", comp, comp, removeFrom(bucketScenes, path))
	}
	return nil
}

func parseChoiceKey(key string, value any) (string, error) {
	switch t := value.(type) {
	case nil:
		return "", fmt.Errorf("empty choice branch detected")
	case string:
		key = t
	case []any:
		first, ok := t[0].(string)
		if !ok {
			return "", fmt.Errorf("expected value %v of type %T to be string", t[0], t[0])
		}

		if len(t) <= 1 {
			return "", fmt.Errorf("empty choice branch detected at \"%s\"", first)
		} else if strings.Contains(first, ":=") {
			return "", fmt.Errorf("assignment \"%s\" detected at beginning of choice branch", first)
		}

		if regex.IsCondition(first) {
			second, ok := t[1].(string)
			if !ok {
				return "", fmt.Errorf("expected value %v of type %T to be string", t[1], t[1])
			} else if regex.IsCondition(second) {
				return "", fmt.Errorf("extra condition entry \"%s\" detected", second)
			} else if strings.Contains(second, ":=") {
				return "", fmt.Errorf("assignment \"%s\" detected after condition \"%s\"", second, first)
			}
			key = fmt.Sprintf("%s %s %s", second, conditionTag, first)
		} else {
			key = first
		}
	}

	return key, nil
}
