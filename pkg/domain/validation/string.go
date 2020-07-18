package validation

import "strings"

// IsBlank returns true if str is empty
// or composed entirely of whitespace.
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// IsTrimmed returns true if str contains no
// whitespace as a prefix or suffix.
func IsTrimmed(str string) bool {
	return strings.TrimSpace(str) == str
}
