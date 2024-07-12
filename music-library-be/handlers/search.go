package handlers

import (
	"music-library-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type SearchHandler struct {
	db *gorm.DB
}

func CreateSearchHandler(gorm *gorm.DB) SearchHandler {
	return SearchHandler{
		db: gorm,
	}
}

func (handler SearchHandler) Search(c *gin.Context) {
	searchString := c.Param("question")
	searchPattern := "%" + searchString + "%"

	var songs []models.Song
	var returnSongs []models.ReturnJSONSong
	handler.db.Where("title ILIKE ?", searchPattern).Find(&songs)

	for _, song := range songs {
		returnSongs = append(returnSongs, models.ReturnJSONSong{
			ID:         song.ID,
			Title:      song.Title,
			Length:     song.Length,
			AlbumTitle: getAlbumNameFromId(song.AlbumID, handler),
			ArtistName: getArtistNameFromSongId(song.ID, handler),
		})
	}

	var albums []models.Album
	var returnAlbums []models.ReturnJSONAlbum
	handler.db.Where("title ILIKE ? OR description ILIKE ?", searchPattern, searchPattern).Find(&albums)

	for _, album := range albums {
		returnAlbums = append(returnAlbums, models.ReturnJSONAlbum{
			ID:          album.ID,
			Title:       album.Title,
			Description: album.Description,
			ArtistName:  getArtistNameFromAlbumId(album.ID, handler),
		})
	}

	var artists []models.Artist
	var returnArtists []models.ReturnJSONArtist
	handler.db.Where("name ILIKE ?", searchPattern).Find(&artists)

	for i := range artists {
		returnArtists = append(returnArtists, models.ReturnJSONArtist{
			ID:   artists[i].ID,
			Name: artists[i].Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"songs":   returnSongs,
		"albums":  returnAlbums,
		"artists": returnArtists,
	})
}
