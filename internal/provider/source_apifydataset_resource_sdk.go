// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceApifyDatasetResourceModel) ToCreateSDKType() *shared.SourceApifyDatasetCreateRequest {
	clean := new(bool)
	if !r.Configuration.Clean.IsUnknown() && !r.Configuration.Clean.IsNull() {
		*clean = r.Configuration.Clean.ValueBool()
	} else {
		clean = nil
	}
	datasetID := new(string)
	if !r.Configuration.DatasetID.IsUnknown() && !r.Configuration.DatasetID.IsNull() {
		*datasetID = r.Configuration.DatasetID.ValueString()
	} else {
		datasetID = nil
	}
	sourceType := shared.SourceApifyDatasetApifyDataset(r.Configuration.SourceType.ValueString())
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceApifyDataset{
		Clean:      clean,
		DatasetID:  datasetID,
		SourceType: sourceType,
		Token:      token,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceApifyDatasetCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceApifyDatasetResourceModel) ToGetSDKType() *shared.SourceApifyDatasetCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceApifyDatasetResourceModel) ToUpdateSDKType() *shared.SourceApifyDatasetPutRequest {
	clean := new(bool)
	if !r.Configuration.Clean.IsUnknown() && !r.Configuration.Clean.IsNull() {
		*clean = r.Configuration.Clean.ValueBool()
	} else {
		clean = nil
	}
	datasetID := new(string)
	if !r.Configuration.DatasetID.IsUnknown() && !r.Configuration.DatasetID.IsNull() {
		*datasetID = r.Configuration.DatasetID.ValueString()
	} else {
		datasetID = nil
	}
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceApifyDatasetUpdate{
		Clean:     clean,
		DatasetID: datasetID,
		Token:     token,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceApifyDatasetPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceApifyDatasetResourceModel) ToDeleteSDKType() *shared.SourceApifyDatasetCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceApifyDatasetResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceApifyDatasetResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
