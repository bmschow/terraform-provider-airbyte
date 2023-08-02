// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceAlloydbUpdateSSLModesVerifyCa1 struct {
	CaCertificate        types.String `tfsdk:"ca_certificate"`
	ClientCertificate    types.String `tfsdk:"client_certificate"`
	ClientKey            types.String `tfsdk:"client_key"`
	ClientKeyPassword    types.String `tfsdk:"client_key_password"`
	Mode                 types.String `tfsdk:"mode"`
	AdditionalProperties types.String `tfsdk:"additional_properties"`
}
