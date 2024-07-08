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
	Title  string `json:"title"`
	Length string `json:"length"`
}
