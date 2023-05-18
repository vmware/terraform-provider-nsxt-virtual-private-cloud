---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_subnet_ip_address_allocation"
sidebar_current: "docs-nsxt-vpc-subnet-ipaddressallocation"
description: |-
  Creates and manages IpAddressAllocation.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_subnet_ip_address_allocation

The IpAddressAllocation resource allows the creation and management of Nsxt IpAddressAllocation

## Example Usage
```hcl
resource "nsxt_vpc_subnet_ip_address_allocation" "test-ipaddressallocation" {
    	parent_path = nsxt_vpc_parentResource.resource_name.path
	nsx_id = "test-IpAddressAllocation-abc"
	allocation_ip = "192.168.0.6"

  }
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `parent_path` - (Required) The policy path of immediate parent resource. This path will be used to create the resource.
* `resource_type` - (Optional) The type of this resource.
* `sync_realization` - (Optional) Realization of intent will be called synchronously

* `allocation_ip` - (Optional) Address that is allocated from pool
* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of IpAddressAllocation resource.

## Importing

An existing IpAddressAllocation can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```
terraform import nsxt_vpc_subnet_ip_address_allocation.test-ipaddressallocation ID
```

The above would import NSX `IpAddressAllocation` as a resource named test-ipaddressallocation with the terraform ID `ID`, 
which is the external ID of IpAddressAllocation, with value as full policy path of this resource.
