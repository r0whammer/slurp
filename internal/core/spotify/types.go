package spotify

type SavedTrackRow struct {
	ID          string `json:"id"`
	CoverURL    string `json:"coverUrl"`
	TrackTitle  string `json:"trackTitle"`
	ArtistTitle string `json:"artistTitle"`
	AlbumTitle  string `json:"albumTitle"`
	AddedAt     string `json:"addedAt"`
}
