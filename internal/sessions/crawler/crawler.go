package crawler

import (
	"encoding/json"
	"io"
	"net/url"
	"time"

	"github.com/gojek/heimdall/hystrix"
	"github.com/kitanoyoru/effective-mobile-task/internal/config"
	"github.com/kitanoyoru/effective-mobile-task/internal/dtos"
	"github.com/kitanoyoru/effective-mobile-task/internal/sessions/store"
	"github.com/kitanoyoru/effective-mobile-task/pkg/pool"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

const (
	defaultSourceRetryCount = 5
	defaultSourceTimeout    = 1000 * time.Millisecond

	agePersonSource         = "https://api.agify.io/?name=Dmitriy"
	genderPersonSource      = "https://api.genderize.io/?name=Dmitriy"
	nationalityPersonSource = "https://api.nationalize.io/?name=Dmitriy"
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
}

func NewCrawlerSession(cfg *config.CrawlerConfig, store *store.StoreSession) *CrawlerSession {
	client := hystrix.NewClient(
		hystrix.WithHTTPTimeout(defaultSourceTimeout),
		hystrix.WithHystrixTimeout(3*defaultSourceTimeout),
	)
	pool := pool.NewPool(cfg.Size)

	return &CrawlerSession{
		client,
		pool,
		store,
	}
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
	var person *dtos.PersonPatchDTO

	for _, source := range sources {
		res, err := c.client.Get(source, nil)
		if err != nil {
			log.Error(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Error(err)
		}

		err = json.Unmarshal(body, &person)
		if err != nil {
			log.Error(err)
		}
	}

	c.store.PersonRepository.PatchByID(id, person)
}
