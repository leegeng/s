package ifttt

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("ifttt", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for IFTTT.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.ifttt.com/recipes/search?q=%s", url.QueryEscape(q))
}
