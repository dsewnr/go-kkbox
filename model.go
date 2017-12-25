package kkbox

import (
	"time"
)

// Summary page data
type Summary struct {
	Total int `json:"total"`
}

// Owner data
type Owner struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Image for music
type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

// Paging data
type Paging struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

// Artist struct
type Artist struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	URL    string  `json:"url"`
	Images []Image `json:"images"`
}

// Album data
type Album struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	URL                  string   `json:"url"`
	Explicitness         bool     `json:"explicitness"`
	AvailableTerritories []string `json:"available_territories"`
	ReleaseDate          string   `json:"release_date"`
	Images               []Image  `json:"images"`
	Artist               Artist   `json:"artist"`
}

// Track data
type Track struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	Duration             int      `json:"duration"`
	URL                  string   `json:"url"`
	TrackNumber          int      `json:"track_number"`
	Explicitness         bool     `json:"explicitness"`
	AvailableTerritories []string `json:"available_territories"`
	Album                Album    `json:"album"`
}

// GroupListData for song list
type GroupListData struct {
	Data []struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		Images      []Image   `json:"images"`
		UpdatedAt   time.Time `json:"updated_at"`
		Owner       Owner     `json:"owner"`
	} `json:"data"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// PlayListData to retrieve information of playlist
type PlayListData struct {
	Tracks struct {
		Data    []Track `json:"data"`
		Paging  Paging  `json:"paging"`
		Summary Summary `json:"summary"`
	} `json:"tracks"`
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Images      []Image   `json:"images"`
	UpdatedAt   time.Time `json:"updated_at"`
	Owner       Owner     `json:"owner"`
}

// TrackData List tracks of a chart playlist.
type TrackData struct {
	Data    []Track `json:"data"`
	Paging  Paging  `json:"paging"`
	Summary Summary `json:"summary"`
}

// Param for http get parameter
type Param struct {
	Territory string
	Page      int
	PerPage   int
}