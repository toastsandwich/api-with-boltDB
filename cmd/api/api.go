package api

import (
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/toastsandwich/restraunt-api-system/handler"
)

type APIServer struct {
	db          *bolt.DB
	userHandler *handler.UserHandler
	Addr        string
}

func NewAPIServer(addr string, db *bolt.DB) (*APIServer, error) {
	userHandler, err := handler.NewUserHandler(db)
	if err != nil {
		return nil, err
	}
	return &APIServer{
		Addr:        addr,
		db:          db,
		userHandler: userHandler,
	}, nil
}

func (a *APIServer) Run() error {
	server := http.Server{
		Addr:    a.Addr,
		Handler: a.routes(),
	}
	log.Println("server started at", a.Addr)
	return server.ListenAndServe()
}

func (a *APIServer) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /user/get", a.userHandler.GetUser)
	mux.HandleFunc("POST /user/create", a.userHandler.CreateUser)
	return mux
}
