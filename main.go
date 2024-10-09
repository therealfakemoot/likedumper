package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/pelletier/go-toml/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
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

	ctx := context.Background()
	oauthConf := &oauth2.Config{
		ClientID:     conf.ID,
		ClientSecret: conf.Secret,
		Scopes:       []string{"user-library-read"},
		Endpoint:     spotify.Endpoint,
		RedirectURL:  "https://likedumper.ndumas.com/callback",
	}

	verifier := oauth2.GenerateVerifier()

	url := oauthConf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		logger.Error("did not receive valid auth code")
		return
	}

	tok, err := oauthConf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		logger.Error("unable to verify PKCE")
		return
	}
	client := oauthConf.Client(ctx, tok)

	// here is where i need to start handling pagination
	resp, err := client.Get("https://api.spotify.com/v1/me/tracks")
	if err != nil {
		logger.Error("unable to fetch Tracks")
		return
	}

	var tracks SpotifyTracksResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&tracks)
	if err != nil {
		logger.Error("couldn't decode Spotify response")
		return
	}

	for _, item := range tracks.Items {
		fmt.Printf("%v - %v\n", item.Track.Artists[0].Name, item.Track.Name)
	}

}
