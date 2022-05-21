package provider

import (
	"context"
	"fmt"
	"os"
	"sync"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client"
)

// Generate the Meraki API client from OpenAPI spec
//go:generate swagger generate client /f meraki-openapi.json /t .

// provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type provider struct {
	// client can contain the upstream provider SDK or HTTP client used to
	// communicate with the upstream service. Resource and DataSource
	// implementations can then make calls using this client.
	client *client.MerakiDashboardAPI

	// configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the provider was previously configured.
	configured bool

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string

	// deleteNetworkMutex ensures multiple networks are not deleted simultaneously.
	deleteNetworkMutex sync.Mutex
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	Host   types.String `tfsdk:"host"`
	ApiKey types.String `tfsdk:"api_key"`
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	var data providerData
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	host, ok := os.LookupEnv("MERAKI_HOST")
	if !ok {
		if data.Host.Unknown {
			resp.Diagnostics.AddError("Meraki host must be constant", "Meraki host must be a constant value")
			return
		}
		if !data.Host.Null {
			host = data.Host.Value
		} else {
			host = "api.meraki.com"
		}
	}

	var apiKey string
	if !data.ApiKey.Null {
		apiKey = data.ApiKey.Value
	} else if data.ApiKey.Unknown {
		resp.Diagnostics.AddError("Meraki API key must be constant", "Meraki API key must be a constant value")
		return
	} else {
		apiKey, ok = os.LookupEnv("MERAKI_API_KEY")
		if !ok {
			resp.Diagnostics.AddError("Missing Meraki API key", "Meraki API key must be specified in configuration or in the MERAKI_API_KEY environment variable.")
			return
		}
	}

	tr := httptransport.New(host, "api/v1", []string{"https"})
	tr.DefaultAuthentication = httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", apiKey)
	p.client = client.New(tr, strfmt.Default)

	p.configured = true
}

func (p *provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"meraki_network":                    networkResourceType{},
		"meraki_network_webhook_httpserver": networkWebhookHttpserverResourceType{},
	}, nil
}

func (p *provider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"meraki_organization":    organizationDataSourceType{},
		"meraki_organizations":   organizationsDataSourceType{},
		"meraki_devices":         devicesDataSourceType{},
		"meraki_device_statuses": deviceStatusesDataSourceType{},
		"meraki_networks":        networksDataSourceType{},
	}, nil
}

func (p *provider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"host": {
				MarkdownDescription: "Host name of the Meraki API",
				Optional:            true,
				Type:                types.StringType,
			},
			"api_key": {
				MarkdownDescription: "Meraki API key",
				Optional:            true,
				Type:                types.StringType,
				Sensitive:           true,
			},
		},
	}, nil
}

func New(version string) func() tfsdk.Provider {
	return func() tfsdk.Provider {
		return &provider{
			version: version,
		}
	}
}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*provider)), however using this can prevent
// potential panics.
func convertProviderType(in tfsdk.Provider) (*provider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*provider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return nil, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return nil, diags
	}

	return p, diags
}
