---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_subnet"
sidebar_current: "docs-nsxt-vpcsubnet"
description: |-
  Get information of NSX-T VpcSubnet.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_subnet

This data source is used to to get nsxt_vpc_subnet objects.

## Example Usage

```hcl
data "nsxt_vpc_subnet" "foo_vpcsubnet" {
  display_name = "VpcSubnet-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search VpcSubnet by its display_name.
* `nsx_id` - (Optional) Search VpcSubnet by its NSX ID.
* `parent_path` - (Optional) Search VpcSubnet by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the VpcSubnet.
* `path` - Full policy path of the VpcSubnet.

