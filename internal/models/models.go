package models

// TransactionRequest represents a transaction request (deposit or withdrawal)
// @Description Represents the request payload for initiating a transaction (either deposit or withdrawal)
// @Param user_id body int true "User ID associated with the transaction"
// @Param amount body float64 true "Amount for the transaction (in decimal format)"
type TransactionRequest struct {
	Id     int     `json:"id" xml:"id" swaggerignore:"true"`
	Status string  `json:"status" xml:"status" swaggerignore:"true"`
	UserID int     `json:"user_id" xml:"user_id"`
	Amount float64 `json:"amount" xml:"amount"`
}

// APIResponse represents the standard response structure for the API
// @Description Represents the response payload returned after processing a transaction request
// @Param status_code body int true "HTTP status code"
// @Param message body string true "Message describing the result of the request"
// @Param data body object false "The data returned by the request, if applicable"
type APIResponse struct {
	StatusCode int         `json:"status_code" xml:"status_code"`
	Message    string      `json:"message" xml:"message"`
	Data       interface{} `json:"data,omitempty" xml:"data,omitempty"`
}
