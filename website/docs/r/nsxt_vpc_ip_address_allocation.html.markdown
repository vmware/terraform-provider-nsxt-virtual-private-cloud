---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_ip_address_allocation"
sidebar_current: "docs-nsxt-vpcipaddressallocation"
description: |-
  Creates and manages VpcIpAddressAllocation.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_ip_address_allocation

The VpcIpAddressAllocation resource allows the creation and management of Nsxt VpcIpAddressAllocation

## Example Usage
```hcl
resource "nsxt_vpc_ip_address_allocation" "test-vpcipaddressallocation" {
    	nsx_id = "test-VpcIpAddressAllocation-abc"
	allocation_ip = "192.168.0.6"
	ip_address_block_visibility = "EXTERNAL"

  }
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `resource_type` - (Optional) The type of this resource.
* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `allocation_ip` - (Optional) Single IP Address that is allocated from external ip block or IPv6 block based on IP address type. If not
specified, any available IP will be allocated from respective IP block. If specified, it has to be within range of respective
IP blocks. If IP is already in use then validation error will be thrown.

* `ip_address_block_visibility` - (Optional) Represents visibility of IP address block. This field is not applicable if IPAddressType at VPC is IPv6.

* `ip_address_type` - (Optional) This defines the type of IP address block that will be used to allocate IP. This field is applicable only
if IP addressType at VPC is DUAL. In case of IPv4, external blocks will be used, and in case of IPv6, IPv6 blocks will be used.

* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of VpcIpAddressAllocation resource.

## Importing

An existing VpcIpAddressAllocation can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```hcl
terraform import nsxt_vpc_ip_address_allocation.test-vpcipaddressallocation ID
```

The above would import NSX `VpcIpAddressAllocation` as a resource named test-vpcipaddressallocation with the terraform ID `ID`, 
which is the external ID of VpcIpAddressAllocation, with value as full policy path of this resource.
