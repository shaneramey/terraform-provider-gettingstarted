package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/shaneramey/terraform-provider-gettingstarted/gettingstarted"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return gettingstarted.Provider()
		},
	})
}
