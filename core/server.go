package core

import (
	"Tenancity/API/core/middleware"
	structs "Tenancity/API/core/structs"
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httpcoala"
	"github.com/go-chi/render"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var err error

type Instance structs.Instance

var Server *Instance

func NewInstance() *Instance {
	//TODO fetch the address from viper config.
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
		version := viper.GetString("version")
		//sourceUrl := "file:///" + path + version
		sourceUrl := "file://" + path + "/" + version

		connection := Server.DB.Credentials["main"]
		dbSource := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", connection.Username, connection.Password, connection.Host, connection.Port, connection.Name)

		m, err := migrate.New(sourceUrl, dbSource)

		if err != nil {
			logrus.Fatal(err)
		}

		e := m.Up()
		if e != nil {
			logrus.Warn(e)
		}
	}
}
