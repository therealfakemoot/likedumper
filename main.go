package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	ID     string `toml:"SPOTIFY_APP_ID"`
	Secret string `toml:"SPOTIFY_APP_TOKEN"`
}

func LoadConfigFile(fn string) (Config, error) {
	var c Config
	f, err := os.Open(fn)
	if err != nil {
		return c, fmt.Errorf("couldn't open config file: %w", err)
	}
	dec := toml.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		return c, fmt.Errorf("error parsing config file: %w", err)
	}

	return c, nil
}

func WriteCSV(w io.Writer, tracks []SpotifyTrack) error {
	enc := csv.NewWriter(w)
	defer enc.Flush()

	enc.Write([]string{"Title", "Album", "Artist"})
	for _, track := range tracks {
		enc.Write([]string{
			track.Name,
			track.Album.Name,
			track.Artists[0].Name,
		})
	}

	return nil
}

func main() {
	var outFile string

	flag.StringVar(&outFile, "dest", "tracks.csv", "Destination CSV file")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	conf, err := LoadConfigFile("config.toml")
	if err != nil {
		logger.With("error", err).Error("error loading config file")
	}

	tracks, err := ListLikedTracks(conf)
	if err != nil {
		logger.With("error", err).Error("error fetching tracks")
	}

	f, err := os.Create(outFile)
	if err != nil {
		logger.With("error", err).Error("could not create destination file")
		os.Exit(1)
	}

	err = WriteCSV(f, tracks)
	if err != nil {
		logger.With("error", err).Error("could not write destination file")
		os.Exit(2)
	}
}
