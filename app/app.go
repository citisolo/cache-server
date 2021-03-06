package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JSainsburyPLC/third-party-token-server/app/context"
	"github.com/JSainsburyPLC/third-party-token-server/app/handler"
	"github.com/JSainsburyPLC/third-party-token-server/app/logging"
	"github.com/JSainsburyPLC/third-party-token-server/config"

	"github.com/gorilla/mux"
)

type RequestHandlerFunction func(ctx *context.AppContext, w http.ResponseWriter, r *http.Request)
type NewMiddleWare func(next http.Handler) http.Handler

type App struct {
	Router  *mux.Router
	Context *context.AppContext
	Config  *config.Config
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware", r.Method)

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "https://localhost:3241")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			return
		}

		next.ServeHTTP(w, r)
		log.Println("Executing middleware again")
	})
}

func (app *App) Initialize(config *config.Config) {
	context := &context.AppContext{}
	context.Cache = &config.CacheManager
	context.Cache.InitDB()
	context.Logger = logging.NewAppLogger()
	app.Context = context
	app.Config = config
	app.Router = mux.NewRouter()

	app.Router.Use(app.NewRequestLoggingMiddleWare())
	app.setRouters()
}

func (app *App) setRouters() {
	app.Get("/token", app.NewHandler(handler.GetToken))
	app.Post("/token", app.NewHandler(handler.PostToken))
}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

func (app *App) NewHandler(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.Context, w, r)
	}
}

func (app *App) NewRequestLoggingMiddleWare() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			app.Context.Logger.Log(fmt.Sprintf("%v", r))
			next.ServeHTTP(w, r)
			app.Context.Logger.Log(fmt.Sprintf("%v", r.Response))
		})
	}
}

func (a *App) Run(host string) {
	a.Context.Logger.Log(fmt.Sprintf("Server listening at Port:%v", host))

	log.Fatal(http.ListenAndServeTLS(host, a.Config.CertFile, a.Config.KeyFile, a.Router))
	//log.Fatal(http.ListenAndServe(host, corsMiddleware(a.Router)))
}
