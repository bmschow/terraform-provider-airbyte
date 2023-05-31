// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
var _ resource.Resource = &SourceTrustpilotResource{}
var _ resource.ResourceWithImportState = &SourceTrustpilotResource{}

func NewSourceTrustpilotResource() resource.Resource {
	return &SourceTrustpilotResource{}
}

// SourceTrustpilotResource defines the resource implementation.
type SourceTrustpilotResource struct {
	client *sdk.SDK
}

// SourceTrustpilotResourceModel describes the resource data model.
type SourceTrustpilotResourceModel struct {
	Configuration SourceTrustpilot `tfsdk:"configuration"`
	Name          types.String     `tfsdk:"name"`
	SecretID      types.String     `tfsdk:"secret_id"`
	SourceID      types.String     `tfsdk:"source_id"`
	SourceType    types.String     `tfsdk:"source_type"`
	WorkspaceID   types.String     `tfsdk:"workspace_id"`
}

func (r *SourceTrustpilotResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_trustpilot"
}

func (r *SourceTrustpilotResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceTrustpilot Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"business_units": schema.ListAttribute{
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplace(),
						},
						Required:    true,
						ElementType: types.StringType,
					},
					"credentials": schema.SingleNestedAttribute{
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplace(),
						},
						Required: true,
						Attributes: map[string]schema.Attribute{
							"source_trustpilot_authorization_method_o_auth_2_0": schema.SingleNestedAttribute{
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"auth_type": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oauth2.0",
											),
										},
									},
									"client_id": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"client_secret": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"refresh_token": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"token_expiry_date": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											validators.IsRFC3339(),
										},
									},
								},
							},
							"source_trustpilot_authorization_method_api_key": schema.SingleNestedAttribute{
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"apikey",
											),
										},
									},
									"client_id": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
								},
								Description: `The API key authentication method gives you access to only the streams which are part of the Public API. When you want to get streams available via the Consumer API (e.g. the private reviews) you need to use authentication method OAuth 2.0.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"source_type": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"trustpilot",
							),
						},
					},
					"start_date": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
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

func (r *SourceTrustpilotResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceTrustpilotResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceTrustpilotResourceModel
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
	res, err := r.client.Sources.CreateSourceTrustpilot(ctx, request)
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

func (r *SourceTrustpilotResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceTrustpilotResourceModel
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

func (r *SourceTrustpilotResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceTrustpilotResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceTrustpilotResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceTrustpilotResourceModel
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
	request := operations.DeleteSourceTrustpilotRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceTrustpilot(ctx, request)
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

func (r *SourceTrustpilotResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
