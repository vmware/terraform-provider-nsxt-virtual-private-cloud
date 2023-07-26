---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_service"
sidebar_current: "docs-nsxt-vpc-service"
description: |-
  Get information of NSX-T Service.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_service

This data source is used to to get nsxt_vpc_service objects.

## Example Usage

```hcl
data "nsxt_vpc_service" "foo_service" {
  context {
    scope = "vpc"
    }
  display_name = "service-Test"
}
```

## Argument Reference

* `display_name` - (Optional) Search Service by its display_name.
* `nsx_id` - (Optional) Search Service by its NSX ID.
* `context` - (Required) Provide context information for Service.
  * `scope` - (Required) Provide scope for searching the Service. It can be any one of vpc, project or infra. Defaulted to vpc if not provided.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Unique identifier for this entity in NSX.
* `description` - Description of the Service.
* `path` - Full policy path of the Service.

