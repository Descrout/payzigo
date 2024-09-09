package responses

import "fmt"

type ErrorResponse struct {
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (e *ErrorResponse) HasError() bool {
	return e.ErrorCode != ""
}

func (e *ErrorResponse) Error() error {
	return fmt.Errorf("Error Code [%s]: %s", e.ErrorCode, e.ErrorMessage)
}
