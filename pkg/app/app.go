package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Sqlite driver based on GGO
// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details

type App struct {
	router *mux.Router
	db     *gorm.DB
}

func NewApp() (*App, error) {
	db, err := gorm.Open(sqlite.Open("word.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("app.NewApp : %w", err)
	}
	app := &App{
		router: mux.NewRouter().StrictSlash(true),
		db:     db,
	}
	app.router.HandleFunc("/", app.HelloWorld)
	app.router.HandleFunc("/word", app.ListWord).Methods(http.MethodGet)
	app.router.HandleFunc("/word", app.ListWord).Methods(http.MethodPost)
	return app, nil
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(r.(error).Error()))
		}
	}()
	a.router.ServeHTTP(w, r)
}

func (a *App) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		panic(fmt.Errorf("App.HelloWorld : %w", err))
	}
}

func (a *App) ListWord(w http.ResponseWriter, r *http.Request) {
	wordList, err := ListWord(r.Context(), a.db)
	if err != nil {
		panic(fmt.Errorf("App.GetWord Fetching Data : %w", err))
	}
	words := make([]string, len(wordList))
	for i, v := range wordList {
		words[i] = v.Word
	}
	err = json.NewEncoder(w).Encode(words)
	if err != nil {
		panic(fmt.Errorf("App.GetWord Writing response: %w", err))
	}
	w.WriteHeader(http.StatusOK)
}

func (a *App) CreateWord(w http.ResponseWriter, r *http.Request) {
	word := &Words{}
	err := json.NewDecoder(r.Body).Decode(word)
	if err != nil {
		panic(fmt.Errorf("App.CreateWord reading request: %w", err))
	}
	err = word.Create(r.Context(), a.db)
	if err != nil {
		panic(fmt.Errorf("App.CreateWord creating record: %w", err))
	}
}
