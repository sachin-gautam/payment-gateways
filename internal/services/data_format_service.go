package services

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"payment-gateway/internal/models"
)

// DecodeRequest decodes the incoming request based on content type
func DecodeRequest(r *http.Request, request *models.TransactionRequest) error {
	contentType := r.Header.Get("Content-Type")

	// Check if Content-Type is missing
	if contentType == "" {
		return fmt.Errorf("missing content type")
	}

	defer r.Body.Close()

	// Handle the different content types
	switch contentType {
	case "application/json":
		// Decode JSON request
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			return fmt.Errorf("failed to decode JSON: %w", err)
		}
	case "text/xml", "application/xml":
		// Decode XML request
		if err := xml.NewDecoder(r.Body).Decode(request); err != nil {
			return fmt.Errorf("failed to decode XML: %w", err)
		}
	default:
		// Unsupported content type
		return fmt.Errorf("unsupported content type: %s", contentType)
	}

	return nil
}

// EncodeResponse encodes the response based on content type
func EncodeResponse(w http.ResponseWriter, r *http.Request, response interface{}) error {
	contentType := r.Header.Get("Accept")

	// Default to "application/json" if Accept header is missing or set to "*/*"
	if contentType == "" || contentType == "*/*" {
		contentType = "application/json"
	}

	// Set the Content-Type header for the response
	w.Header().Set("Content-Type", contentType)

	// Encode based on content type
	switch contentType {
	case "application/json":
		// Encode JSON response
		if err := json.NewEncoder(w).Encode(response); err != nil {
			return fmt.Errorf("failed to encode JSON: %w", err)
		}
	case "text/xml", "application/xml":
		// Encode XML response
		if err := xml.NewEncoder(w).Encode(response); err != nil {
			return fmt.Errorf("failed to encode XML: %w", err)
		}
	default:
		// Unsupported content type
		return fmt.Errorf("unsupported accept content type: %s", contentType)
	}

	return nil
}
