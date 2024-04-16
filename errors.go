package atol_client

import "fmt"

type ErrorType string

const (
	AuthErrorType     ErrorType = "auth"
	JsonErrorType     ErrorType = "json"
	ParsingErrorType  ErrorType = "parsing"
	ExternalErrorType ErrorType = "external"
)

type ATOLClientError struct {
	Type    ErrorType `json:"type"`
	IsInner bool      `json:"isInner"`
	Message string    `json:"message"`
}

func NewAuthError(msg string, isInner bool) *ATOLClientError {
	return newATOLClientError(AuthErrorType, msg, isInner)
}

func NewJsonError(msg string, isInner bool) *ATOLClientError {
	return newATOLClientError(JsonErrorType, msg, isInner)
}

func NewParsingError(msg string, isInner bool) *ATOLClientError {
	return newATOLClientError(ParsingErrorType, msg, isInner)
}

func NewExternalError(msg string, isInner bool) *ATOLClientError {
	return newATOLClientError(ExternalErrorType, msg, isInner)
}

func newATOLClientError(errType ErrorType, msg string, isInner bool) *ATOLClientError {
	return &ATOLClientError{
		Type:    errType,
		IsInner: isInner,
		Message: msg,
	}
}

func (err *ATOLClientError) Error() string {
	return fmt.Sprintf("ATOL client error (type: %s, message: %s, isInner: %t)", err.Type, err.Message, err.IsInner)
}
