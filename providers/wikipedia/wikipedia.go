package wikipedia

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("wikipedia", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Wikipedia.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://en.wikipedia.org/wiki/Special:Search?search=%s", url.QueryEscape(q))
}
