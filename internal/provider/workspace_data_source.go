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
var _ datasource.DataSource = &WorkspaceDataSource{}
var _ datasource.DataSourceWithConfigure = &WorkspaceDataSource{}

func NewWorkspaceDataSource() datasource.DataSource {
	return &WorkspaceDataSource{}
}

// WorkspaceDataSource is the data source implementation.
type WorkspaceDataSource struct {
	client *sdk.SDK
}

// WorkspaceDataSourceModel describes the data model.
type WorkspaceDataSourceModel struct {
	DataResidency types.String `tfsdk:"data_residency"`
	Name          types.String `tfsdk:"name"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *WorkspaceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workspace"
}

// Schema defines the schema for the data source.
func (r *WorkspaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Workspace DataSource",

		Attributes: map[string]schema.Attribute{
			"data_residency": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"auto",
						"us",
						"eu",
					),
				},
				Description: `must be one of [auto, us, eu]`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `Name of the workspace`,
			},
			"workspace_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *WorkspaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *WorkspaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *WorkspaceDataSourceModel
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

	workspaceID := data.WorkspaceID.ValueString()
	request := operations.GetWorkspaceRequest{
		WorkspaceID: workspaceID,
	}
	res, err := r.client.Workspaces.GetWorkspace(ctx, request)
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
	if res.WorkspaceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.WorkspaceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
