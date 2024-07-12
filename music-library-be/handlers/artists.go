package handlers

import (
	// "encoding/json"

	"music-library-be/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/exp/rand"
)

type Artist = models.Artist

type ArtistHandler struct {
	db *gorm.DB
}

func CreateArtistHandler(gorm *gorm.DB) ArtistHandler {
	return ArtistHandler{
		db: gorm,
	}
}

func (handler ArtistHandler) GetArtists(c *gin.Context) {
	var artists []Artist
	handler.db.Find(&artists)

	var returnArtists []models.ReturnJSONArtist

	for i := range artists {
		returnArtists = append(returnArtists, models.ReturnJSONArtist{
			ID:   artists[i].ID,
			Name: artists[i].Name,
		})
	}

	c.JSON(http.StatusOK, returnArtists)
}

func (handler ArtistHandler) GetArtist(c *gin.Context) {
	var artist Artist
	id := c.Param("id")
	handler.db.First(&artist, "id = ?", id)

	var returnArtist models.ReturnJSONArtist

	returnArtist.ID = artist.ID
	returnArtist.Name = artist.Name

	c.JSON(http.StatusOK, returnArtist)
}

func (handler ArtistHandler) CreateArtist(c *gin.Context) {
	var artist Artist
	err := c.ShouldBindJSON(&artist)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.db.Create(&artist)

	c.JSON(http.StatusCreated, artist)
}

func (handler ArtistHandler) UpdateArtist(c *gin.Context) {
	var artist Artist
	id := c.Param("id")
	handler.db.First(&artist, "id = ?", id)

	err := c.ShouldBindJSON(&artist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.db.Update(&artist)

	c.JSON(http.StatusOK, artist)
}

func (handler ArtistHandler) DeleteArtist(c *gin.Context) {
	var artist Artist
	id := c.Param("id")
	handler.db.First(&artist, "id = ?", id)

	handler.db.Delete(&artist)

	c.JSON(http.StatusOK, gin.H{"message": "artist deleted"})
}

func (handler ArtistHandler) GetRandomArtists(c *gin.Context) {
	var artists []Artist
	handler.db.Find(&artists)

	rand.Seed(uint64(time.Now().UnixNano()))
	rand.Shuffle(len(artists), func(i, j int) { artists[i], artists[j] = artists[j], artists[i] })

	c.JSON(http.StatusOK, artists)
}
