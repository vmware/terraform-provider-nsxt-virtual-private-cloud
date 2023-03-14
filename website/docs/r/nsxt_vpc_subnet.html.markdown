---
layout: "nsxt"
page_title: "Nsxt: nsxt_vpc_subnet"
sidebar_current: "docs-nsxt-vpcsubnet"
description: |-
  Creates and manages VpcSubnet.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

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
* `dhcp_config` - (Optional) Vpc Subnet DHCP config

  * `dhcp_relay_config_path` - (Optional) Policy path of DHCP-relay-config. If configured then all the subnets will be configured with the DHCP relay server.
If not specified, then the local DHCP server will be configured for all connected subnets.

  * `dns_client_config` - (Optional) 
    * `dns_server_ips` - (Optional) IPs of the DNS servers which need to be configured on teh workload VMs

  * `static_pool_config` - (Optional) 
    * `ipv4_pool_size` - (Optional) Number of IPs to be reserved in static ip pool.
By default, if dhcp is enabled then static ipv4 pool size will be zero and all available IPs will be reserved in
local dhcp pool.
Maximum allowed value is 'subnet size - 4'. Configure maximum value if dhcp pool is not required.

  * `enable_dhcp` - (Optional) If enabled, the DHCP server will be configured based on IP address type.
If disabled then neither DHCP server nor relay shall be configured.

* `skip_ipam` - (Optional) Temperory workaround, will be removed once ipam integration is merged.

* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `access_mode` - (Optional) There are three kinds of Access Types supported for an Application.
Private  - Subnet is accessbile only within the application and its IPs are allocated from
           private IP address pool from VPC configuration unless specified explicitly by user.
Public   - Subnet is accessible from external networks and its IPs are allocated from public IP
           address pool from VPC configuration unless specified explicitly by user.
Isolated - Subnet is not accessible from other subnets within the same VPC.

* `ipv4_subnet_size` - (Optional) If IP Addresses are not provided, this field will be used to carve out the ips
from respective ip block defined in the parent vpc. The default is 64.
If ip_addresses field is provided then ipv4_subnet_size field is ignored.
This field cannot be modified after creating a subnet.

* `ip_addresses` - (Optional) If not provided, Ip assignment will be done based on VPC CIDRs
This represents the subnet that is associated with tier.
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

```hcl
terraform import nsxt_vpc_subnet.test-vpcsubnet ID
```

The above would import NSX `VpcSubnet` as a resource named test-vpcsubnet with the terraform ID `ID`, 
which is the external ID of VpcSubnet, with value as full policy path of this resource.
