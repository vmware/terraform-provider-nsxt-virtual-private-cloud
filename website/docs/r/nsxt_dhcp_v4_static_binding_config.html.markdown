---
layout: "nsxt"
page_title: "Nsxt: nsxt_vpc_dhcp_v4_static_binding_config"
sidebar_current: "docs-nsxt-vpc-dhcpv4staticbindingconfig"
description: |-
  Creates and manages DhcpV4StaticBindingConfig.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# nsxt_vpc_dhcp_v4_static_binding_config

The DhcpV4StaticBindingConfig resource allows the creation and management of Nsxt DhcpV4StaticBindingConfig

## Example Usage
```hcl
resource "nsxt_vpc_dhcp_v4_static_binding_config" "test-dhcpv4staticbindingconfig" {
    	resource_type = "DhcpV4StaticBindingConfig"
	gateway_address = "10.1.1.1"
	mac_address = "11:22:33:44:55:67"
	host_name = "vm1.vmware.com"
	ip_address = "30.30.30.175"
	options {
		option121 {
			static_routes {
	network = "10.22.12.1/23"
	next_hop = "2.2.2.2"
}
static_routes {
	network = "129.0.0.1/32"
	next_hop = "0.0.0.0"
}
		}
	}
	parent_path = nsxt_vpc_parentResource.resource_name.path
	nsx_id = "test-DhcpStaticBindingConfig-abc"

  }
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `parent_path` - (Required) The policy path of immediate parent resource. This path will be used to create the resource.
* `description` - (Optional) Description of this resource
* `display_name` - (Optional) Defaults to ID if not set
* `gateway_address` - (Optional) When not specified, gateway address is auto-assigned from segment
configuration.

* `host_name` - (Optional) Hostname to assign to the host.

* `ip_address` - (Required) IP assigned to host. The IP address must belong to the subnet, if any,
configured on Segment.

* `lease_time` - (Optional) DHCP lease time in seconds.

* `mac_address` - (Required) MAC address of the host.

* `options` - (Optional) DHCP options for IPv4 server.
  * `option121` - (Optional) 
    * `static_routes` - (Required) Classless static route of DHCP option 121.
  * `others` - (Optional) To define DHCP options other than option 121 in generic format.
Please note, only the following options can be defined in generic
format. Those other options will be accepted without validation
but will not take effect.
--------------------------
  Code    Name
--------------------------
    2   Time Offset
    6   Domain Name Server
    13  Boot File Size
    19  Forward On/Off
    26  MTU Interface
    28  Broadcast Address
    35  ARP Timeout
    40  NIS Domain
    41  NIS Servers
    42  NTP Servers
    44  NETBIOS Name Srv
    45  NETBIOS Dist Srv
    46  NETBIOS Node Type
    47  NETBIOS Scope
    58  Renewal Time
    59  Rebinding Time
    64  NIS+-Domain-Name
    65  NIS+-Server-Addr
    66  TFTP Server-Name (used by PXE)
    67  Bootfile-Name (used by PXE)
    117 Name Service Search
    119 Domain Search
    150 TFTP server address (used by PXE)
    209 PXE Configuration File
    210 PXE Path Prefix
    211 PXE Reboot Time

    * `code` - (Required) Code of the dhcp option.
    * `values` - (Required) Value of the option.
* `resource_type` - (Required) 
* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of DhcpV4StaticBindingConfig resource.

## Importing

An existing DhcpV4StaticBindingConfig can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```hcl
terraform import nsxt_vpc_dhcp_v4_static_binding_config.test-dhcpv4staticbindingconfig ID
```

The above would import NSX `DhcpV4StaticBindingConfig` as a resource named test-dhcpv4staticbindingconfig with the terraform ID `ID`, 
which is the external ID of DhcpV4StaticBindingConfig, with value as full policy path of this resource.
