---
layout: "nsxt"
page_title: "NSXT: nsxt_policy_vpc_vm"
sidebar_current: "docs-nsxt-datasource-vpc-vm"
description: |-
  Get information of NSX-T VPC VM.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# nsxt_policy_vpc_vm

This data source is used to to get nsxt_vm objects.

## Example Usage

```hcl
data "nsxt_policy_vpc_vm" "foo_vpc_vm" {
  display_name = "VM-Test"
}
```

OR

```hcl
data "nsxt_policy_vpc_vm" "foo_vpc_vm" {
  external_id = "vm-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
}
```

## Argument Reference

* `display_name` - (Optional) Search VM by its display_name.
* `external_id` - (Optional) Search VM by its external ID.
* `power_state` - (Optional) Search VM by its power state.

