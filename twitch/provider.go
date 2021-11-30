package twitch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nicklaw5/helix"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user_auth_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWITCH_ACCESS_TOKEN", nil),
			},
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWITCH_CLIENT_ID", nil),
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ShizukaSakuya",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"twitch_channel_point": resourceChannelPoint(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	auth := d.Get("user_auth_token").(string)
	clientId := d.Get("client_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client, err := helix.NewClient(&helix.Options{
		ClientID:        clientId,
		UserAccessToken: auth,
	})

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create twitch client",
			Detail:   "Unable to auth user for authenticated twitch client",
		})
		return nil, diags
	}

	return client, diags
}
