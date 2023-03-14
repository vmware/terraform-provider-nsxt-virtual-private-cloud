---
layout: "nsxt"
page_title: "NSXT: nsxt_shared_project_infra_ip_address_pool"
sidebar_current: "docs-nsxt-shared-projectinfraipaddresspool"
description: |-
  Get information of NSX-T ProjectInfraIpAddressPool.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# nsxt_shared_project_infra_ip_address_pool

This data source is used to to get nsxt_shared_project_infra_ip_address_pool objects.

**Note:**
* `Infra shared` means that the entity represented by this datasource is shared with all the Orgs, Projects, and VPCs that the Share container references.
* `Project shared` means the entity represented by this datasource is shared with all VPCs in this project. 

## Example Usage

```hcl
data "nsxt_shared_project_infra_ip_address_pool" "foo_projectinfraipaddresspool" {
  display_name = "ProjectInfraIpAddressPool-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search ProjectInfraIpAddressPool by its display_name.
* `nsx_id` - (Optional) Search ProjectInfraIpAddressPool by its NSX ID.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the ProjectInfraIpAddressPool.
* `path` - Full policy path of the ProjectInfraIpAddressPool.

