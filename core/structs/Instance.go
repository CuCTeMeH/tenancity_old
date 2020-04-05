package core

import (
	"github.com/go-chi/chi"
	"net/http"
)

type Instance struct {
	AppName    string
	Debug      bool
	Version    string
	Addr       string
	DB         *DB
	Endpoints  *Routes
	Router     *chi.Mux
	HttpServer *http.Server
	Settings   *Setting
}
