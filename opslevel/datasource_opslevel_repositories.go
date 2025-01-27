package opslevel

import (
	"github.com/opslevel/opslevel-go"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceRepositories() *schema.Resource {
	return &schema.Resource{
		Read: wrap(datasourceRepositoriesRead),
		Schema: map[string]*schema.Schema{
			"filter": getDatasourceFilter(false, []string{"tier"}),
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"urls": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func datasourceRepositoriesRead(d *schema.ResourceData, client *opslevel.Client) error {
	field := d.Get("filter.0.field").(string)
	value := d.Get("filter.0.value").(string)

	var teams []opslevel.Repository
	var err error
	switch field {
	case "tier":
		teams, err = client.ListRepositoriesWithTier(value)
	default:
		teams, err = client.ListRepositories()
	}
	if err != nil {
		return err
	}

	count := len(teams)
	ids := make([]string, count)
	names := make([]string, count)
	urls := make([]string, count)
	for i, item := range teams {
		ids[i] = item.Id.(string)
		names[i] = item.Name
		urls[i] = item.Url
	}

	d.SetId(timeID())
	d.Set("ids", ids)
	d.Set("names", names)
	d.Set("urls", urls)

	return nil
}
