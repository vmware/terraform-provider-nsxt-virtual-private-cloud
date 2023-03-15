---
layout: "nsxt"
page_title: "NSXT: nsxt_shared_project_infra_service"
sidebar_current: "docs-nsxt-shared-projectinfraservice"
description: |-
  Get information of NSX-T ProjectInfraService.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_shared_project_infra_service

This data source is used to to get nsxt_shared_project_infra_service objects.

**Note:**
* `Infra shared` means that the entity represented by this datasource is shared with all the Orgs, Projects, and VPCs that the Share container references.
* `Project shared` means the entity represented by this datasource is shared with all VPCs in this project. 

## Example Usage

```hcl
data "nsxt_shared_project_infra_service" "foo_projectinfraservice" {
  display_name = "ProjectInfraService-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search ProjectInfraService by its display_name.
* `nsx_id` - (Optional) Search ProjectInfraService by its NSX ID.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the ProjectInfraService.
* `path` - Full policy path of the ProjectInfraService.

