---
layout: "nsxt"
page_title: "NSXT: nsxt_vpc_policy_nat_rule"
sidebar_current: "docs-nsxt-vpc-policynatrule"
description: |-
  Creates and manages PolicyNatRule.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_policy_nat_rule

The PolicyNatRule resource allows the creation and management of Nsxt PolicyNatRule

## Example Usage
```hcl
resource "nsxt_vpc_policy_nat_rule" "test-policynatrule" {
    	translated_network = "192.168.1.1"
	display_name = "MyNATRuleDemo"
	description = "Example of a NAT rule"
	service = "/infra/services/AD_Server"
	enabled = true
	parent_path = "/orgs/default/projects/Dev_project/vpcs/dev_vpc/nat/USER"
	nsx_id = "test-PolicyNatRule-abc"
	translated_ports = "80-82"
	action = "DNAT"
	logging = false
	firewall_match = "MATCH_EXTERNAL_ADDRESS"
	destination_network = "10.117.5.19"
	sequence_number = 10

  }
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `parent_path` - (Required) The policy path of NAT under VPC. Under VPC there will be 2 different NATs(sections).
(USER and NAT64) 
* `service` - (Optional) It represents the path of Service on which the NAT rule will be applied.
If not provided or if it is blank then Policy manager will consider it
as ANY.
Please note, if this is a DNAT, the destination_port of the service will
be realized on NSX Manager as the translated_port. And if this is a SNAT,
the destination_port will be ignored.

* `resource_type` - (Optional) The type of this resource.
* `logging` - (Optional) The flag, which suggests whether the logging of NAT rule is enabled or
disabled. The default is False.

* `enabled` - (Optional) The flag, which suggests whether the NAT rule is enabled or
disabled. The default is True.

* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `policy_based_vpn_mode` - (Optional) It indicates how the NSX edge applies Nat Policy for VPN traffic. It is supported only for
Nat Rule action type DNAT and NO_DNAT. For all other NAT action, leave it unassigned.
BYPASS - Default vpn mode. It indicates that Nat policy will be applied to the inbound traffic
         on Routed Based VPN tunnel, if the policy based VTI is in the "scope" for this rule.
         Default value will be set to BYPASS if
MATCH - It indicates that this NAT rule will only match the Policy Based VPN traffic.

* `translated_ports` - (Optional) Please note, if there is service configured in this NAT rule, the translated_port
will be realized on NSX Manager as the destination_port. If there is no sevice configured,
the port will be ignored.

* `translated_network` - (Optional) This supports single IP address or comma separated list of single IP
addresses or CIDR.
If user specify the CIDR, this value is actually used as an IP pool that includes both the subnet and
broadcast addresses as valid for NAT translations.
This does not support IP range or IP sets.
For SNAT, DNAT, NAT64 and REFLEXIVE rules, this ia a mandatory field, which
represents the translated network address.
For NO_SNAT and NO_DNAT this should be empty.
For VPC SNAT and Refelexive NATRule, translated network address should be
IPv4 address allocated from External Block associated with VPC.

* `destination_network` - (Optional) This supports single IP address or comma separated list of single IP
addresses or CIDR. This does not support IP range or IP sets.
For DNAT and NO_DNAT rules, this is a mandatory field, and represents
the destination network for the incoming packets.
For other type of rules, optionally it can contain destination network
of outgoing packets.
NULL value for this field represents ANY network.
For VPC DNAT NATRule, destination network address should be IPv4 address
allocated from External Block associated with VPC.

* `source_network` - (Optional) This supports single IP address or comma separated list of single IP
addresses or CIDR. This does not support IP range or IP sets.
For SNAT, NO_SNAT, NAT64 and REFLEXIVE rules, this is a mandatory field and
represents the source network of the packets leaving the network.
For DNAT and NO_DNAT rules, optionally it can contain source network
of incoming packets.
NULL value for this field represents ANY network.

* `sequence_number` - (Required) The sequence_number decides the rule_priority of a NAT rule.
Sequence_number and rule_priority have 1:1 mapping.For each NAT section,
there will be reserved rule_priority numbers.The valid range of
rule_priority number is from 0 to 2147483647(MAX_INT).
1. INTERNAL section
    rule_priority reserved from 0 - 1023 (1024 rules)
    valid sequence_number range  0 - 1023
2. USER section
   rule_priority reserved from 1024 - 2147482623 (2147481600 rules)
   valid sequence_number range  0 - 2147481599
3. DEFAULT section
   rule_priority reserved from 2147482624 - 2147483647 (1024 rules)
   valid sequence_number range  0 - 1023

* `firewall_match` - (Optional) It indicates how the firewall matches the address after NATing if firewall
stage is not skipped.

MATCH_EXTERNAL_ADDRESS indicates the firewall will be applied to external address
of a NAT rule. For SNAT, the external address is the translated source address
after NAT is done. For DNAT, the external address is the original destination
address before NAT is done. For REFLEXIVE, to egress traffic, the firewall
will be applied to the translated source address after NAT is done; To ingress
traffic, the firewall will be applied to the original destination address
before NAT is done.

MATCH_INTERNAL_ADDRESS indicates the firewall will be applied to internal
address of a NAT rule. For SNAT, the internal address is the original source
address before NAT is done. For DNAT, the internal address is the translated
destination address after NAT is done. For REFLEXIVE, to egress traffic, the
firewall will be applied to the original source address before NAT is done;
To ingress traffic, the firewall will be applied to the translated destination
address after NAT is done.

BYPASS indicates the firewall stage will be skipped.

For NO_SNAT or NO_DNAT, it must be BYPASS or leave it unassigned

* `scope` - (Optional) Represents the array of policy paths of ProviderInterface or NetworkInterface or
labels of type ProviderInterface or NetworkInterface or IPSecVpnSession on which the NAT rule should
get enforced.
The interfaces must belong to the same router for which the NAT Rule is created.

* `action` - (Required) Source NAT(SNAT) - translates a source IP address in an outbound packet so that
the packet appears to originate from a different network. SNAT is only supported
when the logical router is running in active-standby mode.
Destination NAT(DNAT) - translates the destination IP address of inbound packets
so that packets are delivered to a target address into another network. DNAT is
only supported when the logical router is running in active-standby mode.
Reflexive NAT(REFLEXIVE) - IP-Range and CIDR are supported to define the "n".
The number of original networks should be exactly the same as that of
translated networks. The address translation is deterministic. Reflexive is
supported on both Active/Standby and Active/Active LR.
NO_SNAT and NO_DNAT - These do not have support for translated_fields, only
source_network and destination_network fields are supported.
NAT64 - translates an external IPv6 address to a internal IPv4 address.

* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of PolicyNatRule resource.

## Importing

An existing PolicyNatRule can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```hcl
terraform import nsxt_vpc_policy_nat_rule.test-policynatrule ID
```

The above would import NSX `PolicyNatRule` as a resource named test-policynatrule with the terraform ID `ID`, 
which is the external ID of PolicyNatRule, with value as full policy path of this resource.
