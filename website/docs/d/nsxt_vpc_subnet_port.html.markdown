---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_subnet_port"
sidebar_current: "docs-nsxt-vpcsubnetport"
description: |-
  Get information of NSX-T VpcSubnetPort.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_subnet_port

This data source is used to to get nsxt_vpc_subnet_port objects.

## Example Usage

```hcl
data "nsxt_vpc_subnet_port" "foo_vpcsubnetport" {
  display_name = "VpcSubnetPort-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search VpcSubnetPort by its display_name.
* `nsx_id` - (Optional) Search VpcSubnetPort by its NSX ID.
* `parent_path` - (Optional) Search VpcSubnetPort by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the VpcSubnetPort.
* `path` - Full policy path of the VpcSubnetPort.

