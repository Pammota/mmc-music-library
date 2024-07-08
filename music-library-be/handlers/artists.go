package handlers

import (
	// "encoding/json"

	"music-library-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	c.JSON(http.StatusOK, artists)
}

func (handler ArtistHandler) GetArtist(c *gin.Context) {
	var artist Artist
	id := c.Param("id")
	handler.db.First(&artist, id)

	c.JSON(http.StatusOK, artist)
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
	handler.db.First(&artist, id)

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
	handler.db.First(&artist, id)

	handler.db.Delete(&artist)

	c.JSON(http.StatusOK, gin.H{"message": "artist deleted"})
}
