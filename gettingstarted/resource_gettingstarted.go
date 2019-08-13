package gettingstarted

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func resourceGettingStarted() *schema.Resource {
        return &schema.Resource{
                Create: resourceGettingStartedCreate,
                Read:   resourceGettingStartedRead,
                Update: resourceGettingStartedUpdate,
                Delete: resourceGettingStartedDelete,

                Schema: map[string]*schema.Schema{
                        "address": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func resourceGettingStartedCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return resourceGettingStartedRead(d, m)
}

func resourceGettingStartedRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceGettingStartedUpdate(d *schema.ResourceData, m interface{}) error {
        return resourceGettingStartedRead(d, m)
}

func resourceGettingStartedDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}
