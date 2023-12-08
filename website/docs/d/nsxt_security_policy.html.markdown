---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_security_policy"
sidebar_current: "docs-nsxt-vpc-securitypolicy"
description: |-
  Get information of NSX-T SecurityPolicy.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_security_policy

This data source is used to to get nsxt_vpc_security_policy objects.

## Example Usage

```hcl
data "nsxt_vpc_security_policy" "foo_securitypolicy" {
  display_name = "SecurityPolicy-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search SecurityPolicy by its display_name.
* `nsx_id` - (Optional) Search SecurityPolicy by its NSX ID.
* `parent_path` - (Optional) Search SecurityPolicy by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the SecurityPolicy.
* `path` - Full policy path of the SecurityPolicy.

