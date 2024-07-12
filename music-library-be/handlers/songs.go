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

	var returnSongs []models.ReturnJSONSong

	for i := range songs {
		returnSongs = append(returnSongs, models.ReturnJSONSong{
			ID:         songs[i].ID,
			Title:      songs[i].Title,
			Length:     songs[i].Length,
			AlbumTitle: getAlbumNameFromId(songs[i].AlbumID, handler),
			ArtistName: getArtistNameFromSongId(songs[i].ID, handler),
		})
	}

	c.JSON(http.StatusOK, returnSongs)
}

func (handler SongHandler) GetSong(c *gin.Context) {
	var song Song
	id := c.Param("id")
	handler.db.First(&song, "id = ?", id)

	var returnSong models.ReturnJSONSong

	returnSong.ID = song.ID
	returnSong.Title = song.Title
	returnSong.Length = song.Length
	returnSong.AlbumTitle = getAlbumNameFromId(song.AlbumID, handler)
	returnSong.ArtistName = getArtistNameFromSongId(song.ID, handler)

	c.JSON(http.StatusOK, returnSong)
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
	handler.db.First(&song, "id = ?", id)

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
	handler.db.First(&song, "id = ?", id)

	handler.db.Delete(&song)

	c.JSON(http.StatusOK, gin.H{"message": "song deleted"})
}

func (handler SongHandler) GetSongsByAlbum(c *gin.Context) {
	var songs []Song
	albumId := c.Param("albumId")

	handler.db.Where("album_id = ?", albumId).Find(&songs)

	var returnSongs []models.ReturnJSONSong

	artistName := getArtistNameFromAlbumId(albumId, handler)
	albumName := getAlbumNameFromSongId(albumId, handler)

	for _, song := range songs {
		returnSongs = append(returnSongs, models.ReturnJSONSong{
			ID:         song.ID,
			Title:      song.Title,
			Length:     song.Length,
			AlbumTitle: albumName,
			ArtistName: artistName,
		})
	}

	c.JSON(http.StatusOK, returnSongs)
}

func (handler SongHandler) GetSongsByArtist(c *gin.Context) {
	var songs []Song
	artistId := c.Param("artistId")

	handler.db.Where("artist_id = ?", artistId).Find(&songs)

	var returnSongs []models.ReturnJSONSong

	artistName := getArtistNameFromId(artistId, handler)
	albumName := getAlbumNameFromId(artistId, handler)

	for i := range songs {
		returnSongs = append(returnSongs, models.ReturnJSONSong{
			ID:         songs[i].ID,
			Title:      songs[i].Title,
			Length:     songs[i].Length,
			AlbumTitle: albumName,
			ArtistName: artistName,
		})
	}

	c.JSON(http.StatusOK, returnSongs)
}

func (handler SongHandler) GetRandomSongs(c *gin.Context) {
	var songs []Song
	handler.db.Find(&songs)

	rand.Seed(uint64(time.Now().UnixNano()))
	rand.Shuffle(len(songs), func(i, j int) { songs[i], songs[j] = songs[j], songs[i] })

	c.JSON(http.StatusOK, songs)
}
