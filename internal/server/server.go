package server

import (
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

	// db *gorm.DB

	users *models.UserModel
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

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
	parsedTim, err := time.Parse("2006-01-02T15:04:05-07:00", "2020-01-01T12:12:34+00:00")
	if err != nil {
		log.Printf("error parsing time: ", err)
	}

	newUser := models.User{ID: "ca3ae13c-3c97-49fc-b4bb-86f04d934142", Name: "Joseph", Email: "email@gmail.com", DateOfBirth: parsedTim}
	err = s.users.Insert(newUser)
	if err != nil {
		log.Printf("unable to create User: %v", err)
	}

	respondWithJSON(w, http.StatusCreated, struct{}{})
}
func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	user, err := s.users.GetOne("ca2ae13c-3c97-49fc-b4bb-86f04d934142")
	if err != nil {
		log.Printf("unable to get user from db: %v", err)
	}

	respondWithJSON(w, http.StatusOK, user)
}
