package app

import (
	"net/http"

	"github.com/kitanoyoru/effective-mobile-task/internal/config"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/cache"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/events"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/store"
)

type App struct {
	cfg *config.Config

	store  *store.StoreSession
	cache  *cache.CacheSession
	events *events.EventBusSession

	server *http.Server
}
