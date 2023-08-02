// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceCoinAPIDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceCoinAPIDataSource{}

func NewSourceCoinAPIDataSource() datasource.DataSource {
	return &SourceCoinAPIDataSource{}
}

// SourceCoinAPIDataSource is the data source implementation.
type SourceCoinAPIDataSource struct {
	client *sdk.SDK
}

// SourceCoinAPIDataSourceModel describes the data model.
type SourceCoinAPIDataSourceModel struct {
	Configuration SourceCoinAPI `tfsdk:"configuration"`
	Name          types.String  `tfsdk:"name"`
	SecretID      types.String  `tfsdk:"secret_id"`
	SourceID      types.String  `tfsdk:"source_id"`
	WorkspaceID   types.String  `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceCoinAPIDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_coin_api"
}

// Schema defines the schema for the data source.
func (r *SourceCoinAPIDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceCoinAPI DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"api_key": schema.StringAttribute{
						Computed:    true,
						Description: `API Key`,
					},
					"end_date": schema.StringAttribute{
						Computed: true,
						MarkdownDescription: `The end date in ISO 8601 format. If not supplied, data will be returned` + "\n" +
							`from the start date to the current time, or when the count of result` + "\n" +
							`elements reaches its limit.` + "\n" +
							``,
					},
					"environment": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"sandbox",
								"production",
							),
						},
						MarkdownDescription: `must be one of [sandbox, production]` + "\n" +
							`The environment to use. Either sandbox or production.` + "\n" +
							``,
					},
					"limit": schema.Int64Attribute{
						Computed: true,
						MarkdownDescription: `The maximum number of elements to return. If not supplied, the default` + "\n" +
							`is 100. For numbers larger than 100, each 100 items is counted as one` + "\n" +
							`request for pricing purposes. Maximum value is 100000.` + "\n" +
							``,
					},
					"period": schema.StringAttribute{
						Computed:    true,
						Description: `The period to use. See the documentation for a list. https://docs.coinapi.io/#list-all-periods-get`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"coin-api",
							),
						},
						Description: `must be one of [coin-api]`,
					},
					"start_date": schema.StringAttribute{
						Computed:    true,
						Description: `The start date in ISO 8601 format.`,
					},
					"symbol_id": schema.StringAttribute{
						Computed: true,
						MarkdownDescription: `The symbol ID to use. See the documentation for a list.` + "\n" +
							`https://docs.coinapi.io/#list-all-symbols-get` + "\n" +
							``,
					},
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"secret_id": schema.StringAttribute{
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *SourceCoinAPIDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceCoinAPIDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceCoinAPIDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.GetSourceCoinAPIRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceCoinAPI(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
