/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

// Auto generated code. DO NOT EDIT.

// nolint
package nsxt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFederationConnectivityConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func resourceBridgeProfileConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"uplink_teaming_policy_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bridge_profile_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_transport_zone_path": {
				Type:     schema.TypeString,
				Required: true,
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
				MinItems: 1,
				MaxItems: 10,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

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

func resourceServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"service_entries": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceServiceEntryCustomSchema(),
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSubnetAdvancedConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"static_ip_allocation": {
				Type:     schema.TypeSet,
				MaxItems: 1,
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

func resourceSegmentDhcpConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"server_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dns_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 2,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lease_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceLocalEgressSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"optimized_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 1,
				MaxItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

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

func resourceIPProtocolServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"protocol_number": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSegmentDhcpV4ConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"options": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceDhcpV4OptionsSchema(),
			},
			"server_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dns_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 2,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lease_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

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

func resourceDhcpV4OptionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"option121": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceDhcpOption121Schema(),
			},
			"others": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 255,
				Elem:     resourceGenericDhcpOptionSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIdentityGroupInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distinguished_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_base_distinguished_name": {
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

func resourceServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourcePolicyAttributesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"isalgtype": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_url_partial_match": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"attribute_source": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"SYSTEM", "CUSTOM"}, false),
				Default:      "SYSTEM",
			},
			"datatype": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sub_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourcePolicySubAttributesSchema(),
			},
			"metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceContextProfileAttributesMetadataSchema(),
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

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

func resourceSegmentExtraConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config_pair": {
				Type:     schema.TypeSet,
				MaxItems: 1,
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
				MaxItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourcePolicyPoolUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func resourceApplicationConnectivityStrategySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"logging_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"application_connectivity_strategy": {
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

func resourcePortAttachmentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"context_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bms_interface_config": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceAttachedInterfaceEntrySchema(),
			},
			"allocate_addresses": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"IP_POOL", "MAC_POOL", "BOTH", "NONE", "DHCP", "DHCPV6", "SLAAC"}, false),
				Computed:     true,
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
			"hyperbus_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ENABLE", "DISABLE"}, false),
				Default:      "DISABLE",
			},
			"evpn_vlans": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 1000,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

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

func resourceConjunctionOperatorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"conjunction_operator": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGuestInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func resourceIPAddressExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"ip_addresses": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 6000,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceL2ExtensionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"local_egress": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceLocalEgressSchema(),
			},
			"l2vpn_path": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: "This is a deprecated property. Please refer the documentation for details, and refrain from use as this will be removed in future versions.",
			},
			"l2vpn_paths": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceICMPTypeServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
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
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSegmentPortSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attachment": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourcePortAttachmentSchema(),
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"init_state": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"UNBLOCKED_VLAN", "RESTORE_VIF"}, false),
				Computed:     true,
			},
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
			"extra_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceSegmentExtraConfigSchema(),
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"address_bindings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 512,
				Elem:     resourcePortAddressBindingEntrySchema(),
			},
			"ignored_address_bindings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 16,
				Elem:     resourcePortAddressBindingEntrySchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceConditionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exclude": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceExcludedMembersListSchema(),
			},
			"scope_operator": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"EQUALS", "NOTEQUALS"}, false),
				Computed:     true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operator": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"EQUALS", "CONTAINS", "STARTSWITH", "ENDSWITH", "NOTEQUALS", "NOTIN", "MATCHES", "IN"}, false),
				Computed:     true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"member_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceL4PortSetServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"l4_protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"destination_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 15,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"source_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 15,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVirtualMachineSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest_info": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceGuestInfoSchema(),
			},
			"source": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceResourceReferenceSchema(),
			},
			"runtime_info": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceVirtualMachineRuntimeInfoSchema(),
			},
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_id_on_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"power_state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceDiscoveredResourceScopeSchema(),
			},
			"compute_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVirtualNetworkInterfaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"device_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lport_attachment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"host_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"owner_vm_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vm_local_id_on_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_address_info": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceIpAddressInfoSchema(),
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
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceExcludedMembersListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_address_expression": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceIPAddressExpressionSchema(),
			},
			"path_expression": {
				Type:     schema.TypeSet,
				MaxItems: 1,
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

func resourceMACAddressExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"mac_addresses": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 4000,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceExternalIDExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"member_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"external_ids": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceStaticPoolConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ipv4_pool_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDhcpOption121Schema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"static_routes": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 27,
				Elem:     resourceClasslessStaticRouteSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIGMPTypeServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSegmentDhcpV6ConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"server_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dns_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 2,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sntp_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 2,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"preferred_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lease_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"excluded_ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 128,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVpcSubnetDhcpConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_dhcp": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"dhcp_relay_config_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_client_config": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceDnsClientConfigSchema(),
			},
			"static_pool_config": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceStaticPoolConfigSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceNestedServiceServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
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
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIpAddressInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func resourceEtherTypeServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ether_type": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceStaticIpAllocationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
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

func resourceAttachedInterfaceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"migrate_intf": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_intf_name": {
				Type:     schema.TypeString,
				Required: true,
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

func resourceRouterNexthopSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scope": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"admin_distance": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSegmentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_pool_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_config_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"evpn_tenant_config_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connectivity_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transport_zone_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"federation_config": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceFederationConnectivityConfigSchema(),
			},
			"l2_extension": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceL2ExtensionSchema(),
			},
			"advanced_config": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceSegmentAdvancedConfigSchema(),
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ls_id": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: "This is a deprecated property. Please refer the documentation for details, and refrain from use as this will be removed in future versions.",
			},
			"replication_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"MTEP", "SOURCE"}, false),
				Default:      "MTEP",
			},
			"admin_state": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"UP", "DOWN"}, false),
				Default:      "UP",
			},
			"bridge_profiles": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceBridgeProfileConfigSchema(),
			},
			"metadata_proxy_paths": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"subnets": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceSegmentSubnetSchema(),
			},
			"extra_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceSegmentExtraConfigSchema(),
			},
			"vlan_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"overlay_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"address_bindings": {
				Type:       schema.TypeList,
				Optional:   true,
				Computed:   true,
				MaxItems:   512,
				Deprecated: "This is a deprecated property. Please refer the documentation for details, and refrain from use as this will be removed in future versions.",
				Elem:       resourcePortAddressBindingEntrySchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSegmentSubnetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dhcp_config": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem:     resourceSegmentDhcpConfigSchema(),
			},
			"gateway_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_ranges": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 1,
				MaxItems: 99,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceALGTypeServiceEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"alg": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_ports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 15,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"destination_ports": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTagSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"scope": {
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

func resourceSegmentAdvancedConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_local_switch": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"multicast": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"origin_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uplink_teaming_policy_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ndra_profile_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_egress": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"origin_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"VCENTER"}, false),
				Computed:     true,
			},
			"connectivity": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ON", "OFF"}, false),
				Default:      "ON",
			},
			"urpf_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "STRICT"}, false),
				Default:      "STRICT",
			},
			"local_egress_routing_policies": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 1,
				Elem:     resourceLocalEgressRoutingEntrySchema(),
			},
			"address_pool_paths": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

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
				MinItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourcePathExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"paths": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIdentityGroupExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"identity_groups": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 500,
				Elem:     resourceIdentityGroupInfoSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceNestedExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"_revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 30,
				Elem:     resourceTagSchema(),
			},
			"expressions": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem:     resourceExpressionSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
