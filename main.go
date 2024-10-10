package main

import (
	"fmt"
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

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	conf, err := LoadConfigFile("config.toml")
	if err != nil {
		logger.With("error", err).Error("error loading config file")
	}

	tracks, err := ListLikedTracks(conf)
	if err != nil {
		logger.With("error", err).Error("error fetching tracks")
	}
	fmt.Printf("track count: %d\n", len(tracks))
}
