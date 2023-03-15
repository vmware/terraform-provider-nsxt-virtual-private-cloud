---
layout: "nsxt"
page_title: "Nsxt: nsxt_vpc_group"
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
    	display_name = "test-group-abc"
	description = "Group description"
	expression {
	resource_type = "NestedExpression"
	expressions {
	member_type = "VirtualMachine"
	value = "vm_1"
	key = "Name"
	operator = "CONTAINS"
	resource_type = "Condition"
}
expressions {
	resource_type = "ConjunctionOperator"
	conjunction_operator = "AND"
}
expressions {
	member_type = "VirtualMachine"
	value = "London"
	key = "Tag"
	operator = "EQUALS"
	resource_type = "Condition"
}
	tags {
	scope = "scope1"
	tag = "webvm"
}
}
expression {
	resource_type = "ConjunctionOperator"
	conjunction_operator = "OR"
}
expression {
	resource_type = "IPAddressExpression"
	ip_addresses = ["10.112.10.1"]
}
expression {
	resource_type = "ConjunctionOperator"
	conjunction_operator = "OR"
}
expression {
	resource_type = "PathExpression"
	paths = ["/orgs/default/projects/project-1/vpcs/vpc-1/groups/default"]
}
	nsx_id = "test-Group-abc"

  }
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `description` - (Optional) Description of this resource
* `display_name` - (Optional) Defaults to ID if not set
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

  * `description` - (Optional) Description of this resource
  * `display_name` - (Optional) Defaults to ID if not set
  * `resource_type` - (Required) 
  * `tags` - (Optional) Opaque identifiers meaningful to the API user
    * `scope` - (Optional) Tag searches may optionally be restricted by scope
    * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters
* `group_type` - (Optional) Group type can be specified during create and update of a group.
Empty group type indicates a 'generic' group, ie group can
include any entity from the valid GroupMemberType.

* `resource_type` - (Optional) The type of this resource.
* `state` - (Optional) Realization state of this group
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

```hcl
terraform import nsxt_vpc_group.test-group ID
```

The above would import NSX `Group` as a resource named test-group with the terraform ID `ID`, 
which is the external ID of Group, with value as full policy path of this resource.
