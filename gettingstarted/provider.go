package gettingstarted

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"gettingstarted": resourceGettingStarted(),
		},
	}
}
