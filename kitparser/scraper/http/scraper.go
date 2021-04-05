package http

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/dsxack/go/v2/kitparser/config"
	"github.com/dsxack/go/v2/kitparser/resultlistener"
	"github.com/gocolly/colly/v2"
	"net/url"
	"runtime"
)

const baseEventType = "scraper"
const errEventType = baseEventType + ":error"

type Driver interface {
	Visit(url string) error
	OnError(func(*Response, error))
	OnResponse(func(*Response))
	OnHTML(selector string, f func(elem HTMLElement))
	Clone() Driver
	SetScraper(*Scraper)
	Wait()
	SetContext(ctx context.Context)
}

type Scraper struct {
	driver   Driver
	ctx      context.Context
	listener *resultlistener.Listener
}

func (s Scraper) Visit(url string) error {
	return s.driver.Visit(url)
}
func (s Scraper) OnResponse(f func(*Response)) {
	s.driver.OnResponse(f)
}
func (s Scraper) OnHTML(selector string, f func(elem HTMLElement)) {
	s.driver.OnHTML(selector, f)
}
func (s Scraper) OnError(f func(*Response, error)) {
	s.driver.OnError(f)
}

func (s Scraper) Clone() *Scraper {
	clone := &Scraper{
		ctx:      s.ctx,
		listener: s.listener,
	}
	driver := s.driver.Clone()
	driver.SetScraper(clone)
	clone.driver = driver

	return clone
}
func (s Scraper) Wait() {
	s.driver.Wait()
}
func (s *Scraper) SetResultsListener(listener *resultlistener.Listener) {
	s.listener = listener
}

func (s Scraper) WithContext(ctx context.Context) *Scraper {
	clone := s.Clone()
	clone.ctx = ctx
	clone.driver.SetContext(ctx)
	return clone
}

func (s Scraper) Init() {
	s.driver.SetScraper(&s)
}

func NewScraper(driver Driver, opts ...Option) *Scraper {
	s := &Scraper{driver: driver}
	for _, opt := range opts {
		opt(s)
	}
	s.Init()
	return s
}

func NewEnvScraper(
	listener *resultlistener.Listener,
	cfg config.ScraperHTTP,
	opts ...Option,
) (*Scraper, error) {
	// TODO: detect env scraper here instead of CollyScraper
	parallelism := cfg.Parallelism
	if parallelism == 0 {
		parallelism = runtime.NumCPU() * 3
	}
	randomDelay := cfg.RandomDelay
	if randomDelay == 0 {
		randomDelay = 0
	}

	driver, err := NewCollyDriver(
		[]*colly.LimitRule{
			{
				DomainGlob:  "*",
				Parallelism: parallelism,
				RandomDelay: randomDelay,
			},
		},
		colly.Async(true),
		colly.AllowURLRevisit(),
	)
	if err != nil {
		return nil, err
	}
	opts = append(opts, WithDefaultResultsListener(listener))

	return NewScraper(driver, opts...), nil
}

type Request struct {
	URL   *url.URL
	Retry func() error
}

type Response struct {
	Body    []byte
	Request *Request
}

type HTMLElement interface {
	Text() string
	DOM() *goquery.Selection
	Attr(name string) string
	URL() *url.URL
	Request() *Request
}

type Option func(*Scraper)

func WithResultsListener(listener *resultlistener.Listener) Option {
	return func(scraper *Scraper) {
		scraper.SetResultsListener(listener)
	}
}

func WithDefaultResultsListener(listener *resultlistener.Listener) Option {
	return func(scraper *Scraper) {
		if scraper.listener == nil {
			scraper.SetResultsListener(listener)
		}
	}
}

func WithContext(ctx context.Context) Option {
	return func(scraper *Scraper) {
		scraper.WithContext(ctx)
	}
}
