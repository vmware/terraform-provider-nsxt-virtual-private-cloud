---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_subnet_port"
sidebar_current: "docs-nsxt-vpcsubnetport"
description: |-
  Creates and manages VpcSubnetPort.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_subnet_port

The VpcSubnetPort resource allows the creation and management of Nsxt VpcSubnetPort

## Example Usage

```
resource "nsxt_vpc_subnet_port" "test-vpcsubnetport" {
  address_bindings {
    ip_address = "1.1.1.1"
    mac_address = "aa:bb:cc:dd:ee:ff"
  }
  address_bindings {
    ip_address = "1.1.1.2"
    mac_address = "aa:bb:cc:dd:ee:f1"
  }
  nsx_id = "test-VpcSubnetPort-abc"
  parent_path = nsxt_vpc_parentResource.resource_name.path
}
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `parent_path` - (Required) The policy path of immediate parent resource. This path will be used to create the resource.
* `resource_type` - (Optional) The type of this resource.
* `attachment` - (Optional) Detail information about port attachment
  * `hyperbus_mode` - (Optional) Flag to indicate if hyperbus configuration is required.
  * `context_id` - (Optional) If type is CHILD and the parent port is on the same segment as the child port, then this
field should be VIF ID of the parent port.
If type is CHILD and the parent port is on a different segment, then this
field should be policy path of the parent port.
If type is INDEPENDENT/STATIC, then this field should be transport node ID.

  * `evpn_vlans` - (Optional) List of Evpn tenant VLAN IDs the Parent logical-port serves in Evpn Route-Server mode. Only effective when attachment type is PARENT and the logical-port is attached to vRouter VM.
  * `bms_interface_config` - (Optional) 
    * `routing_table` - (Optional) Routing rules
    * `migrate_intf` - (Optional) IP configuration on migrate_intf will migrate to app_intf_name. It is used for Management and Application sharing the same IP.
    * `app_intf_name` - (Required) The name of application interface
    * `default_gateway` - (Optional) Gateway IP
  * `app_id` - (Optional) ID used to identify/look up a child attachment behind a parent attachment

  * `traffic_tag` - (Optional) Not valid when type field is INDEPENDENT, mainly used to identify
traffic from different ports in container use case.

  * `allocate_addresses` - (Optional) Indicate how IP will be allocated for the port
  * `type` - (Optional) Type of port attachment. STATIC is added to replace INDEPENDENT. INDEPENDENT type and PARENT type are deprecated.
* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `init_state` - (Optional) Set initial state when a new logical port is created. 'UNBLOCKED_VLAN'
means new port will be unblocked on traffic in creation, also VLAN will
be set with corresponding logical switch setting. This port setting
can only be configured at port creation, and cannot be modified.
'RESTORE_VIF' fetches and restores VIF attachment from ESX host.

* `admin_state` - (Optional) Represents desired state of the segment port
* `extra_configs` - (Optional) This property could be used for vendor specific configuration in key value
string pairs. Segment port setting will override segment setting if
the same key was set on both segment and segment port.

  * `config_pair` - (Required) 
    * `value` - (Required) Value
    * `key` - (Required) Key
* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters
* `address_bindings` - (Optional) Static address binding used for the port.
  * `ip_address` - (Optional) IP Address for port binding
  * `vlan_id` - (Optional) VLAN ID for port binding
  * `mac_address` - (Optional) Mac address for port binding
* `ignored_address_bindings` - (Optional) IP Discovery module uses various mechanisms to discover address
bindings being used on each segment port. If a user would like to
ignore any specific discovered address bindings or prevent the
discovery of a particular set of discovered bindings, then those
address bindings can be provided here. Currently IP range in CIDR format
is not supported.

  * `ip_address` - (Optional) IP Address for port binding
  * `vlan_id` - (Optional) VLAN ID for port binding
  * `mac_address` - (Optional) Mac address for port binding

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of VpcSubnetPort resource.

## Importing

An existing VpcSubnetPort can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```
terraform import nsxt_vpc_subnet_port.test-vpcsubnetport ID
```

The above would import NSX `VpcSubnetPort` as a resource named test-vpcsubnetport with the terraform ID `ID`, 
which is the external ID of VpcSubnetPort, with value as full policy path of this resource.
