package api

import (
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func (a *APIServer) routes() *chi.Mux {
	router := chi.NewRouter()

	// use necessary middleware
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Route("/user", func(r chi.Router) {
		r.Get("/get", a.userHandler.GetUser)
		r.Get("/getall", a.userHandler.GetAllUsers)
		r.Post("/create", a.userHandler.CreateUser)
		r.Delete("/delete", a.userHandler.DeleteUser)
	})

	return router
}
