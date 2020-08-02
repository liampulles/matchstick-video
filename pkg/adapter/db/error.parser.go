package db

import "regexp"

var uniqConRegExp = regexp.MustCompile(`(?m)violates unique constraint`)
var noRowsRegExp = regexp.MustCompile(`(?m)no rows in result set`)

// ErrorParser analyses external errors to create matchstick-video variants.
type ErrorParser interface {
	FromDBRowScan(err error, _type string) error
}

// ErrorParserImpl implements ErrorParser
type ErrorParserImpl struct{}

// Check we implement the interface
var _ ErrorParser = &ErrorParserImpl{}

// NewErrorParserImpl is a constructor
func NewErrorParserImpl() *ErrorParserImpl {
	return &ErrorParserImpl{}
}

// FromDBRowScan tries to extract errors from the response to db row scans
func (e *ErrorParserImpl) FromDBRowScan(err error, _type string) error {
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
