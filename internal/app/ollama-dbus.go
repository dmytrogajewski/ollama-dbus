package app

import (
	"fmt"
	"time"

	l "github.com/dmytrogajewski/ollama-dbus/internal/logger"
	"github.com/dmytrogajewski/ollama-dbus/internal/ollama"
	"github.com/godbus/dbus/v5"
	"github.com/urfave/cli/v2"
)

func New() *cli.App {
	return &cli.App{
		Name:  "ollama-dbus",
		Usage: "",
		Action: func(ctx *cli.Context) error {
			debounceTime := 5 * time.Second
			sp := ollama.NewSearchProvider(debounceTime, "gemma2:2b", l.Logger)
			conn, err := dbus.SessionBus()

			if err != nil {
				return fmt.Errorf("Failed to connect to session bus: %w", err)
			}

			err = sp.Serve(conn)

			if err != nil {
				return fmt.Errorf("Failed to start search provider: %w", err)
			}

			l.Logger.Info("Search provider is running...")

			select {}
		},
		Description: "Ollama D-Bus implements GNOME search provider https://developer.gnome.org/documentation/tutorials/search-provider.html",
	}
}
