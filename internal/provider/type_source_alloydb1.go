// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceAlloydb1 struct {
	Database          types.String                     `tfsdk:"database"`
	Host              types.String                     `tfsdk:"host"`
	JdbcURLParams     types.String                     `tfsdk:"jdbc_url_params"`
	Password          types.String                     `tfsdk:"password"`
	Port              types.Int64                      `tfsdk:"port"`
	ReplicationMethod *SourceAlloydbReplicationMethod1 `tfsdk:"replication_method"`
	Schemas           []types.String                   `tfsdk:"schemas"`
	SourceType        types.String                     `tfsdk:"source_type"`
	SslMode           *SourceAlloydbSSLModes1          `tfsdk:"ssl_mode"`
	TunnelMethod      *SourceAlloydbSSHTunnelMethod    `tfsdk:"tunnel_method"`
	Username          types.String                     `tfsdk:"username"`
}
