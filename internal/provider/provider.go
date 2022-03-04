package provider

import (
	"context"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client"
)

//go:generate swagger generate client /f meraki-openapi.json /t .

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"host": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("MERAKI_HOST", "api.meraki.com"),
				},
				"api_key": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("MERAKI_API_KEY", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"meraki_device_statuses": dataSourceDeviceStatuses(),
			},
			ResourcesMap: map[string]*schema.Resource{},
		}

		p.ConfigureContextFunc = providerConfigure

		return p
	}
}

type providerContext struct {
	*client.MerakiDashboardAPI
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("host").(string)
	apiKey := d.Get("api_key").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if apiKey == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Meraki client",
			Detail:   "Unable to create anonymous Meraki client",
		})
		return nil, diags
	}

	tr := httptransport.New(host, "api/v1", []string{"https"})
	tr.DefaultAuthentication = httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", apiKey)
	c := client.New(tr, strfmt.Default)

	return &providerContext{
		MerakiDashboardAPI: c,
	}, diags
}
