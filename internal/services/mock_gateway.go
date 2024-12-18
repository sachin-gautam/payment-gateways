package services

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"payment-gateway/db"
	"payment-gateway/internal/models"
	"time"
)

func ProcessTransactionAsync(transaction db.Transaction) {
	time.Sleep(time.Duration(rand.Intn(3)+2) * time.Second)

	transactionStatus := "success"
	if rand.Intn(2) == 0 {
		transactionStatus = "failed"
	}

	log.Printf("Transaction %d processed with status: %s", transaction.ID, transactionStatus)

	err := db.UpdateTransactionStatus(db.Db, transaction.ID, transactionStatus)
	if err != nil {
		log.Printf("Failed to update transaction status for ID %d: %v", transaction.ID, err)
	} else {
		log.Printf("Transaction status for ID %d updated to: %s", transaction.ID, transactionStatus)
	}

	callbackData := models.TransactionRequest{
		Id:     transaction.ID,
		Status: transactionStatus,
		UserID: transaction.UserID,
		Amount: transaction.Amount,
	}

	callbackURL := "http://localhost:8080/callback"
	jsonData, _ := json.Marshal(callbackData)
	_, err = http.Post(callbackURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Failed to send callback to gateway: %v", err)
	} else {
		log.Println("Callback sent successfully to gateway")
	}
}
