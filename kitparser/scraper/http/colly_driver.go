package http

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"net/url"
)

type CollyDriver struct {
	collector *colly.Collector
}

func NewCollyDriver(limitRules []*colly.LimitRule, options ...colly.CollectorOption) (*CollyDriver, error) {
	collector := colly.NewCollector(options...)
	err := collector.Limits(limitRules)
	if err != nil {
		return nil, err
	}

	return &CollyDriver{
		collector: collector,
	}, nil
}

func (s CollyDriver) SetScraper(scraper *Scraper) {
	s.collector.SetDebugger(&collyDebugger{scraper: scraper})
}
func (s CollyDriver) Visit(url string) error { return s.collector.Visit(url) }
func (s *CollyDriver) Wait()                 { s.collector.Wait() }
func (s CollyDriver) Clone() Driver          { return &CollyDriver{collector: s.collector.Clone()} }

func (s CollyDriver) OnError(f func(*Response, error)) {
	s.collector.OnError(func(response *colly.Response, err error) {
		f(convertCollyResponse(response), err)
	})
}

func (s CollyDriver) OnResponse(f func(*Response)) {
	s.collector.OnResponse(func(response *colly.Response) {
		f(convertCollyResponse(response))
	})
}

func (s CollyDriver) OnHTML(selector string, f func(element HTMLElement)) {
	s.collector.OnHTML(selector, func(element *colly.HTMLElement) {
		f(&CoollyHTMLElement{
			Element: element,
		})
	})
}

func (s *CollyDriver) SetContext(ctx context.Context) {
	s.collector.Context = ctx
}

type collyDebugger struct {
	scraper *Scraper
}

func (d collyDebugger) Init() error { return nil }

func (d collyDebugger) Event(e *debug.Event) {
	ctx := d.scraper.ctx
	listener := d.scraper.listener
	listener.Debug(ctx, baseEventType+":"+e.Type, e.Values)
}

type CoollyHTMLElement struct {
	Element *colly.HTMLElement
}

func (c CoollyHTMLElement) DOM() *goquery.Selection { return c.Element.DOM }
func (c CoollyHTMLElement) URL() *url.URL           { return c.Element.Request.URL }
func (c CoollyHTMLElement) Attr(name string) string { return c.Element.Attr(name) }
func (c CoollyHTMLElement) Text() string            { return c.Element.Text }
func (c CoollyHTMLElement) Request() *Request       { return convertCollyRequest(c.Element.Request) }

func convertCollyResponse(r *colly.Response) *Response {
	return &Response{
		Request: convertCollyRequest(r.Request),
		Body:    r.Body,
	}
}

func convertCollyRequest(r *colly.Request) *Request {
	return &Request{
		Retry: r.Retry,
		URL:   r.URL,
	}
}
