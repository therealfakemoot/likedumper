package main

import (
	"context"
	"encoding/json"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

func ListLikedTracks(conf Config) ([]SpotifyTrack, error) {
	tracks := make([]SpotifyTrack, 0)

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
		return tracks, fmt.Errorf("did not receive valid auth code: %w", err)
	}

	tok, err := oauthConf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		return tracks, fmt.Errorf("unable to verify PKCE: %w", err)
	}
	client := oauthConf.Client(ctx, tok)

	var nextURL string
	for {
		if nextURL == "" {
			nextURL = "https://api.spotify.com/v1/me/tracks"
		}
		resp, err := client.Get(nextURL)
		if err != nil {
			return tracks, fmt.Errorf("unable to fetch Tracks: %w", err)
		}

		var trackResponse SpotifyTracksResponse
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&trackResponse)
		if err != nil {
			return tracks, fmt.Errorf("couldn't decode Spotify response: %w", err)
		}

		nextURL = trackResponse.Next
		count := trackResponse.Total

		for _, item := range trackResponse.Items {
			tracks = append(tracks, item.Track)
			// fmt.Printf("%v - %v\n", item.Track.Artists[0].Name, item.Track.Name)
		}

		if len(tracks) == count {
			return tracks, nil
		}

	}
}
