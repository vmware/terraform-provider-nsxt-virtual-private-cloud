---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_ip_address_pool"
sidebar_current: "docs-nsxt-vpc-ipaddresspool"
description: |-
  Get information of NSX-T IpAddressPool.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_ip_address_pool

This data source is used to to get nsxt_vpc_ip_address_pool objects.

## Example Usage

```hcl
data "nsxt_vpc_ip_address_pool" "foo_ip_address_pool" {
  context {
    scope = "vpc"
    }
  display_name = "ip_address_pool-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search IpAddressPool by its display_name.
* `nsx_id` - (Optional) Search IpAddressPool by its NSX ID.
* `context` - (Required) Provide context information for IpAddressPool.
  * `scope` - (Required) Provide scope for searching the IpAddressPool. It can be any one of vpc, project or infra. Defaulted to vpc if not provided.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the IpAddressPool.
* `path` - Full policy path of the IpAddressPool.

