---
layout: "nsxt"
page_title: "NSXT: nsxt_shared_project_infra_policy_context_profile"
sidebar_current: "docs-nsxt-shared-projectinfrapolicycontextprofile"
description: |-
  Get information of NSX-T ProjectInfraPolicyContextProfile.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# nsxt_shared_project_infra_policy_context_profile

This data source is used to to get nsxt_shared_project_infra_policy_context_profile objects.

**Note:**
* `Infra shared` means that the entity represented by this datasource is shared with all the Orgs, Projects, and VPCs that the Share container references.
* `Project shared` means the entity represented by this datasource is shared with all VPCs in this project. 

## Example Usage

```hcl
data "nsxt_shared_project_infra_policy_context_profile" "foo_projectinfrapolicycontextprofile" {
  display_name = "ProjectInfraPolicyContextProfile-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search ProjectInfraPolicyContextProfile by its display_name.
* `nsx_id` - (Optional) Search ProjectInfraPolicyContextProfile by its NSX ID.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the ProjectInfraPolicyContextProfile.
* `path` - Full policy path of the ProjectInfraPolicyContextProfile.

