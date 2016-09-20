package gerrit

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider
func Provider() terraform.ResourceProvider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GERRIT_AUTH", "noauth"),
				Description: descriptions["auth"],
			},
			"gitCookieFile": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GERRIT_GITCOOKIE_FILE", ""),
				Description: descriptions["gitCookieFile"],
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GERRIT_PASSWORD", ""),
				Description: descriptions["password"],
			},
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GERRIT_USER", ""),
				Description: descriptions["user"],
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GERRIT_URL", ""),
				Description: descriptions["url"],
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"gerrit_project": resourceGerritProject(),
		},

		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth":          "Gerrit authorization: basic|digest|gitcookies|gitcookiefile|noauth",
		"gitCookieFile": "Filename for Gerrit gitcookiefile authorization",
		"password":      "Gerrit account password for basic|digest authorization",
		"url":           "Gerrit API URL",
		"user":          "Gerrit account username for basic|digest authorization",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Auth:          d.Get("auth").(string),
		GitCookieFile: d.Get("gitCookieFile").(string),
		Password:      d.Get("password").(string),
		User:          d.Get("user").(string),
		URL:           d.Get("url").(string),
	}

	return config.Client()
}
