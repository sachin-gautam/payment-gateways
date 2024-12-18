package services

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"payment-gateway/internal/models"
)

func DecodeRequest(r *http.Request, request *models.TransactionRequest) error {
	contentType := r.Header.Get("Content-Type")

	if contentType == "" {
		return fmt.Errorf("missing content type")
	}

	defer r.Body.Close()

	switch contentType {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			return fmt.Errorf("failed to decode JSON: %w", err)
		}
	case "text/xml", "application/xml":
		if err := xml.NewDecoder(r.Body).Decode(request); err != nil {
			return fmt.Errorf("failed to decode XML: %w", err)
		}
	default:
		return fmt.Errorf("unsupported content type: %s", contentType)
	}

	return nil
}

func EncodeResponse(w http.ResponseWriter, r *http.Request, response interface{}) error {
	contentType := r.Header.Get("Accept")

	if contentType == "" || contentType == "*/*" {
		contentType = "application/json"
	}

	w.Header().Set("Content-Type", contentType)

	switch contentType {
	case "application/json":
		if err := json.NewEncoder(w).Encode(response); err != nil {
			return fmt.Errorf("failed to encode JSON: %w", err)
		}
	case "text/xml", "application/xml":
		if err := xml.NewEncoder(w).Encode(response); err != nil {
			return fmt.Errorf("failed to encode XML: %w", err)
		}
	default:
		return fmt.Errorf("unsupported accept content type: %s", contentType)
	}

	return nil
}
