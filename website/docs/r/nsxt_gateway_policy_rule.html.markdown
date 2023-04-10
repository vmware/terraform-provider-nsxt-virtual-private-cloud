---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_gateway_policy_rule"
sidebar_current: "docs-nsxt-vpc-gatewaypolicyrule"
description: |-
  Creates and manages GatewayPolicyRule.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_gateway_policy_rule

The GatewayPolicyRule resource allows the creation and management of Nsxt GatewayPolicyRule

## Example Usage
```hcl
resource "nsxt_vpc_gateway_policy_rule" "test-gatewaypolicyrule" {
    	parent_path = nsxt_vpc_parentResource.resource_name.path
	profiles = ["ANY"]
	services = ["ANY"]
	nsx_id = "test-Rule-abc"
	destination_groups = ["ANY"]
	action = "DROP"
	sequence_number = 1
	source_groups = ["ANY"]

  }
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `parent_path` - (Required) The policy path of immediate parent resource. This path will be used to create the resource.
* `resource_type` - (Optional) The type of this resource.
* `tag` - (Optional) User level field which will be printed in CLI and packet logs.
Even though there is no limitation on length of a tag, internally
tag will get truncated after 32 characters.

* `disabled` - (Optional) Flag to disable the rule. Default is enabled.
* `logged` - (Optional) Flag to enable packet logging. Default is disabled.
* `destinations_excluded` - (Optional) If set to true, the rule gets applied on all the groups that are
NOT part of the destination groups. If false, the rule applies to the
destination groups

* `sources_excluded` - (Optional) If set to true, the rule gets applied on all the groups that are
NOT part of the source groups. If false, the rule applies to the
source groups

* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `notes` - (Optional) Text for additional notes on changes.
* `action` - (Optional) The action to be applied to all the services
The JUMP_TO_APPLICATION action is only supported for rules created in the
Environment category. Once a match is hit then the rule processing
will jump to the rules present in the Application category, skipping
all further rules in the Environment category. If no rules match in
the Application category then the default application rule will be hit.
This is applicable only for DFW.

* `ip_protocol` - (Optional) Type of IP packet that should be matched while enforcing the rule.
The value is set to IPV4_IPV6 for Layer3 rule if not specified.
For Layer2/Ether rule the value must be null.

* `direction` - (Optional) Define direction of traffic.

* `sequence_number` - (Required) This field is used to resolve conflicts between multiple
Rules under Security or Gateway Policy for a Domain
If no sequence number is specified in the payload, a value of 0 is
assigned by default. If there are multiple rules with the same
sequence number then their order is not deterministic. If a specific
order of rules is desired, then one has to specify unique sequence
numbers or use the POST request on the rule entity with
a query parameter action=revise to let the framework assign a
sequence number

* `profiles` - (Optional) Holds the list of layer 7 service profile paths. These profiles accept
attributes and sub-attributes of various network services
(e.g. L4 AppId, encryption algorithm, domain name, etc) as key value
pairs. Instead of Layer 7 service profiles you can use a L7 access profile.
One of either Layer 7 service profiles or L7 Access Profile can be used in firewall rule.
In case of L7 access profile only one is allowed.

* `services` - (Optional) In order to specify all services, use the constant "ANY".
This is case insensitive. If "ANY" is used, it should
be the ONLY element in the services array. Error will be thrown
if ANY is used in conjunction with other values.

* `service_entries` - (Optional) In order to specify raw services this can be used,
along with services which contains path to services.
This can be empty or null.

  * `display_name` - (Optional) Defaults to ID if not set
  * `description` - (Optional) Description of this resource
  * `tags` - (Optional) Opaque identifiers meaningful to the API user
    * `scope` - (Optional) Tag searches may optionally be restricted by scope
    * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters
  * `resource_type` - (Required) 
* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters
* `scope` - (Optional) The list of policy paths where the rule is applied
LR/Edge/T0/T1/LRP etc. Note that a given rule can be applied
on multiple LRs/LRPs.

* `destination_groups` - (Optional) We need paths as duplicate names may exist for groups under different
domains. Along with paths we support IP Address of type IPv4 and IPv6.
IP Address can be in one of the format(CIDR, IP Address, Range of IP Address).
In order to specify all groups, use the constant "ANY". This
is case insensitive. If "ANY" is used, it should be the ONLY element
in the group array. Error will be thrown if ANY is used in conjunction
with other values.

* `source_groups` - (Optional) We need paths as duplicate names may exist for groups under different
domains. Along with paths we support IP Address of type IPv4 and IPv6.
IP Address can be in one of the format(CIDR, IP Address, Range of IP Address).
In order to specify all groups, use the constant "ANY". This
is case insensitive. If "ANY" is used, it should be the ONLY element
in the group array. Error will be thrown if ANY is used in conjunction
with other values.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of GatewayPolicyRule resource.

## Importing

An existing GatewayPolicyRule can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```hcl
terraform import nsxt_vpc_gateway_policy_rule.test-gatewaypolicyrule ID
```

The above would import NSX `GatewayPolicyRule` as a resource named test-gatewaypolicyrule with the terraform ID `ID`, 
which is the external ID of GatewayPolicyRule, with value as full policy path of this resource.
