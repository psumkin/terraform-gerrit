package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/psumkin/terraform-gerrit/gerrit"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gerrit.Provider,
	})
}
