package main

import (
	"log"

	"github.com/danisbagus/semimarket-product/internal/core/service"
	"github.com/danisbagus/semimarket-product/internal/handler"
	"github.com/danisbagus/semimarket-product/internal/repo"
	"github.com/danisbagus/semimarket-product/pkg/logger"

	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	// sql driver
	client, err := sqlx.Open("mysql", "root:danisbagus@tcp(localhost:9001)/semimarket")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	// multiplexer
	router := mux.NewRouter()

	// wiring
	productRepo := repo.NewProductRepo(client)
	productService := service.NewProductService(productRepo)
	productHandler := handler.ProductHandler{Service: productService}

	transactionRepo := repo.NewTransactionRepo(client)
	TransactionService := service.NewTransactionService(transactionRepo, productRepo)
	transactionHandler := handler.TransactionHandler{Service: TransactionService}

	// routing
	router.HandleFunc("/products", productHandler.GetProductList).Methods(http.MethodGet)
	router.HandleFunc("/products/{product_id:[0-9]+}", productHandler.GetProductDetail).Methods(http.MethodGet)
	router.HandleFunc("/products", productHandler.NewProduct).Methods(http.MethodPost)

	router.HandleFunc("/transactions", transactionHandler.NewTransaction).Methods(http.MethodPost)

	// starting server
	logger.Info("Starting the application")
	log.Fatal(http.ListenAndServe("localhost:9000", router))
}
