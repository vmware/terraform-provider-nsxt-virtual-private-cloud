---
layout: "nsxt"
page_title: "NSXT: nsxt_shared_infra_group"
sidebar_current: "docs-nsxt-shared-infragroup"
description: |-
  Get information of NSX-T InfraGroup.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# nsxt_shared_infra_group

This data source is used to to get nsxt_shared_infra_group objects.

**Note:**
* `Infra shared` means that the entity represented by this datasource is shared with all the Orgs, Projects, and VPCs that the Share container references.
* `Project shared` means the entity represented by this datasource is shared with all VPCs in this project. 

## Example Usage

```hcl
data "nsxt_shared_infra_group" "foo_infragroup" {
  display_name = "InfraGroup-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search InfraGroup by its display_name.
* `nsx_id` - (Optional) Search InfraGroup by its NSX ID.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the InfraGroup.
* `path` - Full policy path of the InfraGroup.

