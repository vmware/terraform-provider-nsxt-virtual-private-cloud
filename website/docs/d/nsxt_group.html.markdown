---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_group"
sidebar_current: "docs-nsxt-vpc-group"
description: |-
  Get information of NSX-T Group.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_group

This data source is used to to get nsxt_vpc_group objects.

## Example Usage

```hcl
data "nsxt_vpc_group" "foo_group" {
  context {
    scope  = "vpc"
    domain = "default"
  }
  display_name = "group-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search Group by its display_name.
* `nsx_id` - (Optional) Search Group by its NSX ID.
* `context` - (Required) Provide context information for Group.
  * `scope` - (Required) Provide scope for searching the Group. It can be any one of vpc, project or infra. Defaulted to vpc if not provided.
  * `domain` - (Optional) Domain ID to search the Group in a particular domain.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the Group.
* `path` - Full policy path of the Group.

