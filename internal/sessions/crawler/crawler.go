package crawler

import (
	"context"
	"encoding/json"
	"io"
	"net/url"
	"time"

	"github.com/gojek/heimdall/hystrix"
	"github.com/kitanoyoru/effective-mobile-task/internal/config"
	"github.com/kitanoyoru/effective-mobile-task/internal/requests"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/events"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/store"
	"github.com/kitanoyoru/effective-mobile-task/pkg/pool"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

const (
	defaultSourceRetryCount = 5
	defaultSourceTimeout    = 5 * time.Second
)

const (
	agePersonSource         = "https://api.agify.io"
	genderPersonSource      = "https://api.genderize.io"
	nationalityPersonSource = "https://api.nationalize.io"
)

var sources = []string{
	agePersonSource,
	genderPersonSource,
	nationalityPersonSource,
}

type CrawlerSession struct {
	client *hystrix.Client
	pool   *pool.Pool
	store  *store.StoreSession

	bus                 *events.EventBusSession
	personPostCtxCancel context.CancelFunc
}

func NewCrawlerSession(cfg *config.CrawlerConfig, bus *events.EventBusSession, store *store.StoreSession) *CrawlerSession {
	client := hystrix.NewClient(
		hystrix.WithHTTPTimeout(defaultSourceTimeout),
		hystrix.WithHystrixTimeout(3*defaultSourceTimeout),
	)
	pool := pool.NewPool(cfg.Size)

	personPostCtx, personPostCtxCancel := context.WithCancel(context.Background())

	c := &CrawlerSession{
		client,
		pool,
		store,
		bus,
		personPostCtxCancel,
	}

	go bus.AsyncConsumeEvents(personPostCtx, events.PersonPostEventTopic, c.onPersonPostHandler)

	return c
}

func (c *CrawlerSession) AssignPatchPersonTask(id int, name string) {
	parsedSources := lo.Map(sources, func(base string, _ int) string {
		u, err := url.Parse(base)
		if err != nil {
			log.Error(err)
		}

		queryParams := url.Values{}
		queryParams.Set("name", name)

		u.RawQuery = queryParams.Encode()

		return u.String()
	})

	c.pool.Submit(func() {
		c.crawlPatchPersonHandler(id, parsedSources)
	})
}

func (c *CrawlerSession) crawlPatchPersonHandler(id int, sources []string) {
	var patchPersonRequest requests.PatchPersonRequest

	for _, source := range sources {
		resp, err := c.client.Get(source, nil)
		if err != nil {
			log.Error(err)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
			return
		}

		err = json.Unmarshal(body, &patchPersonRequest)
		if err != nil {
			log.Error(err)
			return
		}
	}

	c.store.PersonRepository.PatchByID(context.Background(), id, &patchPersonRequest)
}

func (c *CrawlerSession) onPersonPostHandler(event events.PersonPostEvent) {
	c.AssignPatchPersonTask(event.Payload.ID, event.Payload.Name)
}
