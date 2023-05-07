---
layout: "nsxt"
page_title: "NSXT: nsxt_policy_vpc_nat_rule"
sidebar_current: "docs-nsxt-policyvpcnatrule"
description: |-
  Creates and manages PolicyVpcNatRule.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_policy_vpc_nat_rule

The PolicyVpcNatRule resource allows the creation and management of Nsxt PolicyVpcNatRule

## Example Usage
```hcl
resource "nsxt_policy_vpc_nat_rule" "test-policyvpcnatrule" {
    	translated_network = "192.168.1.1"
	display_name = "MyNATRuleDemo"
	description = "Example of a NAT rule"
	enabled = true
	parent_path = "/orgs/default/projects/Dev_project/vpcs/dev_vpc/nat/USER"
	nsx_id = "test-PolicyVpcNatRule-abc"
	action = "DNAT"
	logging = false
	destination_network = "10.117.5.19"
	sequence_number = 10

  }
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `parent_path` - (Required) The policy path of NAT under VPC. Under VPC there will be 2 different NATs(sections).
(USER and NAT64) 
* `resource_type` - (Optional) The type of this resource.
* `logging` - (Optional) The flag, which suggests whether the logging of NAT rule is enabled or
disabled. The default is False.

* `enabled` - (Optional) The flag, which suggests whether the NAT rule is enabled or
disabled. The default is True.

* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `destination_network` - (Optional) This supports single IP address and it does not support IP range or IP sets.
For DNAT rules, this is a mandatory field, and represents
the destination network for the incoming packets.
For other type of rules, optionally it can contain destination network
of outgoing packets.
NULL value for this field represents ANY network.
In case of DNAT NATRule, destination network address should be IPv4 address
allocated from External Block associated with VPC.

* `translated_network` - (Optional) This supports single IP address or comma separated list of single IP
addresses or CIDR.
If user specify the CIDR, this value is actually used as an IP pool that includes both the subnet and
broadcast addresses as valid for NAT translations.
This does not support IP range or IP sets.
For SNAT, DNAT and REFLEXIVE rules, this ia a mandatory field, which
represents the translated network address.
In case of SNAT and Refelexive NATRule, translated network address should be single
IPv4 address allocated from External Block associated with VPC.

* `source_network` - (Optional) This supports single IP address or comma separated list of single IP
addresses or CIDR. This does not support IP range or IP sets.
For SNAT and REFLEXIVE rules, this is a mandatory field and
represents the source network of the packets leaving the network.
For DNAT rules, optionally it can contain source network
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

* `action` - (Required) Source NAT(SNAT) - translates a source IP address into an outbound packet so that
the packet appears to originate from a different network.
Destination NAT(DNAT) - translates the destination IP address of inbound packets
so that packets are delivered to a target address into another network.
Reflexive NAT(REFLEXIVE) - one-to-one mapping of source and destination IP addresses.

* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of PolicyVpcNatRule resource.

## Importing

An existing PolicyVpcNatRule can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```hcl
terraform import nsxt_policy_vpc_nat_rule.test-policyvpcnatrule ID
```

The above would import NSX `PolicyVpcNatRule` as a resource named test-policyvpcnatrule with the terraform ID `ID`, 
which is the external ID of PolicyVpcNatRule, with value as full policy path of this resource.
