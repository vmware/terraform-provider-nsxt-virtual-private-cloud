---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_ip_address_block"
sidebar_current: "docs-nsxt-vpc-ipaddressblock"
description: |-
  Get information of NSX-T IpAddressBlock.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_ip_address_block

This data source is used to to get nsxt_vpc_ip_address_block objects.

## Example Usage

```hcl
data "nsxt_vpc_ip_address_block" "foo_ip_address_block" {
  context_info {
    context = "vpc"
    }
  display_name = "ip_address_block-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search IpAddressBlock by its display_name.
* `nsx_id` - (Optional) Search IpAddressBlock by its NSX ID.
* `context_info` - (Optional) Provide context information for IpAddressBlock.
  * `context` - (Optional) Provide context for searching the IpAddressBlock. It can be any one of vpc, project or infra. Defaulted to vpc if not provided.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the IpAddressBlock.
* `path` - Full policy path of the IpAddressBlock.

