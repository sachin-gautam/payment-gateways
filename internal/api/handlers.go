package api

import (
	"fmt"
	"log"
	"net/http"
	"payment-gateway/db"
	"payment-gateway/internal/models"
	"payment-gateway/internal/services"
	"time"
)

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	var transactionRequest models.TransactionRequest
	err := services.DecodeRequest(r, &transactionRequest)
	if err != nil {
		log.Printf("Failed to decode request: %v", err)
		services.EncodeResponse(w, r, models.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message: fmt.Sprintf("Failed to decode request: %v",
				err)})
		return
	}

	log.Printf("Deposit request received from user_id %+v", transactionRequest.UserID)

	userID := transactionRequest.UserID
	amount := transactionRequest.Amount
	maskedAmount := services.MaskData([]byte(fmt.Sprintf("%.2f", amount)))

	user, err := db.GetUserByID(db.Db, userID)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		services.EncodeResponse(w, r, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message: fmt.Sprintf("Error fetching user: %v",
				err)})
		return
	}

	gateways, err := db.GetAvailableGateways(db.Db, user.CountryID)
	if err != nil || len(gateways) == 0 {
		http.Error(w,
			"No available gateways",
			http.StatusInternalServerError)
		return
	}

	transaction := db.Transaction{
		Amount:    maskedAmount,
		Type:      "deposit",
		Status:    "pending",
		UserID:    userID,
		GatewayID: gateways[0].ID,
		CountryID: user.CountryID,
		CreatedAt: time.Now(),
	}

	err = db.CreateTransaction(db.Db, &transaction)
	if err != nil {
		log.Printf("Error processing transaction: %v", err)
		services.EncodeResponse(w, r, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message: fmt.Sprintf("Error processing transaction: %v",
				err)})
		return
	}

	log.Printf("Transaction created: %+v", transaction)

	go services.ProcessTransactionAsync(transaction)

	err = services.EncodeResponse(w, r, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Transaction initiated successfully",
		Data:       transaction,
	})
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func WithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	var transactionRequest models.TransactionRequest
	err := services.DecodeRequest(r, &transactionRequest)
	if err != nil {
		log.Printf("Failed to decode request: %v", err)
		services.EncodeResponse(w, r, models.APIResponse{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Failed to decode request: %v", err)})
		return
	}

	log.Printf("Withdrawal request received from user_id: %+v", transactionRequest.UserID)

	userID := transactionRequest.UserID
	amount := transactionRequest.Amount
	maskedAmount := services.MaskData([]byte(fmt.Sprintf("%.2f", amount)))

	user, err := db.GetUserByID(db.Db, userID)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		services.EncodeResponse(w, r, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message: fmt.Sprintf("Error fetching user: %v",
				err)})
		return
	}

	gateways, err := db.GetAvailableGateways(db.Db, user.CountryID)
	if err != nil || len(gateways) == 0 {
		http.Error(w, "No available gateways", http.StatusInternalServerError)
		return
	}

	transaction := db.Transaction{
		Amount:    maskedAmount,
		Type:      "withdrawal",
		Status:    "pending",
		UserID:    userID,
		GatewayID: gateways[0].ID,
		CountryID: user.CountryID,
		CreatedAt: time.Now(),
	}

	err = db.CreateTransaction(db.Db, &transaction)
	if err != nil {
		log.Printf("Error processing transaction: %v", err)
		services.EncodeResponse(w, r, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Error processing transaction: %v", err)})
		return
	}

	log.Printf("Transaction created: %+v", transaction)

	go services.ProcessTransactionAsync(transaction)

	err = services.EncodeResponse(w, r, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Withdrawal initiated successfully",
		Data:       transaction,
	})
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func HandleGatewayCallback(w http.ResponseWriter, r *http.Request) {
	var callbackData models.TransactionRequest
	err := services.DecodeRequest(r, &callbackData)
	if err != nil {
		log.Printf("Failed to decode callback: %v", err)
		services.EncodeResponse(w, r, models.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Failed to decode callback: %v", err)})
		return
	}

	log.Printf("Callback received: %+v", callbackData)

	transactionID := callbackData.Id
	transactionStatus := callbackData.Status

	go func() {
		err := db.UpdateTransactionStatus(db.Db, transactionID, transactionStatus)
		if err != nil {
			log.Printf("Failed to update transaction status for ID %d: %v", transactionID, err)
		} else {
			log.Printf("Successfully updated transaction status for ID %d", transactionID)
		}
	}()

	err = services.EncodeResponse(w, r, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Callback processed successfully",
	})
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
