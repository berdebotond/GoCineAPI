package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	database "interview_test_KenTech/pkg/db"
	models "interview_test_KenTech/pkg/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFilm() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Creating a film")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var film models.Film
		defer cancel()

		// Bind the JSON to your film struct
		if err := c.BindJSON(&film); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println("Error binding JSON:" + err.Error())
			return
		}
		// checking if film title already exists
		existingFilm := database.FilmData(database.Client, "films").FindOne(ctx, bson.M{"title": film.Title})
		if existingFilm.Err() == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Film with this title already exists"})
			log.Println("Film with this title already exists")
			return
		}

		// Ensure a new ObjectID is generated for this film.
		film.ID = primitive.NewObjectID()

		// get the username from the context
		username, _ := c.Get("username")

		film.Created_by = username.(string)

		// Your logic to insert the film, ensuring uniqueness by title
		result, err := database.FilmData(database.Client, "films").InsertOne(ctx, film)
		if err != nil {
			log.Println("Error inserting film:" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, result)
	}
}

func GetFilms() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Construct a dynamic query based on query parameters
		query := bson.M{}
		log.Println("Getting films")
		if title := c.Query("title"); title != "" {
			query["title"] = bson.M{"$regex": primitive.Regex{Pattern: title, Options: "i"}}
		}
		if releaseDate := c.Query("releaseDate"); releaseDate != "" {
			query["releaseDate"] = releaseDate
		}

		if genre := c.Query("genre"); genre != "" {
			query["genre"] = genre
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err := database.FilmData(database.Client, "films").Find(ctx, query)
		if err != nil {
			log.Println("Error fetching films:" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while fetching films"})
			return
		}

		var films []models.Film
		if err := cursor.All(ctx, &films); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while decoding films"})
			log.Println("Error decoding films:" + err.Error())
			return
		}

		c.JSON(http.StatusOK, films)
	}
}

// exmaple requst for getfilms endpoint query parameters /films?title=avengers&releaseDate=2021-01-01&genre=action
func GetFilm() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filmID, _ := primitive.ObjectIDFromHex(c.Param("film_id"))
		var film models.Film
		err := database.FilmData(database.Client, "films").FindOne(ctx, bson.M{"_id": filmID}).Decode(&film)
		if err != nil {
			log.Println("Error fetching film:" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Film not found"})
			return
		}
		c.JSON(http.StatusOK, film)
	}
}

func UpdateFilm() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filmID, _ := primitive.ObjectIDFromHex(c.Param("film_id"))
		log.Println("Updating a film with ID: ", filmID)
		var film models.Film
		if err := c.BindJSON(&film); err != nil {
			log.Println("Error binding JSON:" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// checking if film title already exists
		existingFilm := database.FilmData(database.Client, "films").FindOne(ctx, bson.M{"title": film.Title})
		if existingFilm.Err() == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Film with this title already exists"})
			log.Println("Film with this title already exists")
			return
		}
		found_film := models.Film{}
		err := database.FilmData(database.Client, "films").FindOne(ctx, bson.M{"_id": filmID}).Decode(&found_film)
		if err != nil {
			log.Println("Error fetching film:" + err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": "Film not found"})
			return
		}

		// Check if the current user is the creator of the film
		username, _ := c.Get("username")

		if found_film.Created_by != username {
			log.Println("Error updating film: You can only update films you've created")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You can only update films you've created"})
			return
		}
		film.ID = filmID
		film.Created_by = found_film.Created_by
		// Your logic to update the film, ensuring the user is the owner
		result, err := database.FilmData(database.Client, "films").UpdateOne(ctx, bson.M{"_id": filmID}, bson.M{"$set": film})
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while updating the film"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func DeleteFilm() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		log.Println("Deleting a film")
		filmID, _ := primitive.ObjectIDFromHex(c.Param("film_id"))
		var film models.Film
		err := database.FilmData(database.Client, "films").FindOne(ctx, bson.M{"_id": filmID}).Decode(&film)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Film not found"})
			return
		}

		// Check if the current user is the creator of the film
		username, _ := c.Get("username")
		if film.Created_by != username {
			log.Println("Error deleting film: You can only delete films you've created")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You can only delete films you've created"})
			return
		}
		// Your logic to delete the film, ensuring the user is the owner
		result, err := database.FilmData(database.Client, "films").DeleteOne(ctx, bson.M{"_id": filmID})
		if err != nil {
			log.Println("Error deleting film:" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while deleting the film"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
