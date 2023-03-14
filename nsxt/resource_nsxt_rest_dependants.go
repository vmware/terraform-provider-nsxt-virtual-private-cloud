/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

//nolint
package nsxt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGuestInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

//nolint
func resourcePolicyPoolUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

//nolint
func resourceServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_entries": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceServiceEntryCustomSchema(),
			},
			"service_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourcePolicyAttributesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_source": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"SYSTEM", "CUSTOM"}, false),
				Default:      "SYSTEM",
			},
			"custom_url_partial_match": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"datatype": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isalgtype": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceContextProfileAttributesMetadataSchema(),
			},
			"sub_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourcePolicySubAttributesSchema(),
			},
			"value": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceTagSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceSegmentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_state": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"UP", "DOWN"}, false),
				Default:      "UP",
			},
			"advanced_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceSegmentAdvancedConfigSchema(),
			},
			"bridge_profiles": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceBridgeProfileConfigSchema(),
			},
			"connectivity_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_config_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"evpn_tenant_config_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extra_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceSegmentExtraConfigSchema(),
			},
			"federation_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceFederationConnectivityConfigSchema(),
			},
			"l2_extension": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceL2ExtensionSchema(),
			},
			"mac_pool_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadata_proxy_paths": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"overlay_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"replication_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"MTEP", "SOURCE"}, false),
				Default:      "MTEP",
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnets": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceSegmentSubnetSchema(),
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"transport_zone_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceL2ExtensionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"l2vpn_paths": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"local_egress": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceLocalEgressSchema(),
			},
			"tunnel_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceLocalEgressRoutingEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nexthop_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix_list_paths": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourcePolicyRequestParameterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceL4PortSetServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l4_protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceSegmentDhcpV6ConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"domain_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"excluded_ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lease_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"preferred_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sntp_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceVirtualMachineSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compute_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"guest_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceGuestInfoSchema(),
			},
			"host_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_id_on_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"power_state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"runtime_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceVirtualMachineRuntimeInfoSchema(),
			},
			"scope": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceDiscoveredResourceScopeSchema(),
			},
			"source": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceResourceReferenceSchema(),
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceSegmentAdvancedConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address_pool_paths": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"connectivity": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ON", "OFF"}, false),
				Default:      "ON",
			},
			"hybrid": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"inter_router": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"local_egress": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"local_egress_routing_policies": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceLocalEgressRoutingEntrySchema(),
			},
			"multicast": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ndra_profile_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node_local_switch": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"origin_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"origin_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"VCENTER"}, false),
				Computed:     true,
			},
			"uplink_teaming_policy_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"urpf_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "STRICT"}, false),
				Default:      "STRICT",
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceDiscoveredResourceScopeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scope_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scope_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"CONTAINER_CLUSTER", "VPC"}, false),
				Computed:     true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceUnboundedKeyValuePairSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceClasslessStaticRouteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
			"next_hop": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceStaticIpAllocationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceExternalIDExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"member_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceALGTypeServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"alg": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_ports": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceConjunctionOperatorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"conjunction_operator": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceVirtualMachineRuntimeInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vif_runtime_info": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceVirtualNetworkInterfaceRuntimeInfoSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourcePortAddressBindingEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceSubnetAdvancedConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"static_ip_allocation": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceStaticIpAllocationSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceIPAddressExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_addresses": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceExcludedMembersListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_address_expression": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceIPAddressExpressionSchema(),
			},
			"path_expression": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourcePathExpressionSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceMACAddressExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_addresses": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceDhcpOption121Schema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"static_routes": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     resourceClasslessStaticRouteSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceICMPTypeServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icmp_type": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceEtherTypeServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ether_type": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceConditionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exclude": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceExcludedMembersListSchema(),
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"member_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"operator": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"EQUALS", "CONTAINS", "STARTSWITH", "ENDSWITH", "NOTEQUALS", "NOTIN", "MATCHES", "IN"}, false),
				Computed:     true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope_operator": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"EQUALS", "NOTEQUALS"}, false),
				Computed:     true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourcePathExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"paths": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceNestedServiceServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nested_service_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceSegmentExtraConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config_pair": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     resourceUnboundedKeyValuePairSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceRouterNexthopSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_distance": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scope": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceVirtualNetworkInterfaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"device_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"host_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_address_info": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceIpAddressInfoSchema(),
			},
			"lport_attachment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"owner_vm_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceDiscoveredResourceScopeSchema(),
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"vm_local_id_on_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceContextProfileAttributesMetadataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceIdentityGroupExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"identity_groups": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     resourceIdentityGroupInfoSchema(),
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceIPProtocolServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol_number": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourcePortAttachmentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allocate_addresses": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"IP_POOL", "MAC_POOL", "BOTH", "NONE", "DHCP", "DHCPV6", "SLAAC"}, false),
				Computed:     true,
			},
			"app_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bms_interface_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceAttachedInterfaceEntrySchema(),
			},
			"context_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"evpn_vlans": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hyperbus_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ENABLE", "DISABLE"}, false),
				Default:      "DISABLE",
			},
			"traffic_tag": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"PARENT", "CHILD", "INDEPENDENT", "STATIC"}, false),
				Computed:     true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceIdentityGroupInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"distinguished_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_base_distinguished_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceVirtualNetworkInterfaceRuntimeInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uptv2_active": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "True", "False"}, false),
				Computed:     true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceAttachedInterfaceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_intf_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"migrate_intf": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing_table": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceDhcpConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dhcp_relay_config_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_v4_pool_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"dhcp_v6_pool_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"dns_client_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceDnsClientConfigSchema(),
			},
			"enable_dhcp": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceResourceReferenceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceIpAddressInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

//nolint
func resourceNestedExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expressions": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     resourceExpressionSchema(),
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceGenericDhcpOptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"code": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"values": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceLocalEgressSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"optimized_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceSegmentDhcpV4ConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lease_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"options": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceDhcpV4OptionsSchema(),
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourcePolicySubAttributesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"datatype": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceIGMPTypeServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceBridgeProfileConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bridge_profile_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uplink_teaming_policy_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vlan_transport_zone_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceApplicationConnectivityStrategySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_connectivity_strategy": {
				Type:     schema.TypeString,
				Required: true,
			},
			"logging_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceSegmentDhcpConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lease_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceDnsClientConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_server_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceSegmentSubnetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dhcp_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceSegmentDhcpConfigSchema(),
			},
			"dhcp_ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"gateway_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
func resourceFederationConnectivityConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

//nolint
func resourceDhcpV4OptionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"option121": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     resourceDhcpOption121Schema(),
			},
			"others": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceGenericDhcpOptionSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//nolint
