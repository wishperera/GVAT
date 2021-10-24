package handlers

const (
	HandlePathValidateVATID = "/validate/"
)

const (
	errorCodeUnknown = iota + 400000
	errorCodeInvalidRequest
	errorCodeInvalidHeader

	errorCodeDependencyFailure = 500000
)
