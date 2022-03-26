package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
}

func NewApp() (*App, error) {
	return &App{
		router: mux.NewRouter().StrictSlash(true),
	}, nil
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
