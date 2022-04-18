package model

var (
	messageFieldShouldNotBeEmpty = "Should not be null or empty string"
)

const (
	// APIResponseTypeItem is constant for APIResponse.Type
	APIResponseTypeItem = "item"
	// APIResponseTypeList is constant for APIResponse.Type
	APIResponseTypeList = "list"
	// APIResponseTypeRequestError is constant for APIResponse.Type
	APIResponseTypeRequestError = "request_error"
)

// APIResponseType is constant for APIResponse.Type
type APIResponseType string

// APIResponse is common api response for all mmg services
type APIResponse struct {
	Type APIResponseType `json:"type"`
	Data interface{}     `json:"data"`
}

// APIErrorResponse contains API errors
type APIErrorResponse struct {
	Errors interface{} `json:"errors"`
}
