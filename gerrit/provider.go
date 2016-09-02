package gerrit

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider
func Provider() terraform.ResourceProvider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GERRIT_BASE_URL", ""),
				Description: descriptions["base_url"],
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
		"base_url": "The Gerrit Base API URL",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		BaseURL: d.Get("base_url").(string),
	}

	return config.Client()
}
