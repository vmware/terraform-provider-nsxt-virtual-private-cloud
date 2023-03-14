---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_static_routes"
sidebar_current: "docs-nsxt-vpc-staticroutes"
description: |-
  Get information of NSX-T StaticRoutes.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# nsxt_vpc_static_routes

This data source is used to to get nsxt_vpc_static_routes objects.

## Example Usage

```hcl
data "nsxt_vpc_static_routes" "foo_staticroutes" {
  display_name = "StaticRoutes-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search StaticRoutes by its display_name.
* `nsx_id` - (Optional) Search StaticRoutes by its NSX ID.
* `parent_path` - (Optional) Search StaticRoutes by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the StaticRoutes.
* `path` - Full policy path of the StaticRoutes.

