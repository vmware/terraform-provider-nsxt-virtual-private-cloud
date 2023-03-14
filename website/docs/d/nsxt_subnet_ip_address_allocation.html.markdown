---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_subnet_ip_address_allocation"
sidebar_current: "docs-nsxt-vpc-subnet-ipaddressallocation"
description: |-
  Get information of NSX-T IpAddressAllocation.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# nsxt_vpc_subnet_ip_address_allocation

This data source is used to to get nsxt_vpc_subnet_ip_address_allocation objects.

## Example Usage

```hcl
data "nsxt_vpc_subnet_ip_address_allocation" "foo_ipaddressallocation" {
  display_name = "IpAddressAllocation-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search IpAddressAllocation by its display_name.
* `nsx_id` - (Optional) Search IpAddressAllocation by its NSX ID.
* `parent_path` - (Optional) Search IpAddressAllocation by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the IpAddressAllocation.
* `path` - Full policy path of the IpAddressAllocation.

