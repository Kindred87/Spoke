package cache

import "golang.org/x/exp/slices"

// removeFrom returns a modified copy of the given slice in which the given value is removed.
func removeFrom(value string, slice []string) []string {
	return slices.Delete(slice, slices.Index(slice, value), slices.Index(slice, value)+1)
}
