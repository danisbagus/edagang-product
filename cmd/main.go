package main

import (
	"log"

	"github.com/danisbagus/edagang-pkg/logger"
	"github.com/danisbagus/edagang-product/internal/core/service"
	"github.com/danisbagus/edagang-product/internal/handler"
	"github.com/danisbagus/edagang-product/internal/middleware"
	"github.com/danisbagus/edagang-product/internal/repo"

	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	// sql driver
	client := GetClient()

	// multiplexer
	router := mux.NewRouter()

	// wiring
	productRepo := repo.NewProductRepo(client)
	productService := service.NewProductService(productRepo)
	productHandler := handler.ProductHandler{Service: productService}

	// routing
	router.HandleFunc("/products", productHandler.GetProductList).Methods(http.MethodGet).Name("GetProductList")
	router.HandleFunc("/products/{product_id:[0-9]+}", productHandler.GetProductDetail).Methods(http.MethodGet).Name("GetProductDetail")
	router.HandleFunc("/products", productHandler.NewProduct).Methods(http.MethodPost).Name("NewProduct")
	router.HandleFunc("/products/{product_id:[0-9]+}", productHandler.RemoveProduct).Methods(http.MethodDelete).Name("RemoveProduct")

	// middleware
	authMiddleware := middleware.AuthMiddleware{repo.NewAuthRepo()}
	router.Use(authMiddleware.AuthorizationHandler())

	// starting server
	logger.Info("Starting product service")
	log.Fatal(http.ListenAndServe("localhost:9000", router))
}

func GetClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:danisbagus@tcp(localhost:9001)/edagang")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
