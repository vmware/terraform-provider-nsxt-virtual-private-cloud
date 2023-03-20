---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_dhcp_v6_static_binding_config"
sidebar_current: "docs-nsxt-vpc-dhcpv6staticbindingconfig"
description: |-
  Get information of NSX-T DhcpV6StaticBindingConfig.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_dhcp_v6_static_binding_config

This data source is used to to get nsxt_vpc_dhcp_v6_static_binding_config objects.

## Example Usage

```hcl
data "nsxt_vpc_dhcp_v6_static_binding_config" "foo_dhcpv6staticbindingconfig" {
  display_name = "DhcpV6StaticBindingConfig-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search DhcpV6StaticBindingConfig by its display_name.
* `nsx_id` - (Optional) Search DhcpV6StaticBindingConfig by its NSX ID.
* `parent_path` - (Optional) Search DhcpV6StaticBindingConfig by its parent's path.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the DhcpV6StaticBindingConfig.
* `path` - Full policy path of the DhcpV6StaticBindingConfig.

