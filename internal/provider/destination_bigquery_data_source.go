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
var _ datasource.DataSource = &DestinationBigqueryDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationBigqueryDataSource{}

func NewDestinationBigqueryDataSource() datasource.DataSource {
	return &DestinationBigqueryDataSource{}
}

// DestinationBigqueryDataSource is the data source implementation.
type DestinationBigqueryDataSource struct {
	client *sdk.SDK
}

// DestinationBigqueryDataSourceModel describes the data model.
type DestinationBigqueryDataSourceModel struct {
	Configuration DestinationBigquery `tfsdk:"configuration"`
	DestinationID types.String        `tfsdk:"destination_id"`
	Name          types.String        `tfsdk:"name"`
	WorkspaceID   types.String        `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationBigqueryDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_bigquery"
}

// Schema defines the schema for the data source.
func (r *DestinationBigqueryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationBigquery DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"big_query_client_buffer_size_mb": schema.Int64Attribute{
						Computed:    true,
						Description: `Google BigQuery client's chunk (buffer) size (MIN=1, MAX = 15) for each table. The size that will be written by a single RPC. Written data will be buffered and only flushed upon reaching this size or closing the channel. The default 15MB value is used if not set explicitly. Read more <a href="https://googleapis.dev/python/bigquery/latest/generated/google.cloud.bigquery.client.Client.html">here</a>.`,
					},
					"credentials_json": schema.StringAttribute{
						Computed:    true,
						Description: `The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.com/integrations/destinations/bigquery#service-account-key">docs</a> if you need help generating this key. Default credentials will be used if this field is left empty.`,
					},
					"dataset_id": schema.StringAttribute{
						Computed:    true,
						Description: `The default BigQuery Dataset ID that tables are replicated to if the source does not specify a namespace. Read more <a href="https://cloud.google.com/bigquery/docs/datasets#create-dataset">here</a>.`,
					},
					"dataset_location": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"US",
								"EU",
								"asia-east1",
								"asia-east2",
								"asia-northeast1",
								"asia-northeast2",
								"asia-northeast3",
								"asia-south1",
								"asia-south2",
								"asia-southeast1",
								"asia-southeast2",
								"australia-southeast1",
								"australia-southeast2",
								"europe-central1",
								"europe-central2",
								"europe-north1",
								"europe-southwest1",
								"europe-west1",
								"europe-west2",
								"europe-west3",
								"europe-west4",
								"europe-west6",
								"europe-west7",
								"europe-west8",
								"europe-west9",
								"me-west1",
								"northamerica-northeast1",
								"northamerica-northeast2",
								"southamerica-east1",
								"southamerica-west1",
								"us-central1",
								"us-east1",
								"us-east2",
								"us-east3",
								"us-east4",
								"us-east5",
								"us-west1",
								"us-west2",
								"us-west3",
								"us-west4",
							),
						},
						MarkdownDescription: `must be one of [US, EU, asia-east1, asia-east2, asia-northeast1, asia-northeast2, asia-northeast3, asia-south1, asia-south2, asia-southeast1, asia-southeast2, australia-southeast1, australia-southeast2, europe-central1, europe-central2, europe-north1, europe-southwest1, europe-west1, europe-west2, europe-west3, europe-west4, europe-west6, europe-west7, europe-west8, europe-west9, me-west1, northamerica-northeast1, northamerica-northeast2, southamerica-east1, southamerica-west1, us-central1, us-east1, us-east2, us-east3, us-east4, us-east5, us-west1, us-west2, us-west3, us-west4]` + "\n" +
							`The location of the dataset. Warning: Changes made after creation will not be applied. Read more <a href="https://cloud.google.com/bigquery/docs/locations">here</a>.`,
					},
					"destination_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"bigquery",
							),
						},
						Description: `must be one of [bigquery]`,
					},
					"loading_method": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"destination_bigquery_loading_method_gcs_staging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"credential": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"destination_bigquery_loading_method_gcs_staging_credential_hmac_key": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"credential_type": schema.StringAttribute{
														Computed: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"HMAC_KEY",
															),
														},
														Description: `must be one of [HMAC_KEY]`,
													},
													"hmac_key_access_id": schema.StringAttribute{
														Computed:    true,
														Description: `HMAC key access ID. When linked to a service account, this ID is 61 characters long; when linked to a user account, it is 24 characters long.`,
													},
													"hmac_key_secret": schema.StringAttribute{
														Computed:    true,
														Description: `The corresponding secret for the access ID. It is a 40-character base-64 encoded string.`,
													},
												},
												Description: `An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
										Description: `An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.`,
									},
									"file_buffer_count": schema.Int64Attribute{
										Computed:    true,
										Description: `Number of file buffers allocated for writing data. Increasing this number is beneficial for connections using Change Data Capture (CDC) and up to the number of streams within a connection. Increasing the number of file buffers past the maximum number of streams has deteriorating effects`,
									},
									"gcs_bucket_name": schema.StringAttribute{
										Computed:    true,
										Description: `The name of the GCS bucket. Read more <a href="https://cloud.google.com/storage/docs/naming-buckets">here</a>.`,
									},
									"gcs_bucket_path": schema.StringAttribute{
										Computed:    true,
										Description: `Directory under the GCS bucket where data will be written.`,
									},
									"keep_files_in_gcs_bucket": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Delete all tmp files from GCS",
												"Keep all tmp files in GCS",
											),
										},
										MarkdownDescription: `must be one of [Delete all tmp files from GCS, Keep all tmp files in GCS]` + "\n" +
											`This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.`,
									},
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"GCS Staging",
											),
										},
										Description: `must be one of [GCS Staging]`,
									},
								},
								Description: `Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.`,
							},
							"destination_bigquery_loading_method_standard_inserts": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Standard",
											),
										},
										Description: `must be one of [Standard]`,
									},
								},
								Description: `Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.`,
							},
							"destination_bigquery_update_loading_method_gcs_staging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"credential": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"destination_bigquery_update_loading_method_gcs_staging_credential_hmac_key": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"credential_type": schema.StringAttribute{
														Computed: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"HMAC_KEY",
															),
														},
														Description: `must be one of [HMAC_KEY]`,
													},
													"hmac_key_access_id": schema.StringAttribute{
														Computed:    true,
														Description: `HMAC key access ID. When linked to a service account, this ID is 61 characters long; when linked to a user account, it is 24 characters long.`,
													},
													"hmac_key_secret": schema.StringAttribute{
														Computed:    true,
														Description: `The corresponding secret for the access ID. It is a 40-character base-64 encoded string.`,
													},
												},
												Description: `An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
										Description: `An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.`,
									},
									"file_buffer_count": schema.Int64Attribute{
										Computed:    true,
										Description: `Number of file buffers allocated for writing data. Increasing this number is beneficial for connections using Change Data Capture (CDC) and up to the number of streams within a connection. Increasing the number of file buffers past the maximum number of streams has deteriorating effects`,
									},
									"gcs_bucket_name": schema.StringAttribute{
										Computed:    true,
										Description: `The name of the GCS bucket. Read more <a href="https://cloud.google.com/storage/docs/naming-buckets">here</a>.`,
									},
									"gcs_bucket_path": schema.StringAttribute{
										Computed:    true,
										Description: `Directory under the GCS bucket where data will be written.`,
									},
									"keep_files_in_gcs_bucket": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Delete all tmp files from GCS",
												"Keep all tmp files in GCS",
											),
										},
										MarkdownDescription: `must be one of [Delete all tmp files from GCS, Keep all tmp files in GCS]` + "\n" +
											`This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.`,
									},
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"GCS Staging",
											),
										},
										Description: `must be one of [GCS Staging]`,
									},
								},
								Description: `Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.`,
							},
							"destination_bigquery_update_loading_method_standard_inserts": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Standard",
											),
										},
										Description: `must be one of [Standard]`,
									},
								},
								Description: `Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.`,
					},
					"project_id": schema.StringAttribute{
						Computed:    true,
						Description: `The GCP project ID for the project containing the target BigQuery dataset. Read more <a href="https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects">here</a>.`,
					},
					"transformation_priority": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"interactive",
								"batch",
							),
						},
						MarkdownDescription: `must be one of [interactive, batch]` + "\n" +
							`Interactive run type means that the query is executed as soon as possible, and these queries count towards concurrent rate limit and daily limit. Read more about interactive run type <a href="https://cloud.google.com/bigquery/docs/running-queries#queries">here</a>. Batch queries are queued and started as soon as idle resources are available in the BigQuery shared resource pool, which usually occurs within a few minutes. Batch queries don’t count towards your concurrent rate limit. Read more about batch queries <a href="https://cloud.google.com/bigquery/docs/running-queries#batch">here</a>. The default "interactive" value is used if not set explicitly.`,
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

func (r *DestinationBigqueryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *DestinationBigqueryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationBigqueryDataSourceModel
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
	request := operations.GetDestinationBigqueryRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationBigquery(ctx, request)
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
