---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_group"
sidebar_current: "docs-nsxt-vpc-group"
description: |-
  Creates and manages Group.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_group

The Group resource allows the creation and management of Nsxt Group

## Example Usage
```hcl
resource "nsxt_vpc_group" "test-group" {
    	expression {
	expressions {
	key = "Name"
	operator = "CONTAINS"
	resource_type = "Condition"
	value = "vm_1"
	member_type = "VirtualMachine"
}
expressions {
	conjunction_operator = "AND"
	resource_type = "ConjunctionOperator"
}
expressions {
	key = "Tag"
	operator = "EQUALS"
	resource_type = "Condition"
	value = "London"
	member_type = "VirtualMachine"
}
	resource_type = "NestedExpression"
	tags {
	scope = "scope1"
	tag = "webvm"
}
}
expression {
	conjunction_operator = "OR"
	resource_type = "ConjunctionOperator"
}
expression {
	ip_addresses = ["10.112.10.1"]
	resource_type = "IPAddressExpression"
}
expression {
	conjunction_operator = "OR"
	resource_type = "ConjunctionOperator"
}
expression {
	paths = ["/orgs/default/projects/project-1/vpcs/vpc-1/groups/default"]
	resource_type = "PathExpression"
}
	nsx_id = "test-Group-abc"
	display_name = "test-group-abc"
	description = "Group description"

  }
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `resource_type` - (Optional) The type of this resource.
* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `state` - (Optional) Realization state of this group
* `expression` - (Optional) The expression list must follow below criteria:
  1. A non-empty expression list, must be of odd size. In a list, with
  indices starting from 0, all non-conjunction expressions must be at
  even indices, separated by a conjunction expression at odd
  indices.
  2. The total of ConditionExpression and NestedExpression in a list
  should not exceed 5.
  3. The total of IPAddressExpression, MACAddressExpression, external
  IDs in an ExternalIDExpression and paths in a PathExpression must not exceed
  500.
  4. Each expression must be a valid Expression. See the definition of
  the Expression type for more information.

  * `display_name` - (Optional) Defaults to ID if not set
  * `description` - (Optional) Description of this resource
  * `tags` - (Optional) Opaque identifiers meaningful to the API user
    * `scope` - (Optional) Tag searches may optionally be restricted by scope
    * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters
  * `resource_type` - (Required) 
* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of Group resource.

## Importing

An existing Group can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```
terraform import nsxt_vpc_group.test-group ID
```

The above would import NSX `Group` as a resource named test-group with the terraform ID `ID`, 
which is the external ID of Group, with value as full policy path of this resource.
