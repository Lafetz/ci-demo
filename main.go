package main

import (
	"log/slog"
	"os"

	"actions/app"
)

// var version string

func main() {
	a := app.NewApp(8080)
	slog.Info("server running...")
	err := a.Run()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}