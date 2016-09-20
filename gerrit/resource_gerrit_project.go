package gerrit

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/psumkin/golang-build/gerrit"
)

func resourceGerritProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceGerritProjectCreate,
		Read:   resourceGerritProjectRead,
		Update: resourceGerritProjectUpdate,
		Delete: resourceGerritProjectDelete,

		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			// TODO: implement whole fildset
			// https://godoc.org/golang.org/x/build/gerrit#ProjectInfo
		},
	}
}

func resourceGerritProjectCreate(d *schema.ResourceData, m interface{}) error {
	err := resourceGerritProjectRead(d, m)
	if err == nil {
		log.Println("[DEBUG] project already exists")
		return nil
	}

	client := m.(Config).client
	name := d.Get("name").(string)

	_, err = client.CreateProject(name, gerrit.ProjectInput{
		Description: d.Get("description").(string),
	})
	if err != nil {
		return err
	}
	return resourceGerritProjectRead(d, m)
}

func resourceGerritProjectRead(d *schema.ResourceData, m interface{}) error {
	client := m.(Config).client
	name := d.Get("name").(string)

	project, err := client.GetProjectInfo(name)
	if err != nil {
		log.Println("[DEBUG] Getting project info failed:", name, err)
		d.SetId("")
		return err
	}
	log.Println("[DEBUG] Got project info:", project)

	// TODO: implement whole fildset
	// https://godoc.org/golang.org/x/build/gerrit#ProjectInfo
	d.Set("description", project.Description)
	return nil
}

func resourceGerritProjectUpdate(d *schema.ResourceData, m interface{}) error {
	// TODO: implement
	return nil
}

func resourceGerritProjectDelete(d *schema.ResourceData, m interface{}) error {
	// TODO: implement Delete as update for rename?
	err := resourceGerritProjectRead(d, m)
	if err != nil {
		return err
	}
	return nil
}
