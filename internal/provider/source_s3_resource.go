// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceS3Resource{}
var _ resource.ResourceWithImportState = &SourceS3Resource{}

func NewSourceS3Resource() resource.Resource {
	return &SourceS3Resource{}
}

// SourceS3Resource defines the resource implementation.
type SourceS3Resource struct {
	client *sdk.SDK
}

// SourceS3ResourceModel describes the resource data model.
type SourceS3ResourceModel struct {
	Configuration SourceS3     `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	SourceType    types.String `tfsdk:"source_type"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

func (r *SourceS3Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_s3"
}

func (r *SourceS3Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceS3 Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"dataset": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
					},
					"format": schema.SingleNestedAttribute{
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplace(),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_s3_file_format_csv": schema.SingleNestedAttribute{
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_reader_options": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"advanced_options": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"block_size": schema.Int64Attribute{
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"delimiter": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"double_quote": schema.BoolAttribute{
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"encoding": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"escape_char": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"filetype": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"csv",
											),
										},
									},
									"infer_datatypes": schema.BoolAttribute{
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"newlines_in_values": schema.BoolAttribute{
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"quote_char": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
								},
								Description: `This connector utilises <a href="https: // arrow.apache.org/docs/python/generated/pyarrow.csv.open_csv.html" target="_blank">PyArrow (Apache Arrow)</a> for CSV parsing.`,
							},
							"source_s3_file_format_parquet": schema.SingleNestedAttribute{
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"batch_size": schema.Int64Attribute{
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"buffer_size": schema.Int64Attribute{
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"columns": schema.ListAttribute{
										PlanModifiers: []planmodifier.List{
											listplanmodifier.RequiresReplace(),
										},
										Optional:    true,
										ElementType: types.StringType,
									},
									"filetype": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"parquet",
											),
										},
									},
								},
								Description: `This connector utilises <a href="https://arrow.apache.org/docs/python/generated/pyarrow.parquet.ParquetFile.html" target="_blank">PyArrow (Apache Arrow)</a> for Parquet parsing.`,
							},
							"source_s3_file_format_avro": schema.SingleNestedAttribute{
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"filetype": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"avro",
											),
										},
									},
								},
								Description: `This connector utilises <a href="https://fastavro.readthedocs.io/en/latest/" target="_blank">fastavro</a> for Avro parsing.`,
							},
							"source_s3_file_format_jsonl": schema.SingleNestedAttribute{
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"block_size": schema.Int64Attribute{
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"filetype": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"jsonl",
											),
										},
									},
									"newlines_in_values": schema.BoolAttribute{
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"unexpected_field_behavior": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"ignore",
												"infer",
												"error",
											),
										},
										Description: `How JSON fields outside of explicit_schema (if given) are treated. Check <a href="https://arrow.apache.org/docs/python/generated/pyarrow.json.ParseOptions.html" target="_blank">PyArrow documentation</a> for details`,
									},
								},
								Description: `This connector uses <a href="https://arrow.apache.org/docs/python/json.html" target="_blank">PyArrow</a> for JSON Lines (jsonl) file parsing.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"path_pattern": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
					},
					"provider": schema.SingleNestedAttribute{
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplace(),
						},
						Required: true,
						Attributes: map[string]schema.Attribute{
							"aws_access_key_id": schema.StringAttribute{
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
								Optional: true,
							},
							"aws_secret_access_key": schema.StringAttribute{
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
								Optional: true,
							},
							"bucket": schema.StringAttribute{
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
								Required: true,
							},
							"endpoint": schema.StringAttribute{
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
								Optional: true,
							},
							"path_prefix": schema.StringAttribute{
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
								Optional: true,
							},
							"start_date": schema.StringAttribute{
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Validators: []validator.String{
									validators.IsRFC3339(),
								},
							},
						},
						Description: `Use this to load files from S3 or S3-compatible services`,
					},
					"schema": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Optional: true,
					},
					"source_type": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"s3",
							),
						},
					},
				},
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"secret_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional: true,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
			},
			"source_type": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
		},
	}
}

func (r *SourceS3Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceS3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceS3ResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
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

	request := *data.ToCreateSDKType()
	res, err := r.client.Sources.CreateSourceS3(ctx, request)
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
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceS3Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceS3ResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceS3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceS3ResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceS3Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceS3ResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.DeleteSourceS3Request{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceS3(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SourceS3Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
