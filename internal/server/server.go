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
	"whalebone-assignment/internal/validator"

	"github.com/google/uuid"

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
		validator.Validator
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Unable to parse request body: %v", err)
		respondWithError(w, http.StatusBadRequest, "unable to parse request body")
		return
	}

	validId, err := uuid.Parse(params.ID)
	if err != nil {
		log.Printf("Unable to parse UUID: %v", err)
		params.AddProblem("id", "invalid UUID")
	}

	params.Check(validator.NotBlank(params.Name), "name", "field name cannot be empty")
	params.Check(validator.MaxChars(params.Name, 100), "name", "field name must be max 100 chars long")
	params.Check(validator.MatchRegex(params.Email, validator.EmailRegex), "email", "field email must be a valid email address")

	parsedTime, err := time.Parse("2006-01-02T15:04:05-07:00", params.DateOfBirth)
	if err != nil {
		params.AddProblem("date_of_birth", "field needs to be a valid time")
	}

	if !params.Valid() {
		respondWithError(w, http.StatusUnprocessableEntity, "unable to process request")
		return
	}

	err = s.users.Insert(models.DatabaseUser{
		ID:          validId,
		Name:        params.Name,
		Email:       params.Email,
		DateOfBirth: parsedTime,
	})
	if err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, "unable to process request")
		return
	}

	respondWithJSON(w, http.StatusCreated, struct{}{})
}
func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	validId, err := uuid.Parse(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid id")
	}

	user, err := s.users.GetOne(validId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "user not found")
	}

	respondWithJSON(w, http.StatusOK, models.DatabaseUserToResponseUser(user))
}
