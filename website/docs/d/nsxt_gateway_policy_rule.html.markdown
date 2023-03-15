---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_gateway_policy_rule"
sidebar_current: "docs-nsxt-vpc-gatewaypolicyrule"
description: |-
  Get information of NSX-T GatewayPolicyRule.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_gateway_policy_rule

This data source is used to to get nsxt_vpc_gateway_policy_rule objects.

## Example Usage

```hcl
data "nsxt_vpc_gateway_policy_rule" "foo_gatewaypolicyrule" {
  display_name = "GatewayPolicyRule-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search GatewayPolicyRule by its display_name.
* `nsx_id` - (Optional) Search GatewayPolicyRule by its NSX ID.
* `parent_path` - (Optional) Search GatewayPolicyRule by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the GatewayPolicyRule.
* `path` - Full policy path of the GatewayPolicyRule.

