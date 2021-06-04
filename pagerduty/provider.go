package pagerduty

import (
	"context"
	"os"
	"terraform-provider-pagerduty/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("Token", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"pagerduty_user_resource": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"pagerduty_user_data": dataSourceUser(),
		},
		ConfigureContextFunc: providerConfigure,
	}

}
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	Token := d.Get("token").(string)

	os.Setenv("Token", Token)

	var diags diag.Diagnostics

	return client.NewClient(Token), diags
}
