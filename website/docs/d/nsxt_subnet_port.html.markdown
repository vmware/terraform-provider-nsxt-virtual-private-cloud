---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_subnet_port"
sidebar_current: "docs-nsxt-vpc-subnetport"
description: |-
  Get information of NSX-T SubnetPort.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# nsxt_vpc_subnet_port

This data source is used to to get nsxt_vpc_subnet_port objects.

## Example Usage

```hcl
data "nsxt_vpc_subnet_port" "foo_subnetport" {
  display_name = "SubnetPort-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search SubnetPort by its display_name.
* `nsx_id` - (Optional) Search SubnetPort by its NSX ID.
* `parent_path` - (Optional) Search SubnetPort by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the SubnetPort.
* `path` - Full policy path of the SubnetPort.

