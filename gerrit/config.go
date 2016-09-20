package gerrit

import (
	"log"
	"net/url"

	"github.com/psumkin/golang-build/gerrit"
)

// Config represents GerritClient configuration
type Config struct {
	Auth, URL, User, Password, GitCookieFile string
	client                                   *gerrit.Client
}

// Client configures and returns a fully initialized GerritClient
func (c *Config) Client() (interface{}, error) {
	var a gerrit.Auth

	switch c.Auth {
	case "basic":
		a = gerrit.BasicAuth(c.User, c.Password)
	case "digest":
		a = gerrit.DigestAuth(c.User, c.Password)
	case "gitcookies":
		a = gerrit.GitCookiesAuth()
	case "gitcookiefile":
		a = gerrit.GitCookieFileAuth(c.GitCookieFile)
	case "noauth":
		a = gerrit.NoAuth
	default:
		log.Fatal("[ERROR] unrecognized config auth")
	}

	if _, err := url.Parse(c.URL); err != nil {
		log.Fatal("[ERROR] unrecognized config url", err)
	}

	c.client = gerrit.NewClient(c.URL, a)
	return *c, nil
}
