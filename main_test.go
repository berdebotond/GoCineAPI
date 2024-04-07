package main_test

import (
	"bytes"
	"context"
	"fmt"
	"interview_test_KenTech/config"
	db "interview_test_KenTech/pkg/db"
	auth "interview_test_KenTech/pkg/helper"
	"interview_test_KenTech/pkg/models"
	"interview_test_KenTech/pkg/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	// Set up the Gin router and load the routes
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	routes.UserRoutes(router)
	routes.AuthRoutes(router)
	routes.FilmRoutes(router)

	// Connect to the MongoDB test database
	config.CFG.MongoURI = "mongodb://user:password@localhost:27017"
	// reset the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := db.Client.Database("test").Drop(ctx)
	if err != nil {
		fmt.Println("Error dropping test database , make sure mongodb is running: ", err)
	}
	db.Client = db.DBSet()

	// Run the test suite
	os.Exit(m.Run())
}

func TestSignup(t *testing.T) {
	// Test case: Successful signup
	requestBody := []byte(`{
		"username": "testuser",
		"password": "password123",
		"user_type": "USER"
	}`)
	req, _ := http.NewRequest("POST", "/users/signup", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test case: Signup with existing username
	requestBody = []byte(`{
		"username": "testuser",
		"password": "password123",
		"user_type": "USER"
	}`)
	req, _ = http.NewRequest("POST", "/users/signup", bytes.NewBuffer(requestBody))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestLogin(t *testing.T) {
	// Test case: Successful login
	requestBody := []byte(`{
		"username": "testuser",
		"password": "password123"
	}`)
	req, _ := http.NewRequest("POST", "/users/login", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test case: Login with invalid credentials
	requestBody = []byte(`{
		"username": "testuser",
		"password": "wrongpassword"
	}`)
	req, _ = http.NewRequest("POST", "/users/login", bytes.NewBuffer(requestBody))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateFilm(t *testing.T) {
	// Test case: Successful film creation
	requestBody := []byte(`{
		"title": "Test Film",
		"director": "Test Director",
		"releaseDate": "2023-04-07T00:00:00Z",
		"cast": ["Actor 1", "Actor 2"],
		"genre": ["Drama", "Comedy"],
		"synopsis": "This is a test film."
	}`)
	req, _ := http.NewRequest("POST", "/films", bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+getTestToken("testuser", "USER"))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Test case: Create film with existing title
	requestBody = []byte(`{
		"title": "Test Film",
		"director": "Test Director",
		"releaseDate": "2023-04-07T00:00:00Z",
		"cast": ["Actor 1", "Actor 2"],
		"genre": ["Drama", "Comedy"],
		"synopsis": "This is a test film."
	}`)
	req, _ = http.NewRequest("POST", "/films", bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+getTestToken("testuser", "USER"))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetFilms(t *testing.T) {
	// Test case: Get all films
	req, _ := http.NewRequest("GET", "/films", nil)
	req.Header.Set("Authorization", "Bearer "+getTestToken("testuser", "USER"))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test case: Get films by query parameters
	req, _ = http.NewRequest("GET", "/films?title=Test&releaseDate=2023-04-07&genre=Drama", nil)
	req.Header.Set("Authorization", "Bearer "+getTestToken("testuser", "USER"))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateFilm(t *testing.T) {
	// Test case: Successful film update
	film := models.Film{
		Title:       "Test Film",
		Director:    "Test Director",
		ReleaseDate: time.Now(),
		Cast:        []string{"Actor 1", "Actor 2"},
		Genre:       []string{"Drama", "Comedy"},
		Synopsis:    "This is a test film.",
		User_ID:     primitive.NewObjectID(),
		Created_by:  "testuser",
	}
	filmID := insertTestFilm(t, film)

	requestBody := []byte(`{
		"title": "Updated Test Film",
		"director": "Updated Test Director",
		"releaseDate": "2023-04-08T00:00:00Z",
		"cast": ["Actor 1", "Actor 2", "Actor 3"],
		"genre": ["Drama", "Comedy", "Action"],
		"synopsis": "This is an updated test film."
	}`)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/films/%s", filmID.Hex()), bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+getTestToken("testuser", "USER"))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test case: Update film with existing title
	requestBody = []byte(`{
		"title": "Updated Test Film",
		"director": "Test Director",
		"releaseDate": "2023-04-07T00:00:00Z",
		"cast": ["Actor 1", "Actor 2"],
		"genre": ["Drama", "Comedy"],
		"synopsis": "This is a test film."
	}`)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/films/%s", filmID.Hex()), bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+getTestToken("testuser", "USER"))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test case: Update film as unauthorized user
	requestBody = []byte(`{
		"title": "Unauthorized Update",
		"director": "Unauthorized Director",
		"releaseDate": "2023-04-09T00:00:00Z",
		"cast": ["Actor 1", "Actor 2", "Actor 3"],
		"genre": ["Drama", "Comedy", "Action"],
		"synopsis": "This is an unauthorized update."
	}`)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/films/%s", filmID.Hex()), bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", getTestToken("anotheruser", "USER"))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestDeleteFilm(t *testing.T) {
	// Test case: Successful film deletion
	film := models.Film{
		Title:       "Test Film",
		Director:    "Test Director",
		ReleaseDate: time.Now(),
		Cast:        []string{"Actor 1", "Actor 2"},
		Genre:       []string{"Drama", "Comedy"},
		Synopsis:    "This is a test film.",
		User_ID:     primitive.NewObjectID(),
		Created_by:  "testuser",
	}
	filmID := insertTestFilm(t, film)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/films/%s", filmID.Hex()), nil)
	req.Header.Set("Authorization", "Bearer "+getTestToken("testuser", "USER"))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test case: Delete film as unauthorized user
	film = models.Film{
		Title:       "Unauthorized Test Film",
		Director:    "Unauthorized Director",
		ReleaseDate: time.Now(),
		Cast:        []string{"Actor 1", "Actor 2"},
		Genre:       []string{"Drama", "Comedy"},
		Synopsis:    "This is an unauthorized test film.",
		User_ID:     primitive.NewObjectID(),
		Created_by:  "anotheruser",
	}
	filmID = insertTestFilm(t, film)

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/films/%s", filmID.Hex()), nil)
	req.Header.Set("Authorization", "Bearer "+getTestToken("testuser", "USER"))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func getTestToken(username, userType string) string {
	token, _ := auth.GenerateAllTokens(username, userType, "test-user-id")
	return token
}

func insertTestFilm(t *testing.T, film models.Film) primitive.ObjectID {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	film.ID = primitive.NewObjectID()
	_, err := db.FilmData(db.Client, "films").InsertOne(ctx, film)
	assert.NoError(t, err)
	return film.ID
}
