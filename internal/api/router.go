package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.Handle("/deposit", http.HandlerFunc(DepositHandler)).Methods("POST")
	router.Handle("/withdrawal", http.HandlerFunc(WithdrawalHandler)).Methods("POST")
	router.Handle("/callback", http.HandlerFunc(HandleGatewayCallback)).Methods("POST")

	return router
}
