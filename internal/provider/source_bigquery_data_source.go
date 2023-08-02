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
var _ datasource.DataSource = &SourceBigqueryDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceBigqueryDataSource{}

func NewSourceBigqueryDataSource() datasource.DataSource {
	return &SourceBigqueryDataSource{}
}

// SourceBigqueryDataSource is the data source implementation.
type SourceBigqueryDataSource struct {
	client *sdk.SDK
}

// SourceBigqueryDataSourceModel describes the data model.
type SourceBigqueryDataSourceModel struct {
	Configuration SourceBigquery `tfsdk:"configuration"`
	Name          types.String   `tfsdk:"name"`
	SecretID      types.String   `tfsdk:"secret_id"`
	SourceID      types.String   `tfsdk:"source_id"`
	WorkspaceID   types.String   `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceBigqueryDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_bigquery"
}

// Schema defines the schema for the data source.
func (r *SourceBigqueryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceBigquery DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials_json": schema.StringAttribute{
						Computed:    true,
						Description: `The contents of your Service Account Key JSON file. See the <a href="https://docs.airbyte.com/integrations/sources/bigquery#setup-the-bigquery-source-in-airbyte">docs</a> for more information on how to obtain this key.`,
					},
					"dataset_id": schema.StringAttribute{
						Computed:    true,
						Description: `The dataset ID to search for tables and views. If you are only loading data from one dataset, setting this option could result in much faster schema discovery.`,
					},
					"project_id": schema.StringAttribute{
						Computed:    true,
						Description: `The GCP project ID for the project containing the target BigQuery dataset.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"bigquery",
							),
						},
						Description: `must be one of [bigquery]`,
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

func (r *SourceBigqueryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceBigqueryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceBigqueryDataSourceModel
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
	request := operations.GetSourceBigqueryRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceBigquery(ctx, request)
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
