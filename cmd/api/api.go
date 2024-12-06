package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"go-backend-api-jwt-mysql/service/cart"
	"go-backend-api-jwt-mysql/service/order"
	"go-backend-api-jwt-mysql/service/product"
	"go-backend-api-jwt-mysql/service/user"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterStore(subrouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Listening on: ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
