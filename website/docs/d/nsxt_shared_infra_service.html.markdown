---
layout: "nsxt"
page_title: "NSXT: nsxt_shared_infra_service"
sidebar_current: "docs-nsxt-shared-infraservice"
description: |-
  Get information of NSX-T InfraService.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_shared_infra_service

This data source is used to to get nsxt_shared_infra_service objects.

**Note:**
* `Infra shared` means that the entity represented by this datasource is shared with all the Orgs, Projects, and VPCs that the Share container references.
* `Project shared` means the entity represented by this datasource is shared with all VPCs in this project. 

## Example Usage

```hcl
data "nsxt_shared_infra_service" "foo_infraservice" {
  display_name = "InfraService-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search InfraService by its display_name.
* `nsx_id` - (Optional) Search InfraService by its NSX ID.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the InfraService.
* `path` - Full policy path of the InfraService.

