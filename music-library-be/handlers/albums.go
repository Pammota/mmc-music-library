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

type Album = models.Album

type AlbumHandler struct {
	db *gorm.DB
}

func CreateAlbumHandler(gorm *gorm.DB) AlbumHandler {
	return AlbumHandler{
		db: gorm,
	}
}

func (handler AlbumHandler) GetAlbums(c *gin.Context) {
	var albums []Album
	handler.db.Find(&albums)

	c.JSON(http.StatusOK, albums)
}

func (handler AlbumHandler) GetAlbum(c *gin.Context) {
	var album Album
	id := c.Param("id")
	handler.db.First(&album, id)

	c.JSON(http.StatusOK, album)
}

func (handler AlbumHandler) CreateAlbum(c *gin.Context) {
	var album Album
	album.ID = uuid.New().String()
	err := c.ShouldBindJSON(&album)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.db.Create(&album)

	c.JSON(http.StatusOK, album)
}

func (handler AlbumHandler) UpdateAlbum(c *gin.Context) {
	var album Album
	id := c.Param("id")
	handler.db.First(&album, id)

	err := c.ShouldBindJSON(&album)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.db.Update(&album)

	c.JSON(http.StatusOK, album)
}

func (handler AlbumHandler) DeleteAlbum(c *gin.Context) {
	var album Album
	id := c.Param("id")
	handler.db.First(&album, id)

	handler.db.Delete(&album)

	c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
}
