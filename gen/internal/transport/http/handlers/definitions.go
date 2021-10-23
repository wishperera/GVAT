package handlers

const (
	handlePathValidateVATID = "/validate/"
)

const (
	errorCodeUnknown = iota + 400000
	errorCodeInvalidRequest
	errorCodeInvalidHeader
	errorCodeEncoding
	errorCodeDecoding

	errorCodeDependencyFailure = 500000
)
