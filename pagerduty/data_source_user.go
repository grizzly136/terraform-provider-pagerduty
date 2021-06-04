package pagerduty

import (
	"context"
	"fmt"
	"log"
	"terraform-provider-pagerduty/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"contact_methods": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"summary": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	Apiclient := m.(*client.Client)
	fmt.Println("Data source read")
	id := d.Get("id").(string)
	User_response, err := Apiclient.GetUser(id)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return diag.Errorf("error in fetching the user")
	}
	d.Set("id", User_response.User.Id)
	d.Set("name", User_response.User.Name)
	d.Set("email", User_response.User.Email)
	d.Set("type", User_response.User.Type)
	d.Set("role", User_response.User.Role)
	d.SetId(User_response.User.Id)
	contact_methods_list := make([]interface{}, len(User_response.User.Contact_methods))

	for i, com := range User_response.User.Contact_methods {
		contact := make(map[string]interface{})
		contact["type"] = com.Type
		contact["summary"] = com.Summary

		contact_methods_list[i] = contact
	}
	d.Set("contact_methods", contact_methods_list)
	return diags
}
