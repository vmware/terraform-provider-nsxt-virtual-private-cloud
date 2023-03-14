---
layout: "nsxt"
page_title: "NSXT: nsxt_shared_infra_l2_bridge_endpoint_profile"
sidebar_current: "docs-nsxt-shared-infral2bridgeendpointprofile"
description: |-
  Get information of NSX-T InfraL2BridgeEndpointProfile.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# nsxt_shared_infra_l2_bridge_endpoint_profile

This data source is used to to get nsxt_shared_infra_l2_bridge_endpoint_profile objects.

**Note:**
* `Infra shared` means that the entity represented by this datasource is shared with all the Orgs, Projects, and VPCs that the Share container references.
* `Project shared` means the entity represented by this datasource is shared with all VPCs in this project. 

## Example Usage

```hcl
data "nsxt_shared_infra_l2_bridge_endpoint_profile" "foo_infral2bridgeendpointprofile" {
  display_name = "InfraL2BridgeEndpointProfile-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search InfraL2BridgeEndpointProfile by its display_name.
* `nsx_id` - (Optional) Search InfraL2BridgeEndpointProfile by its NSX ID.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the InfraL2BridgeEndpointProfile.
* `path` - Full policy path of the InfraL2BridgeEndpointProfile.

