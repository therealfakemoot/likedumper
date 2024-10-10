package main

import (
	"time"
)

type SpotifyTrackAlbum struct {
	AlbumType    string `json:"album_type"`
	TotalTracks  int    `json:"total_tracks"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	Name                 string `json:"name"`
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	Type                 string `json:"type"`
	URI                  string `json:"uri"`
	Artists              []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"artists"`
	IsPlayable bool `json:"is_playable"`
}

type SpotifyTrackArtist struct {
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href string `json:"href"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type SpotifyTrack struct {
	Album       SpotifyTrackAlbum    `json:"album"`
	Artists     []SpotifyTrackArtist `json:"artists"`
	DiscNumber  int                  `json:"disc_number"`
	DurationMs  int                  `json:"duration_ms"`
	Explicit    bool                 `json:"explicit"`
	ExternalIds struct {
		Isrc string `json:"isrc"`
	} `json:"external_ids"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href        string `json:"href"`
	ID          string `json:"id"`
	IsPlayable  bool   `json:"is_playable"`
	Name        string `json:"name"`
	Popularity  int    `json:"popularity"`
	PreviewURL  string `json:"preview_url"`
	TrackNumber int    `json:"track_number"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
	IsLocal     bool   `json:"is_local"`
}

type SpotifyTracksItem struct {
	AddedAt time.Time    `json:"added_at"`
	Track   SpotifyTrack `json:"track"`
}

type SpotifyTracksResponse struct {
	Href     string              `json:"href"`
	Limit    int                 `json:"limit"`
	Next     string              `json:"next"`
	Offset   int                 `json:"offset"`
	Previous any                 `json:"previous"`
	Total    int                 `json:"total"`
	Items    []SpotifyTracksItem `json:"items"`
}
