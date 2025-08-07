package main

import (
	"context"
	"raw-photo-manager/backend"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	rt "runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetOS() string {
	return rt.GOOS
}

func (a *App) ListDir(path string) []string {
	return backend.ListFiles(path)
}

func (a *App) DeleteFiles(files []string) error {
	return backend.DeleteFiles(files)
}

func (a *App) PickFolder() (string, error) {
	return runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a folder",
	})
}
