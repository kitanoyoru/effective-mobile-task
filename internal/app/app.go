package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	httpApi "github.com/kitanoyoru/effective-mobile-task/internal/api/http/v0"
	"github.com/kitanoyoru/effective-mobile-task/internal/config"
	"github.com/kitanoyoru/effective-mobile-task/internal/service"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/cache"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/crawler"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/events"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/store"
	log "github.com/sirupsen/logrus"
)

const (
	defaultAppShutdownTimeout = 5 * time.Second

	defaultHttpServerReadTimeout  = 10 * time.Second
	defaultHttpServerWriteTimeout = 10 * time.Second

	defaultHttpServerMaxHeaderBytes = 1 << 20
)

type App struct {
	cfg *config.Config

	service *service.Service

	store  *store.StoreSession
	cache  *cache.CacheSession
	events *events.EventBusSession

	crawler *crawler.CrawlerSession

	httpServer *http.Server
}

func NewApp(cfg *config.Config) (*App, error) {
	app := App{}

	app.cfg = cfg

	app.events = events.NewEventBusSession()

	level, err := log.ParseLevel(app.cfg.Log.LogLevel)
	if err != nil {
		return nil, err
	}
	log.SetLevel(level)

	storeSession, err := store.NewStoreSession(&app.cfg.Database, app.events)
	if err != nil {
		return nil, err
	}
	app.store = storeSession

	cacheSession, err := cache.NewCacheSession(&app.cfg.Cache, app.events)
	if err != nil {
		return nil, err
	}
	app.cache = cacheSession

	app.crawler = crawler.NewCrawlerSession(&app.cfg.Crawler, app.events, app.store)

	app.service = service.NewService(app.store, app.cache)

	return &app, nil
}

func (app *App) Run() error {
	router := httpApi.NewHTTPApi(app.service)

	addr := fmt.Sprintf("%s:%s", app.cfg.Server.Host, app.cfg.Server.Port)

	app.httpServer = &http.Server{
		Addr:           addr,
		Handler:        router.GetHTTPRouter(),
		ReadTimeout:    defaultHttpServerReadTimeout,
		WriteTimeout:   defaultHttpServerWriteTimeout,
		MaxHeaderBytes: defaultHttpServerMaxHeaderBytes,
	}

	go func() {
		log.Info("Starting HTTP server...")
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), defaultAppShutdownTimeout)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}
