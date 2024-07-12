package handlers

import (
	"music-library-be/models"

	"github.com/jinzhu/gorm"
)

func getArtistNameFromId(artistId string, handler struct{ db *gorm.DB }) string {
	var artist models.Artist
	handler.db.First(&artist, "id = ?", artistId)
	return artist.Name
}

func getArtistNameFromAlbumId(albumId string, handler struct{ db *gorm.DB }) string {
	var album models.Album
	handler.db.First(&album, "id = ?", albumId)
	var artist models.Artist
	handler.db.First(&artist, "id = ?", album.ArtistID)
	return artist.Name
}

func getArtistNameFromSongId(songId string, handler struct{ db *gorm.DB }) string {
	var song models.Song
	handler.db.First(&song, "id = ?", songId)
	var album models.Album
	handler.db.First(&album, "id = ?", song.AlbumID)
	var artist models.Artist
	handler.db.First(&artist, "id = ?", album.ArtistID)
	return artist.Name
}

func getAlbumNameFromId(albumId string, handler struct{ db *gorm.DB }) string {
	var album models.Album
	handler.db.First(&album, "id = ?", albumId)
	return album.Title
}

func getAlbumNameFromSongId(songId string, handler struct{ db *gorm.DB }) string {
	var song models.Song
	handler.db.First(&song, "id = ?", songId)
	var album models.Album
	handler.db.First(&album, "id = ?", song.AlbumID)
	return album.Title
}
