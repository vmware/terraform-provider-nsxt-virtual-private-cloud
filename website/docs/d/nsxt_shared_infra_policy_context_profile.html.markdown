---
layout: "nsxt"
page_title: "NSXT: nsxt_shared_infra_policy_context_profile"
sidebar_current: "docs-nsxt-shared-infrapolicycontextprofile"
description: |-
  Get information of NSX-T InfraPolicyContextProfile.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_shared_infra_policy_context_profile

This data source is used to to get nsxt_shared_infra_policy_context_profile objects.

**Note:**
* `Infra shared` means that the entity represented by this datasource is shared with all the Orgs, Projects, and VPCs that the Share container references.
* `Project shared` means the entity represented by this datasource is shared with all VPCs in this project. 

## Example Usage

```hcl
data "nsxt_shared_infra_policy_context_profile" "foo_infrapolicycontextprofile" {
  display_name = "InfraPolicyContextProfile-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search InfraPolicyContextProfile by its display_name.
* `nsx_id` - (Optional) Search InfraPolicyContextProfile by its NSX ID.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the InfraPolicyContextProfile.
* `path` - Full policy path of the InfraPolicyContextProfile.

