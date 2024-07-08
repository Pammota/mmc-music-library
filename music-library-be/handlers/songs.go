package handlers

import (
	// "encoding/json"

	"music-library-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Song = models.Song

type SongHandler struct {
	db *gorm.DB
}

func CreateSongHandler(gorm *gorm.DB) SongHandler {
	return SongHandler{
		db: gorm,
	}
}

func (handler SongHandler) GetSongs(c *gin.Context) {
	var songs []Song
	handler.db.Find(&songs)

	c.JSON(http.StatusOK, songs)
}

func (handler SongHandler) GetSong(c *gin.Context) {
	var song Song
	id := c.Param("id")
	handler.db.First(&song, id)

	c.JSON(http.StatusOK, song)
}

func (handler SongHandler) CreateSong(c *gin.Context) {
	var song Song
	song.ID = uuid.New().String()
	err := c.ShouldBindJSON(&song)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.db.Create(&song)

	c.JSON(201, song)
}

func (handler SongHandler) UpdateSong(c *gin.Context) {
	var song Song
	id := c.Param("id")
	handler.db.First(&song, id)

	err := c.ShouldBindJSON(&song)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.db.Update(&song)

	c.JSON(http.StatusOK, song)
}

func (handler SongHandler) DeleteSong(c *gin.Context) {
	var song Song
	id := c.Param("id")
	handler.db.First(&song, id)

	handler.db.Delete(&song)

	c.JSON(http.StatusOK, gin.H{"message": "song deleted"})
}
