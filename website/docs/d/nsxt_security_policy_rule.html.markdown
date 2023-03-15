---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_security_policy_rule"
sidebar_current: "docs-nsxt-vpc-securitypolicyrule"
description: |-
  Get information of NSX-T SecurityPolicyRule.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_security_policy_rule

This data source is used to to get nsxt_vpc_security_policy_rule objects.

## Example Usage

```hcl
data "nsxt_vpc_security_policy_rule" "foo_securitypolicyrule" {
  display_name = "SecurityPolicyRule-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search SecurityPolicyRule by its display_name.
* `nsx_id` - (Optional) Search SecurityPolicyRule by its NSX ID.
* `parent_path` - (Optional) Search SecurityPolicyRule by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the SecurityPolicyRule.
* `path` - Full policy path of the SecurityPolicyRule.

