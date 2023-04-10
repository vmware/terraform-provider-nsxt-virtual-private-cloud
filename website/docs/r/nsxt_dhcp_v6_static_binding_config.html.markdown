---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_dhcp_v6_static_binding_config"
sidebar_current: "docs-nsxt-vpc-dhcpv6staticbindingconfig"
description: |-
  Creates and manages DhcpV6StaticBindingConfig.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_dhcp_v6_static_binding_config

The DhcpV6StaticBindingConfig resource allows the creation and management of Nsxt DhcpV6StaticBindingConfig

## Example Usage
```hcl
resource "nsxt_vpc_dhcp_v6_static_binding_config" "test-dhcpv6staticbindingconfig" {
    nsx_id = "test-dhcpv6staticbindingconfig-abc"
	display_name = "DhcpV6StaticBindingConfig-Test"
    parent_path = nsxt_vpc_subnet.subnet-test.path
    }
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `parent_path` - (Required) The policy path of immediate parent resource. This path will be used to create the resource.
* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `resource_type` - (Required) Resource type of the DhcpStaticBindingConfig

* `mac_address` - (Required) The MAC address of the client host. Either client-duid or mac-address,
but not both.

* `domain_names` - (Optional) When not specified, no domain name will be assigned to client host.

* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters
* `preferred_time` - (Optional) Preferred time, in seconds. If this value is not provided, the value
of lease_time*0.8 will be used.

* `lease_time` - (Optional) Lease time, in seconds.
* `sntp_servers` - (Optional) SNTP server IP addresses.
* `dns_nameservers` - (Optional) When not specified, no DNS nameserver will be set to client host.

* `ip_addresses` - (Optional) When not specified, no ip address will be assigned to client host.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of DhcpV6StaticBindingConfig resource.

## Importing

An existing DhcpV6StaticBindingConfig can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```hcl
terraform import nsxt_vpc_dhcp_v6_static_binding_config.test-dhcpv6staticbindingconfig ID
```

The above would import NSX `DhcpV6StaticBindingConfig` as a resource named test-dhcpv6staticbindingconfig with the terraform ID `ID`, 
which is the external ID of DhcpV6StaticBindingConfig, with value as full policy path of this resource.
