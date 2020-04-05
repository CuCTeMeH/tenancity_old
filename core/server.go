package core

import (
	"Tenancity/API/core/middleware"
	structs "Tenancity/API/core/structs"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httpcoala"
	"github.com/go-chi/render"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var err error

type Instance structs.Instance

var Server *Instance

func NewInstance() *Instance {
	s := &Instance{
		Addr: ":3333",
		// just in case you need some setup here
	}
	s = s.InitConfig()
	return s
}

func (i *Instance) Start() *Instance {
	// Startup all dependencies
	// I usually panic if any essential like the DB fails
	// e.g. due to wrong configurations

	// Startup the http Server in a way that
	// we can gracefully shut it down again
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		//logrus.Error("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}

	if err := chi.Walk(i.Router, walkFunc); err != nil {
		logrus.Panic("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	i.HttpServer = &http.Server{Addr: i.Addr, Handler: i.Router}
	err = i.HttpServer.ListenAndServe() // Blocks!
	if err != http.ErrServerClosed {
		logrus.WithError(err).Error("Http Server stopped unexpected")
		i.Shutdown()
	} else {
		logrus.WithError(err).Info("Http Server stopped")
	}

	return i
}

func (i *Instance) Shutdown() {
	if i.HttpServer != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := i.HttpServer.Shutdown(ctx)
		if err != nil {
			logrus.WithError(err).Error("Failed to shutdown http server gracefully")
		} else {
			i.HttpServer = nil
		}
	}
}

func (i *Instance) AddRoute(path string, handler http.Handler) {
	i.Endpoints.Handlers[path] = handler
}

func (i *Instance) InitRoutes() {
	if i.Endpoints == nil {
		i.Endpoints = &structs.Routes{}
		i.Endpoints.Handlers = make(map[string]http.Handler)
	}
}

func (i *Instance) GetRoutes() *Instance {
	if i.Endpoints == nil {
		i.Endpoints.Handlers = make(map[string]http.Handler)
	}
	return i
}

func (i *Instance) InitRouter() {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.Logger,
		middleware.RealIP,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.URLFormat,
		middleware.Throttle(5),
		httpcoala.Route("*"),
		core.Locale(),
	)

	router.Route("/v1", func(r chi.Router) {
		for k, v := range i.Endpoints.Handlers {
			r.Mount(k, v)
		}
	})

	i.Router = router
}

func AutoMigrateModules(path string) {
	if viper.GetBool("autoMigrate") == true {
		version := "1.0.0"
		//set db configurable here.
		m, err := migrate.New(
			path+version,
			"mysql://root:@tcp(localhost:3306)/Tenancity")
		if err != nil {
			logrus.Fatal(err)
		}
		m.Up()
	}
}
