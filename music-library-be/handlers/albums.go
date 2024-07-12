package handlers

import (
	// "encoding/json"

	"music-library-be/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/exp/rand"
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

	var returnAlbums []models.ReturnJSONAlbum

	for i := range albums {
		returnAlbums = append(returnAlbums, models.ReturnJSONAlbum{
			ID:          albums[i].ID,
			Title:       albums[i].Title,
			Description: albums[i].Description,
			ArtistName:  getArtistNameFromId(albums[i].ArtistID, handler),
		})
	}

	c.JSON(http.StatusOK, returnAlbums)
}

func (handler AlbumHandler) GetAlbum(c *gin.Context) {
	var album Album
	id := c.Param("id")
	handler.db.First(&album, "id = ?", id)

	var returnAlbum models.ReturnJSONAlbum

	returnAlbum.ID = album.ID
	returnAlbum.Title = album.Title
	returnAlbum.Description = album.Description
	returnAlbum.ArtistName = getArtistNameFromId(album.ArtistID, handler)

	c.JSON(http.StatusOK, returnAlbum)
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
	handler.db.First(&album, "id = ?", id)

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
	handler.db.First(&album, "id = ?", id)

	handler.db.Delete(&album)

	c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
}

func (handler AlbumHandler) GetAlbumsByArtist(c *gin.Context) {
	var albums []Album
	artistId := c.Param("artistId")

	if artistId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "artistId is required"})
		return
	}

	handler.db.Where("artist_id = ?", artistId).Find(&albums)

	var returnAlbums []models.ReturnJSONAlbum

	artistName := getArtistNameFromId(artistId, handler)

	for _, album := range albums {
		returnAlbums = append(returnAlbums, models.ReturnJSONAlbum{
			ID:          album.ID,
			Title:       album.Title,
			Description: album.Description,
			ArtistName:  artistName,
		})
	}

	c.JSON(http.StatusOK, returnAlbums)
}

func (handler AlbumHandler) GetRandomAlbums(c *gin.Context) {
	var albums []Album
	handler.db.Find(&albums)

	var returnAlbums []models.ReturnJSONAlbum

	for i := range albums {
		returnAlbums = append(returnAlbums, models.ReturnJSONAlbum{
			ID:          albums[i].ID,
			Title:       albums[i].Title,
			Description: albums[i].Description,
			ArtistName:  getArtistNameFromId(albums[i].ArtistID, handler),
		})

		rand.Seed(uint64(time.Now().UnixNano()))
		rand.Shuffle(len(returnAlbums), func(i, j int) { returnAlbums[i], returnAlbums[j] = returnAlbums[j], returnAlbums[i] })

		c.JSON(http.StatusOK, returnAlbums)
	}
}
