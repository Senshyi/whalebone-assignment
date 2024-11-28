package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"whalebone-assignment/internal/database"
	"whalebone-assignment/internal/models"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func createTmpDB(t testing.TB) (database.Service, func()) {
	t.Helper()
	dbName := os.Getenv("TEST_DB_URL")
	println("<<<<<<<<", dbName)
	service := database.New(dbName)
	teardown := func() {
		os.Remove(dbName)
	}
	return service, teardown

}

func TestGetAndPostUser(t *testing.T) {
	envPath := filepath.Join("..", "..", ".env")
	_ = godotenv.Load(envPath)

	createTmpDB(t)
	dbService, teardown := createTmpDB(t)
	defer teardown()

	t.Run("create user and get user", func(t *testing.T) {
		testID, _ := uuid.Parse("ae593b85-b9a2-4386-ad71-7b62287d7c24")
		testUser := models.ResponseUser{
			ID:          testID,
			Name:        "test name",
			Email:       "test.email@example.com",
			DateOfBirth: "2020-01-01T12:12:34+00:00",
		}
		jsonUser, _ := json.Marshal(testUser)

		s := NewServer(dbService, 7000)

		postResponse := httptest.NewRecorder()
		postReq, _ := http.NewRequest(http.MethodPost, "/save", bytes.NewBuffer(jsonUser))
		s.Handler.ServeHTTP(postResponse, postReq)
		if postResponse.Result().StatusCode != http.StatusCreated {
			t.Fatalf("did not get correct status, got %d, want %d", postResponse.Result().StatusCode, http.StatusCreated)
		}

		getResponse := httptest.NewRecorder()
		getReq, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", testID), nil)
		s.Handler.ServeHTTP(getResponse, getReq)

		var got models.ResponseUser
		err := json.NewDecoder(getResponse.Body).Decode(&got)
		if err != nil {
			t.Fatalf("unable to parse response from server %q into slice of Player, '%v'", getResponse.Body, err)
		}

		if getResponse.Result().StatusCode != http.StatusOK {
			t.Fatalf("did not get correct status, got %d, want %d", getResponse.Result().StatusCode, http.StatusOK)
		}
		if !reflect.DeepEqual(got, testUser) {
			t.Fatalf("did not get correct response, got %+v, want %+v", got, testUser)
		}

	})
}
