package inventory

// Validator validates that objects are "well formed".
type Validator interface {
	ValidateCreateItemVO(*CreateItemVO) error
}
