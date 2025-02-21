package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strconv"
	"terraform-provider-technitium-dns/internal/technitium"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &dhcpScopeResource{}
	_ resource.ResourceWithConfigure = &dhcpScopeResource{}
)

func NewDhcpScopeResource() resource.Resource {
	return &dhcpScopeResource{}
}

type dhcpScopeResource struct {
	client *technitium.Client
}

// Configure adds the provider configured client to the resource.
func (r *dhcpScopeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*technitium.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *technitium.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Metadata returns the resource type name.
func (r *dhcpScopeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dhcp_scope"
}

func (r *dhcpScopeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"starting_address": schema.StringAttribute{
				Required: true,
			},
			"ending_address": schema.StringAttribute{
				Required: true,
			},
			"subnet_mask": schema.StringAttribute{
				Required: true,
			},
			"lease_time_days": schema.Int32Attribute{
				Optional: true,
			},
			"lease_time_hours": schema.Int32Attribute{
				Optional: true,
			},
			"lease_time_minutes": schema.Int32Attribute{
				Optional: true,
			},
			"offer_delay_time": schema.Int32Attribute{
				Optional: true,
			},
			"ping_check_enabled": schema.BoolAttribute{
				Optional: true,
			},
			"ping_check_timeout": schema.Int32Attribute{
				Optional: true,
			},
			"ping_check_retries": schema.Int32Attribute{
				Optional: true,
			},
			"domain_name": schema.StringAttribute{
				Optional: true,
			},
			"domain_search_list": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"dns_updates": schema.BoolAttribute{
				Optional: true,
			},
			"dns_ttl": schema.Int32Attribute{
				Optional: true,
			},
			"server_address": schema.StringAttribute{
				Optional: true,
			},
			"server_host_name": schema.StringAttribute{
				Optional: true,
			},
			"boot_file_name": schema.StringAttribute{
				Optional: true,
			},
			"router_address": schema.StringAttribute{
				Optional: true,
			},
			"use_this_dns_server": schema.BoolAttribute{
				Optional: true,
			},
			"dns_servers": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"wins_servers": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"ntp_servers": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"static_routes": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"destination": schema.StringAttribute{
							Required: true,
						},
						"subnet_mask": schema.StringAttribute{
							Required: true,
						},
						"router": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"vendor_info": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"identifier": schema.StringAttribute{
							Required: true,
						},
						"information": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"capwap_ac_ip_addresses": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"tftp_server_addresses": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"generic_options": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"code": schema.Int32Attribute{
							Required: true,
						},
						"value": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"exclusions": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"starting_address": schema.StringAttribute{
							Required: true,
						},
						"ending_address": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"reserved_leases": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							Optional: true,
						},
						"hardware_address": schema.StringAttribute{
							Required: true,
						},
						"address": schema.StringAttribute{
							Required: true,
						},
						"comments": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"allow_only_reserved_leases": schema.BoolAttribute{
				Optional: true,
			},
			"block_locally_administered_mac_addresses": schema.BoolAttribute{
				Optional: true,
			},
			"ignore_client_identifier_option": schema.BoolAttribute{
				Optional: true,
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *dhcpScopeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan technitium.Scope
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new order
	order, err := r.client.CreateScope(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating order",
			"Could not create order, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(strconv.Itoa(order.ID))
	for orderItemIndex, orderItem := range order.Items {
		plan.Items[orderItemIndex] = orderItemModel{
			Coffee: orderItemCoffeeModel{
				ID:          types.Int64Value(int64(orderItem.Coffee.ID)),
				Name:        types.StringValue(orderItem.Coffee.Name),
				Teaser:      types.StringValue(orderItem.Coffee.Teaser),
				Description: types.StringValue(orderItem.Coffee.Description),
				Price:       types.Float64Value(orderItem.Coffee.Price),
				Image:       types.StringValue(orderItem.Coffee.Image),
			},
			Quantity: types.Int64Value(int64(orderItem.Quantity)),
		}
	}
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

// Read refreshes the Terraform state with the latest data.
func (r *dhcpScopeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *dhcpScopeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *dhcpScopeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
