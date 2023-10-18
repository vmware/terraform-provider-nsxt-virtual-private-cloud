---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_static_routes"
sidebar_current: "docs-nsxt-vpc-staticroutes"
description: |-
  Creates and manages StaticRoutes.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_static_routes

The StaticRoutes resource allows the creation and management of Nsxt StaticRoutes

## Example Usage

```
resource "nsxt_vpc_static_routes" "test-staticroutes" {
  nsx_id = "test-StaticRoutes-abc"
  next_hops {
    ip_address = "41.1.1.1"
    admin_distance = 1
  }
  next_hops {
    ip_address = "42.1.1.1"
    admin_distance = 2
  }
  next_hops {
    ip_address = "43.1.2.3"
    admin_distance = 3
  }
  network = "45.1.1.0/24"
}
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `resource_type` - (Optional) The type of this resource.
* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `network` - (Required) Specify network address in CIDR format.
In case of VPC, user can optionally use allocated IP from one of the external blocks associated with VPC.
Only /32 CIDR is allowed in case IP overlaps with external blocks.

* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters
* `next_hops` - (Required) Specify next hop routes for network.

  * `scope` - (Optional) Interface path associated with current route.
For example: specify a policy path referencing the IPSec VPN Session.
Should not be provided while creating routes under VPC.

  * `ip_address` - (Optional) Next hop gateway IP address
  * `admin_distance` - (Optional) Cost associated with next hop route

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of StaticRoutes resource.

## Importing

An existing StaticRoutes can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```
terraform import nsxt_vpc_static_routes.test-staticroutes ID
```

The above would import NSX `StaticRoutes` as a resource named test-staticroutes with the terraform ID `ID`, 
which is the external ID of StaticRoutes, with value as full policy path of this resource.
