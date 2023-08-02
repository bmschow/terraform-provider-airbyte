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
var _ datasource.DataSource = &DestinationPulsarDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationPulsarDataSource{}

func NewDestinationPulsarDataSource() datasource.DataSource {
	return &DestinationPulsarDataSource{}
}

// DestinationPulsarDataSource is the data source implementation.
type DestinationPulsarDataSource struct {
	client *sdk.SDK
}

// DestinationPulsarDataSourceModel describes the data model.
type DestinationPulsarDataSourceModel struct {
	Configuration DestinationPulsar `tfsdk:"configuration"`
	DestinationID types.String      `tfsdk:"destination_id"`
	Name          types.String      `tfsdk:"name"`
	WorkspaceID   types.String      `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationPulsarDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_pulsar"
}

// Schema defines the schema for the data source.
func (r *DestinationPulsarDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationPulsar DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"batching_enabled": schema.BoolAttribute{
						Computed:    true,
						Description: `Control whether automatic batching of messages is enabled for the producer.`,
					},
					"batching_max_messages": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum number of messages permitted in a batch.`,
					},
					"batching_max_publish_delay": schema.Int64Attribute{
						Computed:    true,
						Description: ` Time period in milliseconds within which the messages sent will be batched.`,
					},
					"block_if_queue_full": schema.BoolAttribute{
						Computed:    true,
						Description: `If the send operation should block when the outgoing message queue is full.`,
					},
					"brokers": schema.StringAttribute{
						Computed:    true,
						Description: `A list of host/port pairs to use for establishing the initial connection to the Pulsar cluster.`,
					},
					"compression_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"NONE",
								"LZ4",
								"ZLIB",
								"ZSTD",
								"SNAPPY",
							),
						},
						MarkdownDescription: `must be one of [NONE, LZ4, ZLIB, ZSTD, SNAPPY]` + "\n" +
							`Compression type for the producer.`,
					},
					"destination_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"pulsar",
							),
						},
						Description: `must be one of [pulsar]`,
					},
					"max_pending_messages": schema.Int64Attribute{
						Computed:    true,
						Description: `The maximum size of a queue holding pending messages.`,
					},
					"max_pending_messages_across_partitions": schema.Int64Attribute{
						Computed:    true,
						Description: `The maximum number of pending messages across partitions.`,
					},
					"producer_name": schema.StringAttribute{
						Computed:    true,
						Description: `Name for the producer. If not filled, the system will generate a globally unique name which can be accessed with.`,
					},
					"producer_sync": schema.BoolAttribute{
						Computed:    true,
						Description: `Wait synchronously until the record has been sent to Pulsar.`,
					},
					"send_timeout_ms": schema.Int64Attribute{
						Computed:    true,
						Description: `If a message is not acknowledged by a server before the send-timeout expires, an error occurs (in ms).`,
					},
					"topic_namespace": schema.StringAttribute{
						Computed:    true,
						Description: `The administrative unit of the topic, which acts as a grouping mechanism for related topics. Most topic configuration is performed at the namespace level. Each tenant has one or multiple namespaces.`,
					},
					"topic_pattern": schema.StringAttribute{
						Computed:    true,
						Description: `Topic pattern in which the records will be sent. You can use patterns like '{namespace}' and/or '{stream}' to send the message to a specific topic based on these values. Notice that the topic name will be transformed to a standard naming convention.`,
					},
					"topic_tenant": schema.StringAttribute{
						Computed:    true,
						Description: `The topic tenant within the instance. Tenants are essential to multi-tenancy in Pulsar, and spread across clusters.`,
					},
					"topic_test": schema.StringAttribute{
						Computed:    true,
						Description: `Topic to test if Airbyte can produce messages.`,
					},
					"topic_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"persistent",
								"non-persistent",
							),
						},
						MarkdownDescription: `must be one of [persistent, non-persistent]` + "\n" +
							`It identifies type of topic. Pulsar supports two kind of topics: persistent and non-persistent. In persistent topic, all messages are durably persisted on disk (that means on multiple disks unless the broker is standalone), whereas non-persistent topic does not persist message into storage disk.`,
					},
					"use_tls": schema.BoolAttribute{
						Computed:    true,
						Description: `Whether to use TLS encryption on the connection.`,
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

func (r *DestinationPulsarDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *DestinationPulsarDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationPulsarDataSourceModel
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
	request := operations.GetDestinationPulsarRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationPulsar(ctx, request)
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
