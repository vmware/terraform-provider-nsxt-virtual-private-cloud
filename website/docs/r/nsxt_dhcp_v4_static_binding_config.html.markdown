---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_dhcp_v4_static_binding_config"
sidebar_current: "docs-nsxt-vpc-dhcpv4staticbindingconfig"
description: |-
  Creates and manages DhcpV4StaticBindingConfig.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_dhcp_v4_static_binding_config

The DhcpV4StaticBindingConfig resource allows the creation and management of Nsxt DhcpV4StaticBindingConfig

## Example Usage

```hcl
resource "nsxt_vpc_dhcp_v4_static_binding_config" "test-dhcpv4staticbindingconfig" {
  gateway_address = "10.1.1.1"
  parent_path     = nsxt_vpc_parentResource.resource_name.path
  nsx_id          = "test-DhcpStaticBindingConfig-abc"
  host_name       = "vm1.vmware.com"
  mac_address     = "11:22:33:44:55:67"
  ip_address      = "30.30.30.175"
  options {
    option121 {
      static_routes {
        next_hop = "2.2.2.2"
        network  = "10.22.12.1/23"
      }
      static_routes {
        next_hop = "0.0.0.0"
        network  = "129.0.0.1/32"
      }
    }
  }
  resource_type = "DhcpV4StaticBindingConfig"
}
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `parent_path` - (Required) The policy path of immediate parent resource. This path will be used to create the resource.
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
* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `host_name` - (Optional) Hostname to assign to the host.

* `gateway_address` - (Optional) When not specified, gateway address is auto-assigned from segment
configuration.

* `ip_address` - (Required) IP assigned to host. The IP address must belong to the subnet, if any,
configured on Segment.

* `mac_address` - (Required) MAC address of the host.

* `resource_type` - (Required) Resource type of the DhcpStaticBindingConfig

* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters
* `lease_time` - (Optional) DHCP lease time in seconds.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of DhcpV4StaticBindingConfig resource.

## Importing

An existing DhcpV4StaticBindingConfig can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```
terraform import nsxt_vpc_dhcp_v4_static_binding_config.test-dhcpv4staticbindingconfig ID
```

The above would import NSX `DhcpV4StaticBindingConfig` as a resource named test-dhcpv4staticbindingconfig with the terraform ID `ID`, 
which is the external ID of DhcpV4StaticBindingConfig, with value as full policy path of this resource.
