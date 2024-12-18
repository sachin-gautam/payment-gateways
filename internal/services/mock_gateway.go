package services

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"payment-gateway/db"
	"payment-gateway/internal/models"
	"strconv"
	"time"
)

func ProcessTransactionAsync(transaction db.Transaction) {
	time.Sleep(time.Duration(rand.Intn(3)+2) * time.Second)

	amountBytes, err := UnmaskData(transaction.Amount)
	if err != nil {
		log.Printf("Failed to unmask transaction amount for ID %d: %v", transaction.ID, err)
		return
	}

	amountStr := string(amountBytes)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.Printf("Failed to convert unmasked amount to float for transaction ID %d: %v", transaction.ID, err)
		return
	}

	transactionStatus := "success"
	if rand.Intn(2) == 0 {
		transactionStatus = "failed"
	}

	log.Printf("Transaction %d processed with status: %s", transaction.ID, transactionStatus)

	err = db.UpdateTransactionStatus(db.Db, transaction.ID, transactionStatus)
	if err != nil {
		log.Printf("Failed to update transaction status for ID %d: %v", transaction.ID, err)
	} else {
		log.Printf("Transaction status for ID %d updated to: %s", transaction.ID, transactionStatus)
	}

	callbackData := models.TransactionRequest{
		Id:     transaction.ID,
		Status: transactionStatus,
		UserID: transaction.UserID,
		Amount: amount,
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
