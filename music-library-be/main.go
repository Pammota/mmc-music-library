package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"music-library-be/handlers"
	"music-library-be/models"
	"music-library-be/parser"
)

func main() {
	var dbString string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DBNAME"))

	fmt.Println(dbString)

	db, err := gorm.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Automigrate the tables
	db.AutoMigrate(&models.Song{})
	db.AutoMigrate(&models.Album{})
	db.AutoMigrate(&models.Artist{})

	songHandler := handlers.CreateSongHandler(db)
	albumHandler := handlers.CreateAlbumHandler(db)
	artistHandler := handlers.CreateArtistHandler(db)
	searchHandler := handlers.CreateSearchHandler(db)

	// Import input data
	parser.ParseJSON("data.json", db)

	router := gin.Default()

	// CORS Setup
	config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "active",
		})
	})

	// Songs CRUD
	router.GET("/songs", songHandler.GetSongs)
	router.GET("/songs/:id", songHandler.GetSong)
	router.POST("/songs", songHandler.CreateSong)
	router.PUT("/songs/:id", songHandler.UpdateSong)
	router.DELETE("/songs/:id", songHandler.DeleteSong)

	// Query params & random
	router.GET("/songs/byAlbum/:albumId", songHandler.GetSongsByAlbum)
	router.GET("/songs/byArtist/:artistId", songHandler.GetSongsByArtist)
	router.GET("/songs/random", songHandler.GetRandomSongs)

	// Albums CRUD
	router.GET("/albums", albumHandler.GetAlbums)
	router.GET("/albums/:id", albumHandler.GetAlbum)
	router.POST("/albums", albumHandler.CreateAlbum)
	router.PUT("/albums/:id", albumHandler.UpdateAlbum)
	router.DELETE("/albums/:id", albumHandler.DeleteAlbum)

	// Query params & random
	router.GET("/albums/byArtist/:artistId", albumHandler.GetAlbumsByArtist)
	router.GET("/albums/random", albumHandler.GetRandomAlbums)

	// Artists CRUD
	router.GET("/artists", artistHandler.GetArtists)
	router.GET("/artists/:id", artistHandler.GetArtist)
	router.POST("/artists", artistHandler.CreateArtist)
	router.PUT("/artists/:id", artistHandler.UpdateArtist)
	router.DELETE("/artists/:id", artistHandler.DeleteArtist)

	router.GET("/artists/random", artistHandler.GetRandomArtists)

	// Global search
	router.GET("/search/:question", searchHandler.Search)

	router.Run("0.0.0.0:8080")
}
