package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	port int
}

func NewServer() *http.Server {

	NewServer := Server{
		port: 8080,
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /save", s.createUser)
	mux.HandleFunc("GET /{id}", s.createUser)

	return mux
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal("create user")
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal("get user")
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
