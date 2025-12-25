package main

import (
	"context"
	"embed"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/r0whammer/slurp/internal/cli/commands"
	"github.com/r0whammer/slurp/internal/core/spotify"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	if len(os.Args) == 1 {
		// Create an instance of the app structure
		app := NewApp()
		// Create application with options
		err := wails.Run(&options.App{
			Title:  "slurp",
			Width:  1024,
			Height: 768,
			AssetServer: &assetserver.Options{
				Assets: assets,
			},
			BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
			OnStartup:        app.startup,
			Bind: []any{
				app,
			},
		})
		if err != nil {
			println("Error:", err.Error())
		}
	} else {
		rootCmd := commands.NewRootCmd()
		if err := fang.Execute(context.Background(), rootCmd); err != nil {
			os.Exit(1)
		}
	}
}

// App struct
type App struct {
	ctx     context.Context
	spotify *spotify.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	tokenFn := func(ctx context.Context) (string, error) {
		return os.Getenv("SPOTIFY_ACCESS_TOKEN"), nil
	}
	return &App{
		spotify: spotify.NewClient(tokenFn),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// func (a *App) GetSavedTracks(limit int, offset int) ([]spotify.SavedTrackRow, error) {
// 	return a.spotify.GetSavedTracks(a.ctx, limit, offset)
// }

func (a *App) GetSavedTracks(limit int, offset int) ([]spotify.SavedTrackRow, error) {
	return a.spotify.GetSavedTracks(a.ctx, limit, offset)
}
