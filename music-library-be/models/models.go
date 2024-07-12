package models

// DB models

type Artist struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Name string `gorm:"type:varchar(100);not null"`
	// Albums []Album `gorm:"foreignKey:ArtistID;references:ArtistID"`
}

type Album struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Title       string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text"`
	ArtistID    string `gorm:"type:uuid"`
	// Songs       []Song `gorm:"foreignKey:AlbumID;references:AlbumID"`
}

type Song struct {
	ID      string `gorm:"type:uuid;primaryKey"`
	Title   string `gorm:"type:varchar(100);not null"`
	Length  string `gorm:"type:varchar(50);not null"`
	AlbumID string `gorm:"type:uuid"`
}

// JSON models

type JSONArtist struct {
	Name   string      `json:"name"`
	Albums []JSONAlbum `json:"albums"`
}

type JSONAlbum struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Songs       []JSONSong `json:"songs"`
}

type JSONSong struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Length string `json:"length"`
}

// JSOM return models

type ReturnJSONArtist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ReturnJSONAlbum struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ArtistName  string `json:"artistName"`
}

type ReturnJSONSong struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Length     string `json:"length"`
	AlbumTitle string `json:"albumTitle"`
	ArtistName string `json:"artistName"`
}
