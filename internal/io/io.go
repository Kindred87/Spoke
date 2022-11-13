package io

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Load unmarshals the given YAML file into a map.
func Load(file string) (map[string]any, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error while loading %s: %w", file, err)
	}

	tree := make(map[string]any)

	err = yaml.Unmarshal(b, &tree)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling %s: %w", file, err)
	}

	if len(tree) == 0 {
		return nil, fmt.Errorf("nothing was unmarshalled from %s", file)
	}

	return tree, nil
}
