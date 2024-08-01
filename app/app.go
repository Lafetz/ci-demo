package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

type App struct {
	port   int
	router *http.ServeMux
	logger *slog.Logger
}

func (a *App) routes() {
	a.router.HandleFunc("GET /", getHandler(a.logger))
	a.router.HandleFunc("POST /", postHandler(a.logger))
}

func (a *App) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(a.port)), a.router)
}

func NewApp(port int) *App {
	a := &App{
		port:   port,
		router: http.NewServeMux(),
	}
	a.routes()
	return a
}
