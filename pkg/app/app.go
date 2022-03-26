package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
}

func NewApp() (*App, error) {
	app := &App{
		router: mux.NewRouter().StrictSlash(true),
	}
	app.router.HandleFunc("/", app.HelloWorld)
	app.router.HandleFunc("/word", app.GetWord)
	return app, nil
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func (a *App) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		panic(fmt.Errorf("App.HelloWorld : %w", err))
	}
}

func (a *App) GetWord(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
