package validation

import "strings"

// IsNotBlank returns true if str is not empty
// and is not composed entirely of whitespace.
func IsNotBlank(str string) bool {
	return len(strings.TrimSpace(str)) != 0
}

// IsTrimmed returns true if str contains no
// whitespace as a prefix or suffix.
func IsTrimmed(str string) bool {
	return strings.TrimSpace(str) == str
}
