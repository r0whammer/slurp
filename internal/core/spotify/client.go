package spotify

import (
	"context"
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"
	// "net/url"
	"os"
	"path/filepath"
	// "strconv"
	// "strings"
	"time"
)

type Client struct {
	http  *http.Client
	token func(ctx context.Context) (string, error)
}

type savedTracksResponse struct {
	Items []struct {
		AddedAt string `json:"added_at"`
		Track   struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			Album struct {
				Name   string `json:"name"`
				Images []struct {
					URL string `json:"url"`
				} `json:"images"`
			} `json:"album"`
		} `json:"track"`
	} `json:"items"`
}

func NewClient(tokenFn func(context.Context) (string, error)) *Client {
	return &Client{
		http:  &http.Client{Timeout: 20 * time.Second},
		token: tokenFn,
	}
}

func (c *Client) GetSavedTracks(ctx context.Context, limit, offset int) ([]SavedTrackRow, error) {
	// dev-only: read from file if set
	if p := os.Getenv("SPOTIFY_TRACKS_JSON"); p != "" {
		return c.getSavedTracksFromFile(p, limit, offset)
	}

	// TODO: real spotify http call later
	return nil, fmt.Errorf("no spotify auth yet (set SPOTIFY_TRACKS_JSON=/abs/or/rel/path/to/tracks.json)")

	// if limit <= 0 {
	// 	limit = 20
	// }
	// if limit > 50 {
	// 	limit = 50 // spotify max for this endpoint :contentReference[oaicite:2]{index=2}
	// }
	// if offset < 0 {
	// 	offset = 0
	// }

	// tok, err := c.token(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// if strings.TrimSpace(tok) == "" {
	// 	return nil, errors.New("spotify access token is empty")
	// }

	// u, _ := url.Parse("https://api.spotify.com/v1/me/tracks")
	// q := u.Query()
	// q.Set("limit", strconv.Itoa(limit))
	// q.Set("offset", strconv.Itoa(offset))
	// u.RawQuery = q.Encode()

	// req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	// req.Header.Set("Authorization", "Bearer "+tok)

	// res, err := c.http.Do(req)
	// if err != nil {
	// 	return nil, err
	// }
	// defer res.Body.Close()

	// if res.StatusCode != 200 {
	// 	return nil, fmt.Errorf("spotify /me/tracks failed: %s", res.Status)
	// }

	// // minimal response shape we care about
	// var payload struct {
	// 	Items []struct {
	// 		AddedAt string `json:"added_at"`
	// 		Track   struct {
	// 			ID      string `json:"id"`
	// 			Name    string `json:"name"`
	// 			Artists []struct {
	// 				Name string `json:"name"`
	// 			} `json:"artists"`
	// 			Album struct {
	// 				Name   string `json:"name"`
	// 				Images []struct {
	// 					URL string `json:"url"`
	// 				} `json:"images"`
	// 			} `json:"album"`
	// 		} `json:"track"`
	// 	} `json:"items"`
	// }

	// if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
	// 	return nil, err
	// }

	// out := make([]SavedTrackRow, 0, len(payload.Items))
	// for _, it := range payload.Items {
	// 	artist := ""
	// 	if len(it.Track.Artists) > 0 {
	// 		artist = it.Track.Artists[0].Name
	// 	}
	// 	cover := ""
	// 	if len(it.Track.Album.Images) > 0 {
	// 		cover = it.Track.Album.Images[0].URL
	// 	}
	// 	out = append(out, SavedTrackRow{
	// 		ID:          it.Track.ID,
	// 		CoverURL:    cover,
	// 		TrackTitle:  it.Track.Name,
	// 		ArtistTitle: artist,
	// 		AlbumTitle:  it.Track.Album.Name,
	// 		AddedAt:     it.AddedAt,
	// 	})
	// }

	// return out, nil
}

func (c *Client) getSavedTracksFromFile(path string, limit, offset int) ([]SavedTrackRow, error) {
	if limit <= 0 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	// allow relative paths from current working dir
	path = filepath.Clean(path)

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read tracks json %q: %w", path, err)
	}

	var payload savedTracksResponse
	if err := json.Unmarshal(b, &payload); err != nil {
		return nil, fmt.Errorf("unmarshal tracks json %q: %w", path, err)
	}

	// apply offset/limit against the file contents
	items := payload.Items
	if offset >= len(items) {
		return []SavedTrackRow{}, nil
	}
	items = items[offset:]
	if limit < len(items) {
		items = items[:limit]
	}

	out := make([]SavedTrackRow, 0, len(items))
	for _, it := range items {
		artist := ""
		if len(it.Track.Artists) > 0 {
			artist = it.Track.Artists[0].Name
		}
		cover := ""
		if len(it.Track.Album.Images) > 0 {
			cover = it.Track.Album.Images[0].URL
		}

		out = append(out, SavedTrackRow{
			ID:          it.Track.ID,
			CoverURL:    cover,
			TrackTitle:  it.Track.Name,
			ArtistTitle: artist,
			AlbumTitle:  it.Track.Album.Name,
			AddedAt:     it.AddedAt,
		})
	}

	return out, nil
}
