package regex

import (
	"fmt"
	"regexp"
)

var (
	condition *regexp.Regexp
)

// getCondition returns a regular expression used for matching conditional statements.
func getCondition() regexp.Regexp {
	if condition == nil {
		e, err := regexp.Compile(`(?:[\w()]*\s*(?:<|>|>=|<=|={2}|&{2}|\|{2}|!=)\s*[\w()]*)*$`)
		if err != nil {
			panic(fmt.Errorf("error while compiling condition expression: %w", err))
		}
		condition = e
	}

	return *condition
}
