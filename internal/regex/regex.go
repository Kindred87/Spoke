package regex

// IsCondition returns a boolean indicating whether the given string is a condition.
func IsCondition(s string) bool {
	exp := getCondition()

	result := exp.FindStringSubmatch(s)

	return len(result) == 1 && result[0] != ""
}
