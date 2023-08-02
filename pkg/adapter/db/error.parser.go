package db

import "regexp"

var uniqConRegExp = regexp.MustCompile(`(?m)violates unique constraint`)
var noRowsRegExp = regexp.MustCompile(`(?m)no rows in result set`)

// FromDBRowScan tries to extract errors from the response to db row scans
var FromDBRowScan = func(err error, _type string) error {
	// See if there is a uniqueness constraint violation
	if isUniquenessConstraintError(err) {
		return NewUniqueConstraintError(err)
	}

	// See if the error indicates no rows were returned
	if isNoRowsError(err) {
		return NewNotFoundError(_type)
	}

	// Else, return the original error
	return err
}

func isUniquenessConstraintError(err error) bool {
	return matchesRegExp(err, uniqConRegExp)
}

func isNoRowsError(err error) bool {
	return matchesRegExp(err, noRowsRegExp)
}

func matchesRegExp(err error, reg *regexp.Regexp) bool {
	return reg.Match([]byte(err.Error()))
}
