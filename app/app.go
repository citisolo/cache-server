package app

import (
	"log"
	"net/http"

	"github.com/JSainsburyPLC/third-party-token-server/app/handler"
	"github.com/JSainsburyPLC/third-party-token-server/config"
	"github.com/JSainsburyPLC/third-party-token-server/db"
	"github.com/gorilla/mux"
)

type RequestHandlerFunction func(db db.Cache, w http.ResponseWriter, r *http.Request)

type App struct {
	Router *mux.Router
	Cache db.Cache
}

func (app *App) Initialize(config *config.Config) {
	app.Cache = &config.CacheManager
	app.Cache.InitDB()
	app.Router = mux.NewRouter()
	app.setRouters()
}

func (app *App) setRouters(){
	app.Get("/token", app.HandleRequest(handler.GetToken))
	app.Post("/token",app.HandleRequest(handler.PostToken))
}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

func (app *App) HandleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.Cache, w, r)
	}
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
