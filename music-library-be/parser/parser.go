package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"music-library-be/models"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// existsArtist checks if an artist exists in the database
func existsArtist(db *gorm.DB, name string) bool {
	var count int64
	db.Model(&models.Artist{}).Where("name = ?", name).Count(&count)
	return count > 0
}

// existsAlbum checks if an album exists in the database
func existsAlbum(db *gorm.DB, title string, artistID string) bool {
	var count int64
	db.Model(&models.Album{}).Where("title = ? AND artist_id = ?", title, artistID).Count(&count)
	return count > 0
}

// existsSong checks if a song exists in the database
func existsSong(db *gorm.DB, title string, albumID string) bool {
	var count int64
	db.Model(&models.Song{}).Where("title = ? AND album_id = ?", title, albumID).Count(&count)
	return count > 0
}

func printArtist(artist models.Artist, albums []models.Album, songsMap map[string][]models.Song) {
	fmt.Printf("Artist: %s\n", artist.Name)
	fmt.Printf("\tID: %s\n", artist.ID)
	for _, album := range albums {
		fmt.Printf("\tAlbum: %s\n", album.Title)
		fmt.Printf("\t\tID: %s\n", album.ID)
		fmt.Printf("\t\tDescription: %s\n", album.Description)
		for _, song := range songsMap[album.ID] {
			fmt.Printf("\t\tSong: %s\n", song.Title)
			fmt.Printf("\t\t\tID: %s\n", song.ID)
			fmt.Printf("\t\t\tLength: %s\n", song.Length)
		}
	}
}

// ParseJSON reads and parses a JSON file and inserts the data into the database
func ParseJSON(inputFile string, db *gorm.DB) {
	var jsonArtists []models.JSONArtist

	jsonFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &jsonArtists)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Build data structures
	var newArtists []models.Artist
	var newAlbums []models.Album
	var newSongs []models.Song

	for _, jsonArtist := range jsonArtists {
		// Check if the artist exists
		if existsArtist(db, jsonArtist.Name) {
			log.Printf("Artist already exists: %s", jsonArtist.Name)
			continue
		}

		artistID := uuid.New().String()
		artist := models.Artist{
			ID:   artistID,
			Name: jsonArtist.Name,
		}
		log.Printf("Creating artist: %s with ID: %s", artist.Name, artist.ID)

		var albums []models.Album
		songsMap := make(map[string][]models.Song)

		// Process albums
		for _, jsonAlbum := range jsonArtist.Albums {
			// Check if the album exists
			if existsAlbum(db, jsonAlbum.Title, artist.ID) {
				log.Printf("Album already exists: %s for artist: %s", jsonAlbum.Title, artist.Name)
				continue
			}

			albumID := uuid.New().String()
			album := models.Album{
				ID:          albumID,
				Title:       jsonAlbum.Title,
				Description: jsonAlbum.Description,
				ArtistID:    artist.ID,
			}
			log.Printf("Creating album: %s with ID: %s for artist ID: %s", album.Title, album.ID, artist.ID)
			albums = append(albums, album)

			// Process songs
			for _, jsonSong := range jsonAlbum.Songs {
				// Check if the song exists
				if existsSong(db, jsonSong.Title, album.ID) {
					log.Printf("Song already exists: %s for album: %s", jsonSong.Title, album.Title)
					continue
				}

				songID := uuid.New().String()
				song := models.Song{
					ID:      songID,
					Title:   jsonSong.Title,
					Length:  jsonSong.Length,
					AlbumID: album.ID,
				}
				log.Printf("Creating song: %s with ID: %s for album ID: %s", song.Title, song.ID, album.ID)
				songsMap[album.ID] = append(songsMap[album.ID], song)
				newSongs = append(newSongs, song)
			}

			newAlbums = append(newAlbums, album)
		}

		printArtist(artist, albums, songsMap)
		newArtists = append(newArtists, artist)
	}

	// Send create requests to the database
	for _, artist := range newArtists {
		db.Create(&artist)
	}

	for _, album := range newAlbums {
		db.Create(&album)
	}

	for _, song := range newSongs {
		db.Create(&song)
	}
}
