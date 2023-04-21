---
layout: "nsxt"
page_title: "NSXT: nsxt_policy_vpc_nat_rule"
sidebar_current: "docs-nsxt-policyvpcnatrule"
description: |-
  Get information of NSX-T PolicyVpcNatRule.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_policy_vpc_nat_rule

This data source is used to to get nsxt_policy_vpc_nat_rule objects.

## Example Usage

```hcl
data "nsxt_policy_vpc_nat_rule" "foo_policyvpcnatrule" {
  display_name = "PolicyVpcNatRule-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search PolicyVpcNatRule by its display_name.
* `nsx_id` - (Optional) Search PolicyVpcNatRule by its NSX ID.
* `parent_path` - (Optional) Search PolicyVpcNatRule by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the PolicyVpcNatRule.
* `path` - Full policy path of the PolicyVpcNatRule.

