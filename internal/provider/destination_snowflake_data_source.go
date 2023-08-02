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
var _ datasource.DataSource = &DestinationSnowflakeDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationSnowflakeDataSource{}

func NewDestinationSnowflakeDataSource() datasource.DataSource {
	return &DestinationSnowflakeDataSource{}
}

// DestinationSnowflakeDataSource is the data source implementation.
type DestinationSnowflakeDataSource struct {
	client *sdk.SDK
}

// DestinationSnowflakeDataSourceModel describes the data model.
type DestinationSnowflakeDataSourceModel struct {
	Configuration DestinationSnowflake1 `tfsdk:"configuration"`
	DestinationID types.String          `tfsdk:"destination_id"`
	Name          types.String          `tfsdk:"name"`
	WorkspaceID   types.String          `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationSnowflakeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_snowflake"
}

// Schema defines the schema for the data source.
func (r *DestinationSnowflakeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationSnowflake DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"destination_snowflake_authorization_method_key_pair_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Key Pair Authentication",
											),
										},
										Description: `must be one of [Key Pair Authentication]`,
									},
									"private_key": schema.StringAttribute{
										Computed:    true,
										Description: `RSA Private key to use for Snowflake connection. See the <a href="https://docs.airbyte.com/integrations/destinations/snowflake">docs</a> for more information on how to obtain this key.`,
									},
									"private_key_password": schema.StringAttribute{
										Computed:    true,
										Description: `Passphrase for private key`,
									},
								},
							},
							"destination_snowflake_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Enter you application's Access Token`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth2.0",
											),
										},
										Description: `must be one of [OAuth2.0]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Client ID`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Client secret`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Refresh Token`,
									},
								},
							},
							"destination_snowflake_authorization_method_username_and_password": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Username and Password",
											),
										},
										Description: `must be one of [Username and Password]`,
									},
									"password": schema.StringAttribute{
										Computed:    true,
										Description: `Enter the password associated with the username.`,
									},
								},
							},
							"destination_snowflake_update_authorization_method_key_pair_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Key Pair Authentication",
											),
										},
										Description: `must be one of [Key Pair Authentication]`,
									},
									"private_key": schema.StringAttribute{
										Computed:    true,
										Description: `RSA Private key to use for Snowflake connection. See the <a href="https://docs.airbyte.com/integrations/destinations/snowflake">docs</a> for more information on how to obtain this key.`,
									},
									"private_key_password": schema.StringAttribute{
										Computed:    true,
										Description: `Passphrase for private key`,
									},
								},
							},
							"destination_snowflake_update_authorization_method_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Enter you application's Access Token`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth2.0",
											),
										},
										Description: `must be one of [OAuth2.0]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Client ID`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Client secret`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your application's Refresh Token`,
									},
								},
							},
							"destination_snowflake_update_authorization_method_username_and_password": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Username and Password",
											),
										},
										Description: `must be one of [Username and Password]`,
									},
									"password": schema.StringAttribute{
										Computed:    true,
										Description: `Enter the password associated with the username.`,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"database": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the name of the <a href="https://docs.snowflake.com/en/sql-reference/ddl-database.html#database-schema-share-ddl">database</a> you want to sync data into`,
					},
					"destination_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"snowflake",
							),
						},
						Description: `must be one of [snowflake]`,
					},
					"file_buffer_count": schema.Int64Attribute{
						Computed:    true,
						Description: `Number of file buffers allocated for writing data. Increasing this number is beneficial for connections using Change Data Capture (CDC) and up to the number of streams within a connection. Increasing the number of file buffers past the maximum number of streams has deteriorating effects`,
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Description: `Enter your Snowflake account's <a href="https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#using-an-account-locator-as-an-identifier">locator</a> (in the format <account_locator>.<region>.<cloud>.snowflakecomputing.com)`,
					},
					"jdbc_url_params": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the additional properties to pass to the JDBC URL string when connecting to the database (formatted as key=value pairs separated by the symbol &). Example: key1=value1&key2=value2&key3=value3`,
					},
					"loading_method": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"destination_snowflake_data_staging_method_recommended_internal_staging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Internal Staging",
											),
										},
										Description: `must be one of [Internal Staging]`,
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
							"destination_snowflake_data_staging_method_aws_s3_staging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_key_id": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your <a href="https://docs.aws.amazon.com/powershell/latest/userguide/pstools-appendix-sign-up.html">AWS access key ID</a>. Airbyte requires Read and Write permissions on your S3 bucket `,
									},
									"encryption": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"destination_snowflake_data_staging_method_aws_s3_staging_encryption_no_encryption": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Computed: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"none",
															),
														},
														Description: `must be one of [none]`,
													},
												},
												Description: `Staging data will be stored in plaintext.`,
											},
											"destination_snowflake_data_staging_method_aws_s3_staging_encryption_aes_cbc_envelope_encryption": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Computed: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"aes_cbc_envelope",
															),
														},
														Description: `must be one of [aes_cbc_envelope]`,
													},
													"key_encrypting_key": schema.StringAttribute{
														Computed:    true,
														Description: `The key, base64-encoded. Must be either 128, 192, or 256 bits. Leave blank to have Airbyte generate an ephemeral key for each sync.`,
													},
												},
												Description: `Staging data will be encrypted using AES-CBC envelope encryption.`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
										Description: `Choose a data encryption method for the staging data`,
									},
									"file_name_pattern": schema.StringAttribute{
										Computed:    true,
										Description: `The pattern allows you to set the file-name format for the S3 staging file(s)`,
									},
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"S3 Staging",
											),
										},
										Description: `must be one of [S3 Staging]`,
									},
									"purge_staging_data": schema.BoolAttribute{
										Computed:    true,
										Description: `Toggle to delete staging files from the S3 bucket after a successful sync`,
									},
									"s3_bucket_name": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your S3 bucket name`,
									},
									"s3_bucket_region": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"",
												"us-east-1",
												"us-east-2",
												"us-west-1",
												"us-west-2",
												"af-south-1",
												"ap-east-1",
												"ap-south-1",
												"ap-northeast-1",
												"ap-northeast-2",
												"ap-northeast-3",
												"ap-southeast-1",
												"ap-southeast-2",
												"ca-central-1",
												"cn-north-1",
												"cn-northwest-1",
												"eu-central-1",
												"eu-west-1",
												"eu-west-2",
												"eu-west-3",
												"eu-south-1",
												"eu-north-1",
												"sa-east-1",
												"me-south-1",
											),
										},
										MarkdownDescription: `must be one of [, us-east-1, us-east-2, us-west-1, us-west-2, af-south-1, ap-east-1, ap-south-1, ap-northeast-1, ap-northeast-2, ap-northeast-3, ap-southeast-1, ap-southeast-2, ca-central-1, cn-north-1, cn-northwest-1, eu-central-1, eu-west-1, eu-west-2, eu-west-3, eu-south-1, eu-north-1, sa-east-1, me-south-1]` + "\n" +
											`Enter the region where your S3 bucket resides`,
									},
									"secret_access_key": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your <a href="https://docs.aws.amazon.com/powershell/latest/userguide/pstools-appendix-sign-up.html">AWS secret access key</a>`,
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
							"destination_snowflake_data_staging_method_google_cloud_storage_staging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"bucket_name": schema.StringAttribute{
										Computed:    true,
										Description: `Enter the <a href="https://cloud.google.com/storage/docs/creating-buckets">Cloud Storage bucket name</a>`,
									},
									"credentials_json": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys">Google Cloud service account key</a> in the JSON format with read/write access to your Cloud Storage staging bucket`,
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
									"project_id": schema.StringAttribute{
										Computed:    true,
										Description: `Enter the <a href="https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects">Google Cloud project ID</a>`,
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
							"destination_snowflake_data_staging_method_select_another_option": schema.SingleNestedAttribute{
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
								Description: `Select another option`,
							},
							"destination_snowflake_update_data_staging_method_recommended_internal_staging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Internal Staging",
											),
										},
										Description: `must be one of [Internal Staging]`,
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
							"destination_snowflake_update_data_staging_method_aws_s3_staging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_key_id": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your <a href="https://docs.aws.amazon.com/powershell/latest/userguide/pstools-appendix-sign-up.html">AWS access key ID</a>. Airbyte requires Read and Write permissions on your S3 bucket `,
									},
									"encryption": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"destination_snowflake_update_data_staging_method_aws_s3_staging_encryption_no_encryption": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Computed: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"none",
															),
														},
														Description: `must be one of [none]`,
													},
												},
												Description: `Staging data will be stored in plaintext.`,
											},
											"destination_snowflake_update_data_staging_method_aws_s3_staging_encryption_aes_cbc_envelope_encryption": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Computed: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"aes_cbc_envelope",
															),
														},
														Description: `must be one of [aes_cbc_envelope]`,
													},
													"key_encrypting_key": schema.StringAttribute{
														Computed:    true,
														Description: `The key, base64-encoded. Must be either 128, 192, or 256 bits. Leave blank to have Airbyte generate an ephemeral key for each sync.`,
													},
												},
												Description: `Staging data will be encrypted using AES-CBC envelope encryption.`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
										Description: `Choose a data encryption method for the staging data`,
									},
									"file_name_pattern": schema.StringAttribute{
										Computed:    true,
										Description: `The pattern allows you to set the file-name format for the S3 staging file(s)`,
									},
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"S3 Staging",
											),
										},
										Description: `must be one of [S3 Staging]`,
									},
									"purge_staging_data": schema.BoolAttribute{
										Computed:    true,
										Description: `Toggle to delete staging files from the S3 bucket after a successful sync`,
									},
									"s3_bucket_name": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your S3 bucket name`,
									},
									"s3_bucket_region": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"",
												"us-east-1",
												"us-east-2",
												"us-west-1",
												"us-west-2",
												"af-south-1",
												"ap-east-1",
												"ap-south-1",
												"ap-northeast-1",
												"ap-northeast-2",
												"ap-northeast-3",
												"ap-southeast-1",
												"ap-southeast-2",
												"ca-central-1",
												"cn-north-1",
												"cn-northwest-1",
												"eu-central-1",
												"eu-west-1",
												"eu-west-2",
												"eu-west-3",
												"eu-south-1",
												"eu-north-1",
												"sa-east-1",
												"me-south-1",
											),
										},
										MarkdownDescription: `must be one of [, us-east-1, us-east-2, us-west-1, us-west-2, af-south-1, ap-east-1, ap-south-1, ap-northeast-1, ap-northeast-2, ap-northeast-3, ap-southeast-1, ap-southeast-2, ca-central-1, cn-north-1, cn-northwest-1, eu-central-1, eu-west-1, eu-west-2, eu-west-3, eu-south-1, eu-north-1, sa-east-1, me-south-1]` + "\n" +
											`Enter the region where your S3 bucket resides`,
									},
									"secret_access_key": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your <a href="https://docs.aws.amazon.com/powershell/latest/userguide/pstools-appendix-sign-up.html">AWS secret access key</a>`,
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
							"destination_snowflake_update_data_staging_method_google_cloud_storage_staging": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"bucket_name": schema.StringAttribute{
										Computed:    true,
										Description: `Enter the <a href="https://cloud.google.com/storage/docs/creating-buckets">Cloud Storage bucket name</a>`,
									},
									"credentials_json": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys">Google Cloud service account key</a> in the JSON format with read/write access to your Cloud Storage staging bucket`,
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
									"project_id": schema.StringAttribute{
										Computed:    true,
										Description: `Enter the <a href="https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects">Google Cloud project ID</a>`,
									},
								},
								Description: `Recommended for large production workloads for better speed and scalability.`,
							},
							"destination_snowflake_update_data_staging_method_select_another_option": schema.SingleNestedAttribute{
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
								Description: `Select another option`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Select a data staging method`,
					},
					"role": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the <a href="https://docs.snowflake.com/en/user-guide/security-access-control-overview.html#roles">role</a> that you want to use to access Snowflake`,
					},
					"schema": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the name of the default <a href="https://docs.snowflake.com/en/sql-reference/ddl-database.html#database-schema-share-ddl">schema</a>`,
					},
					"username": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the name of the user you want to use to access the database`,
					},
					"warehouse": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the name of the <a href="https://docs.snowflake.com/en/user-guide/warehouses-overview.html#overview-of-warehouses">warehouse</a> that you want to sync data into`,
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

func (r *DestinationSnowflakeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *DestinationSnowflakeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationSnowflakeDataSourceModel
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
	request := operations.GetDestinationSnowflakeRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationSnowflake(ctx, request)
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
