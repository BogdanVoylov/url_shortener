package main

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)


type App struct {
	Router *mux.Router
	DB     *sqlx.DB
	config *Config
}

func NewApp(config *Config) *App {
	db, err := sqlx.Connect("postgres", config.DbConnectionString())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	app := &App{r, db, config }
	log.Printf("Successfully connected to %s\n", config.DB.Name)
	return app
}


func (this *App) HandleFunc(url string, h func(http.ResponseWriter, *http.Request)) *mux.Route {
	return this.Router.HandleFunc(url,h)
}

func (this *App) Run() {
	log.Fatal(http.ListenAndServe(this.config.Server.Address, this.Router))
}

func (this *App) Db() *sqlx.DB {
	return this.DB
}
