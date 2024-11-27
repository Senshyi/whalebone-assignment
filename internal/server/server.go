package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"whalebone-assignment/internal/models"

	_ "github.com/joho/godotenv/autoload"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	port int

	users *models.UserModel
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.DatabaseUser{})

	NewServer := Server{
		port:  port,
		users: &models.UserModel{Db: db},
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
	mux.HandleFunc("GET /{id}", s.getUser)

	return mux
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		DateOfBirth string `json:"date_of_birth"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Unable to parse request body: %v", err)
		respondWithError(w, http.StatusBadRequest, "unable to parse request body")
		return
	}

	parsedTime, err := time.Parse("2006-01-02T15:04:05-07:00", params.DateOfBirth)
	if err != nil {
		log.Printf("error parsing time: ", err)
	}

	err = s.users.Insert(models.DatabaseUser{
		ID:          params.ID,
		Name:        params.Name,
		Email:       params.Email,
		DateOfBirth: parsedTime,
	})
	if err != nil {
		log.Printf("unable to create User: %v", err)
	}

	respondWithJSON(w, http.StatusCreated, struct{}{})
}
func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user, err := s.users.GetOne(id)
	if err != nil {
		log.Printf("unable to get user from db: %v", err)
	}

	respondWithJSON(w, http.StatusOK, models.DatabaseUserToResponseUser(user))
}
