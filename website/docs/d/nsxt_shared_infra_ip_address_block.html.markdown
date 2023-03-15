---
layout: "nsxt"
page_title: "NSXT: nsxt_shared_infra_ip_address_block"
sidebar_current: "docs-nsxt-shared-infraipaddressblock"
description: |-
  Get information of NSX-T InfraIpAddressBlock.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_shared_infra_ip_address_block

This data source is used to to get nsxt_shared_infra_ip_address_block objects.

**Note:**
* `Infra shared` means that the entity represented by this datasource is shared with all the Orgs, Projects, and VPCs that the Share container references.
* `Project shared` means the entity represented by this datasource is shared with all VPCs in this project. 

## Example Usage

```hcl
data "nsxt_shared_infra_ip_address_block" "foo_infraipaddressblock" {
  display_name = "InfraIpAddressBlock-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search InfraIpAddressBlock by its display_name.
* `nsx_id` - (Optional) Search InfraIpAddressBlock by its NSX ID.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the InfraIpAddressBlock.
* `path` - Full policy path of the InfraIpAddressBlock.

