---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_vm_tags"
sidebar_current: "docs-nsxt-vpc-resource-vpcvmtags"
description: |-
  Creates and manages Vpc VM Tags.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# nsxt_vpc_vm_tags

This resource allows the creation and management of tags on VMs in scope of VPC.

## Example Usage
```hcl
resource "nsxt_vpc_vm_tags" "test-vpc-vm-tags" {
  virtual_machine_id = "1e10a422-ad92-4642-b56d-47971848cc3d"
  tags {
    scope = "os"
    tag = "windows"
  }
}
```
or you can provide reference to the VM using its datasource

```hcl
data "nsxt_policy_vpc_vm" "foo_vpc_vm" {
  external_id = "vm-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
  display_name = "Dev-VM-Test"
}

resource "nsxt_vpc_vm_tags" "test-vpc-vm-tags" {
  virtual_machine_id = nsxt_policy_vpc_vm.foo_vpc_vm.external_id
  tags {
    scope = "os"
    tag = "windows"
  }
}
```

## Argument Reference

The following arguments are supported:

* `virtual_machine_id` - (Required) This is the external ID of the Virtual machine on which tags will be managed.
* `tags` - (Optional) A list of scope + tag pairs to associate with this Virtual Machine.
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters

## Importing

An existing VpcVmTags collection can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```hcl
terraform import nsxt_vpc_vm_tags.test-vpc-vm-tags ID
```

The above would import NSX Virtual Machine tags as a resource named test-vpc-vm-tags with the NSX ID ID, where ID is external ID of the Virtual Machine.