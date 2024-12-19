// router.go
package api

import (
	"net/http"

	_ "payment-gateway/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Route for deposit
	router.Handle("/deposit", http.HandlerFunc(DepositHandler)).Methods("POST")

	// Route for withdrawal
	router.Handle("/withdrawal", http.HandlerFunc(WithdrawalHandler)).Methods("POST")

	// Route for callback
	router.Handle("/callback", http.HandlerFunc(HandleGatewayCallback)).Methods("POST")

	//Route for swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
