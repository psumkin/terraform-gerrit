package gerrit

import "golang.org/x/build/gerrit"

// Config represents GerritClient configuration
type Config struct {
	BaseURL string
	client  *gerrit.Client
}

// Client configures and returns a fully initialized GerritClient
func (c *Config) Client() (interface{}, error) {
	// c.client = gerrit.NewClient(c.BaseURL, gerrit.NoAuth)
	// c.client = gerrit.NewClient(c.BaseURL, gerrit.BasicAuth("admin", "secret"))
	c.client = gerrit.NewClient(c.BaseURL, gerrit.BasicAuth("terraform-bot", "terraform-bot"))
	return *c, nil
}
