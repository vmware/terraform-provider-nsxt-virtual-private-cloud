---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_policy_nat_rule"
sidebar_current: "docs-nsxt-vpc-policynatrule"
description: |-
  Get information of NSX-T PolicyNatRule.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# nsxt_vpc_policy_nat_rule

This data source is used to to get nsxt_vpc_policy_nat_rule objects.

## Example Usage

```hcl
data "nsxt_vpc_policy_nat_rule" "foo_policynatrule" {
  display_name = "PolicyNatRule-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search PolicyNatRule by its display_name.
* `nsx_id` - (Optional) Search PolicyNatRule by its NSX ID.
* `parent_path` - (Optional) Search PolicyNatRule by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the PolicyNatRule.
* `path` - Full policy path of the PolicyNatRule.

