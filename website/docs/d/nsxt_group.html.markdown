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

# nsxt_vpc_group

This data source is used to to get nsxt_vpc_group objects.

## Example Usage

```hcl
data "nsxt_vpc_group" "foo_group" {
  display_name = "Group-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search Group by its display_name.
* `nsx_id` - (Optional) Search Group by its NSX ID.
* `parent_path` - (Optional) Search Group by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the Group.
* `path` - Full policy path of the Group.

