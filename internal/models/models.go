package models

// a standard request structure for the transactions
type TransactionRequest struct {
	Id     int     `json:"id" xml:"id"`
	Status string  `json:"status" xml:"status"`
	UserID int     `json:"user_id" xml:"user_id"`
	Amount float64 `json:"amount" xml:"amount"`
}

// a standard response structure for the APIs
type APIResponse struct {
	StatusCode int         `json:"status_code" xml:"status_code"`
	Message    string      `json:"message" xml:"message"`
	Data       interface{} `json:"data,omitempty" xml:"data,omitempty"`
}
