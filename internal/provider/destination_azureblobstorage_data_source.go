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
var _ datasource.DataSource = &DestinationAzureBlobStorageDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationAzureBlobStorageDataSource{}

func NewDestinationAzureBlobStorageDataSource() datasource.DataSource {
	return &DestinationAzureBlobStorageDataSource{}
}

// DestinationAzureBlobStorageDataSource is the data source implementation.
type DestinationAzureBlobStorageDataSource struct {
	client *sdk.SDK
}

// DestinationAzureBlobStorageDataSourceModel describes the data model.
type DestinationAzureBlobStorageDataSourceModel struct {
	Configuration DestinationAzureBlobStorage `tfsdk:"configuration"`
	DestinationID types.String                `tfsdk:"destination_id"`
	Name          types.String                `tfsdk:"name"`
	WorkspaceID   types.String                `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationAzureBlobStorageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_azure_blob_storage"
}

// Schema defines the schema for the data source.
func (r *DestinationAzureBlobStorageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationAzureBlobStorage DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"azure_blob_storage_account_key": schema.StringAttribute{
						Computed:    true,
						Description: `The Azure blob storage account key.`,
					},
					"azure_blob_storage_account_name": schema.StringAttribute{
						Computed:    true,
						Description: `The account's name of the Azure Blob Storage.`,
					},
					"azure_blob_storage_container_name": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the Azure blob storage container. If not exists - will be created automatically. May be empty, then will be created automatically airbytecontainer+timestamp`,
					},
					"azure_blob_storage_endpoint_domain_name": schema.StringAttribute{
						Computed:    true,
						Description: `This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.`,
					},
					"azure_blob_storage_output_buffer_size": schema.Int64Attribute{
						Computed:    true,
						Description: `The amount of megabytes to buffer for the output stream to Azure. This will impact memory footprint on workers, but may need adjustment for performance and appropriate block size in Azure.`,
					},
					"azure_blob_storage_spill_size": schema.Int64Attribute{
						Computed:    true,
						Description: `The amount of megabytes after which the connector should spill the records in a new blob object. Make sure to configure size greater than individual records. Enter 0 if not applicable`,
					},
					"destination_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"azure-blob-storage",
							),
						},
						Description: `must be one of [azure-blob-storage]`,
					},
					"format": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"destination_azure_blob_storage_output_format_csv_comma_separated_values": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"flattening": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"No flattening",
												"Root level flattening",
											),
										},
										MarkdownDescription: `must be one of [No flattening, Root level flattening]` + "\n" +
											`Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details.`,
									},
									"format_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"CSV",
											),
										},
										Description: `must be one of [CSV]`,
									},
								},
								Description: `Output data format`,
							},
							"destination_azure_blob_storage_output_format_json_lines_newline_delimited_json": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"format_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"JSONL",
											),
										},
										Description: `must be one of [JSONL]`,
									},
								},
								Description: `Output data format`,
							},
							"destination_azure_blob_storage_update_output_format_csv_comma_separated_values": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"flattening": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"No flattening",
												"Root level flattening",
											),
										},
										MarkdownDescription: `must be one of [No flattening, Root level flattening]` + "\n" +
											`Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details.`,
									},
									"format_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"CSV",
											),
										},
										Description: `must be one of [CSV]`,
									},
								},
								Description: `Output data format`,
							},
							"destination_azure_blob_storage_update_output_format_json_lines_newline_delimited_json": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"format_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"JSONL",
											),
										},
										Description: `must be one of [JSONL]`,
									},
								},
								Description: `Output data format`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Output data format`,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *DestinationAzureBlobStorageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *DestinationAzureBlobStorageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationAzureBlobStorageDataSourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationAzureBlobStorageRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationAzureBlobStorage(ctx, request)
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
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
