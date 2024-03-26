---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_subnet"
sidebar_current: "docs-nsxt-vpcsubnet"
description: |-
  Creates and manages VpcSubnet.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_subnet

The VpcSubnet resource allows the creation and management of Nsxt VpcSubnet

## Example Usage
```hcl
resource "nsxt_vpc_subnet" "test-vpcsubnet" {
    	display_name = "Subnet 1"
	description = "This is test VpcSubnet"
	ip_addresses = ["10.112.2.0/24"]
	nsx_id = "test-VpcSubnet-abc"
	ipv4_subnet_size = 64
	access_mode = "Private"

  }
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `resource_type` - (Optional) The type of this resource.
* `advanced_config` - (Optional) VPC Subnet Advanced Configuration

  * `static_ip_allocation` - (Optional) 
    * `enabled` - (Optional) Enable ip and mac addresse allocation for VPC Subnet ports from static ip pool. To enable this,
dhcp pool shall be empty and static ip pool shall own all available ip addresses.

* `dhcp_config` - (Optional) VPC Subnet DHCP config

  * `dhcp_relay_config_path` - (Optional) Policy path of DHCP-relay-config. If configured then all the subnets will be configured with the DHCP relay server.
If not specified, then the local DHCP server will be configured for all connected subnets.

  * `dns_client_config` - (Optional) 
    * `dns_server_ips` - (Optional) IPs of the DNS servers which need to be configured on the workload VMs

  * `static_pool_config` - (Optional) 
    * `ipv4_pool_size` - (Optional) Number of IPs to be reserved in static ip pool. Maximum allowed value is 'subnet size - 4'.
If dhcp is enabled then by default static ipv4 pool size will be zero and all available IPs will be reserved in
local dhcp pool.
If dhcp is deactivated then by default all IPs will be reserved in static ip pool.

  * `enable_dhcp` - (Optional) If activated, the DHCP server will be configured based on IP address type.
If deactivated then neither DHCP server nor relay shall be configured.

* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `access_mode` - (Optional) There are three kinds of Access Types supported for an Application.
Private  - VPC Subnet is accessible only within the application and its IPs are allocated from
           private IP address pool from VPC configuration unless specified explicitly by user.
Public   - VPC Subnet is accessible from external networks and its IPs are allocated from public IP
           address pool from VPC configuration unless specified explicitly by user.
Isolated - VPC Subnet is not accessible from other VPC Subnets within the same VPC.

* `ipv4_subnet_size` - (Optional) If IP Addresses are not provided, this field will be used to carve out the ips
from respective ip block defined in the parent VPC. The default is 64.
If ip_addresses field is provided then ipv4_subnet_size field is ignored.
This field cannot be modified after creating a VPC Subnet.

* `ip_addresses` - (Optional) If not provided, Ip assignment will be done based on VPC CIDRs
This represents the VPC Subnet that is associated with tier.
If IPv4 CIDR is given, ipv4_subnet_size property is ignored.
For IPv6 CIDR, supported prefix length is /64.

* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of VpcSubnet resource.

## Importing

An existing VpcSubnet can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```
terraform import nsxt_vpc_subnet.test-vpcsubnet ID
```

The above would import NSX `VpcSubnet` as a resource named test-vpcsubnet with the terraform ID `ID`, 
which is the external ID of VpcSubnet, with value as full policy path of this resource.
