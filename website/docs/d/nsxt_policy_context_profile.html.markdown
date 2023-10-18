---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_policy_context_profile"
sidebar_current: "docs-nsxt-vpc-policycontextprofile"
description: |-
  Get information of NSX-T PolicyContextProfile.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_policy_context_profile

This data source is used to to get nsxt_vpc_policy_context_profile objects.

## Example Usage

```hcl
data "nsxt_vpc_policy_context_profile" "foo_policy_context_profile" {
  context {
    scope = "vpc"
  }
  display_name = "policy_context_profile-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search PolicyContextProfile by its display_name.
* `nsx_id` - (Optional) Search PolicyContextProfile by its NSX ID.
* `context` - (Required) Provide context information for PolicyContextProfile.
  * `scope` - (Required) Provide scope for searching the PolicyContextProfile. It can be any one of vpc, project or infra. Defaulted to vpc if not provided.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the PolicyContextProfile.
* `path` - Full policy path of the PolicyContextProfile.

