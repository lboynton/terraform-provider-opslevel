package opslevel

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/opslevel/opslevel-go"
)

func resourceCheckRepositoryIntegrated() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a repository integrated check.",
		Create:      wrap(resourceCheckRepositoryIntegratedCreate),
		Read:        wrap(resourceCheckRepositoryIntegratedRead),
		Update:      wrap(resourceCheckRepositoryIntegratedUpdate),
		Delete:      wrap(resourceCheckDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: getCheckSchema(nil),
	}
}

func resourceCheckRepositoryIntegratedCreate(d *schema.ResourceData, client *opslevel.Client) error {
	input := opslevel.CheckRepositoryIntegratedCreateInput{}
	setCheckCreateInput(d, &input)

	resource, err := client.CreateCheckRepositoryIntegrated(input)
	if err != nil {
		return err
	}
	d.SetId(resource.Id.(string))

	return resourceCheckRepositoryIntegratedRead(d, client)
}

func resourceCheckRepositoryIntegratedRead(d *schema.ResourceData, client *opslevel.Client) error {
	id := d.Id()

	resource, err := client.GetCheck(id)
	if err != nil {
		return err
	}

	if err := setCheckData(d, resource); err != nil {
		return err
	}

	return nil
}

func resourceCheckRepositoryIntegratedUpdate(d *schema.ResourceData, client *opslevel.Client) error {
	input := opslevel.CheckRepositoryIntegratedUpdateInput{}
	setCheckUpdateInput(d, &input)

	_, err := client.UpdateCheckRepositoryIntegrated(input)
	if err != nil {
		return err
	}
	d.Set("last_updated", timeLastUpdated())
	return resourceCheckRepositoryIntegratedRead(d, client)
}
