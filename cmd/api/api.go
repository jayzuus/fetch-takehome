package api

import (
	"log"
	"net/http"
	"takehome/cmd/service/receipts"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	receiptStore := receipts.NewStore()
	receiptService := receipts.NewService(receiptStore)
	receiptHandler := receipts.NewHandler(receiptService)
	receiptHandler.RegisterRoutes(router)
	log.Println("listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
