// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceRetentlyDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceRetentlyDataSource{}

func NewSourceRetentlyDataSource() datasource.DataSource {
	return &SourceRetentlyDataSource{}
}

// SourceRetentlyDataSource is the data source implementation.
type SourceRetentlyDataSource struct {
	client *sdk.SDK
}

// SourceRetentlyDataSourceModel describes the data model.
type SourceRetentlyDataSourceModel struct {
	Configuration SourceRetently1 `tfsdk:"configuration"`
	Name          types.String    `tfsdk:"name"`
	SecretID      types.String    `tfsdk:"secret_id"`
	SourceID      types.String    `tfsdk:"source_id"`
	WorkspaceID   types.String    `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceRetentlyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_retently"
}

// Schema defines the schema for the data source.
func (r *SourceRetentlyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceRetently DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_retently_authentication_mechanism_authenticate_via_retently_o_auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Client",
											),
										},
										Description: `must be one of [Client]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of your Retently developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Client Secret of your Retently developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Retently Refresh Token which can be used to fetch new Bearer Tokens when the current one expires.`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Choose how to authenticate to Retently`,
							},
							"source_retently_authentication_mechanism_authenticate_with_api_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"api_key": schema.StringAttribute{
										Computed:    true,
										Description: `Retently API Token. See the <a href="https://app.retently.com/settings/api/tokens">docs</a> for more information on how to obtain this key.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Token",
											),
										},
										Description: `must be one of [Token]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Choose how to authenticate to Retently`,
							},
							"source_retently_update_authentication_mechanism_authenticate_via_retently_o_auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Client",
											),
										},
										Description: `must be one of [Client]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of your Retently developer application.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The Client Secret of your Retently developer application.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Retently Refresh Token which can be used to fetch new Bearer Tokens when the current one expires.`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Choose how to authenticate to Retently`,
							},
							"source_retently_update_authentication_mechanism_authenticate_with_api_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"api_key": schema.StringAttribute{
										Computed:    true,
										Description: `Retently API Token. See the <a href="https://app.retently.com/settings/api/tokens">docs</a> for more information on how to obtain this key.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Token",
											),
										},
										Description: `must be one of [Token]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Choose how to authenticate to Retently`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Choose how to authenticate to Retently`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"retently",
							),
						},
						Description: `must be one of [retently]`,
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

func (r *SourceRetentlyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceRetentlyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceRetentlyDataSourceModel
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
	request := operations.GetSourceRetentlyRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceRetently(ctx, request)
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
