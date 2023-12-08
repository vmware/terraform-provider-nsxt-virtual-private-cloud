---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_ip_address_allocation"
sidebar_current: "docs-nsxt-vpcipaddressallocation"
description: |-
  Get information of NSX-T VpcIpAddressAllocation.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_ip_address_allocation

This data source is used to to get nsxt_vpc_ip_address_allocation objects.

## Example Usage

```hcl
data "nsxt_vpc_ip_address_allocation" "foo_vpcipaddressallocation" {
  display_name = "VpcIpAddressAllocation-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search VpcIpAddressAllocation by its display_name.
* `nsx_id` - (Optional) Search VpcIpAddressAllocation by its NSX ID.
* `parent_path` - (Optional) Search VpcIpAddressAllocation by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the VpcIpAddressAllocation.
* `path` - Full policy path of the VpcIpAddressAllocation.
* `allocation_ip` - The allocated IP address.

