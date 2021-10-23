package handlers

const (
	HandlePathValidateVATID = "/validate/"
)

const (
	errorCodeUnknown = iota + 400000
	errorCodeInvalidRequest
	errorCodeInvalidHeader
	errorCodeEncoding
	errorCodeDecoding

	errorCodeDependencyFailure = 500000
)
