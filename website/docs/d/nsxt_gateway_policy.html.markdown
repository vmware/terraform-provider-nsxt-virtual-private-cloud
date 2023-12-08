---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_gateway_policy"
sidebar_current: "docs-nsxt-vpc-gatewaypolicy"
description: |-
  Get information of NSX-T GatewayPolicy.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_gateway_policy

This data source is used to to get nsxt_vpc_gateway_policy objects.

## Example Usage

```hcl
data "nsxt_vpc_gateway_policy" "foo_gatewaypolicy" {
  display_name = "GatewayPolicy-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search GatewayPolicy by its display_name.
* `nsx_id` - (Optional) Search GatewayPolicy by its NSX ID.
* `parent_path` - (Optional) Search GatewayPolicy by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the GatewayPolicy.
* `path` - Full policy path of the GatewayPolicy.

