---
layout: "nsxt"
page_title: "Nsxt: nsxt_vpc_security_policy"
sidebar_current: "docs-nsxt-vpc-securitypolicy"
description: |-
  Creates and manages SecurityPolicy.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

<!-- Auto generated code. DO NOT EDIT. -->

# nsxt_vpc_security_policy

The SecurityPolicy resource allows the creation and management of Nsxt SecurityPolicy

## Example Usage
```hcl
resource "nsxt_vpc_security_policy" "test-securitypolicy" {
    	nsx_id = "test-SecurityPolicy-abc"
	display_name = "Test Policy"
	description = "This is Test vpc policy patch operation"
	sequence_number = 1

  }
```

## Argument Reference

The following arguments are supported:

* `nsx_id` - (Required) The NSX ID of this resource. This ID will be used to create the resource.
* `category` - (Optional) - Distributed Firewall -
Policy framework provides five pre-defined categories for classifying
a security policy. They are "Ethernet","Emergency", "Infrastructure"
"Environment" and "Application". There is a pre-determined order in
which the policy framework manages the priority of these security
policies. Ethernet category is for supporting layer 2 firewall rules.
The other four categories are applicable for layer 3 rules. Amongst
them, the Emergency category has the highest priority followed by
Infrastructure, Environment and then Application rules. Administrator
can choose to categorize a security policy into the above categories
or can choose to leave it empty. If empty it will have the least
precedence w.r.t the above four categories.
- Edge Firewall -
Policy Framework for Edge Firewall provides six pre-defined categories
"Emergency", "SystemRules", "SharedPreRules", "LocalGatewayRules",
"AutoServiceRules" and "Default", in order of priority of rules.
All categories are allowed for Gatetway Policies that belong
to 'default' Domain. However, for user created domains, category is
restricted to "SharedPreRules" or "LocalGatewayRules" only. Also, the
users can add/modify/delete rules from only the "SharedPreRules" and
"LocalGatewayRules" categories. If user doesn't specify the category
then defaulted to "Rules". System generated category is used by NSX
created rules, for example BFD rules. Autoplumbed category used by
NSX verticals to autoplumb data path rules. Finally, "Default" category
is the placeholder default rules with lowest in the order of priority.

* `comments` - (Optional) Comments for security policy lock/unlock.
* `tcp_strict` - (Optional) Ensures that a 3 way TCP handshake is done before the data packets
are sent.
tcp_strict=true is supported only for stateful security policies.
If the tcp_strict flag is not specified and the security policy
is stateful, then tcp_strict will be set to true.

* `scheduler_path` - (Optional) Provides a mechanism to apply the rules in this policy for a specified
time duration.

* `stateful` - (Optional) Stateful or Stateless nature of security policy is enforced on all
rules in this security policy. When it is stateful, the state of
the network connects are tracked and a stateful packet inspection is
performed.
Layer3 security policies can be stateful or stateless. By default, they are stateful.
Layer2 security policies can only be stateless.

* `resource_type` - (Optional) The type of this resource.
* `locked` - (Optional) Indicates whether a security policy should be locked. If the
security policy is locked by a user, then no other user would
be able to modify this security policy. Once the user releases
the lock, other users can update this security policy.

* `display_name` - (Optional) Defaults to ID if not set
* `description` - (Optional) Description of this resource
* `connectivity_preference` - (Optional) This field indicates the default connectivity policy for the security
policy. Based on the connectivitiy preference, a default rule for this
security policy will be created. An appropriate action will be set on
the rule based on the value of the connectivity preference. If NONE is
selected or no connectivity preference is specified, then no default
rule for the security policy gets created. The default rule that gets
created will be a any-any rule and applied to entities specified in the
scope of the security policy. Specifying the connectivity_preference
without specifying the scope is not allowed. The scope has to be a
Group and one cannot specify IPAddress directly in the group that is
used as scope. This default rule is only applicable for the Layer3
security policies.
ALLOWLIST - Adds a default drop rule. Administrator can then use "allow"
rules to allow traffic between groups
DENYLIST - Adds a default allow rule. Admin can then use "drop" rules
to block traffic between groups
ALLOWLIST_ENABLE_LOGGING - Allowlisting with logging enabled
DENYLIST_ENABLE_LOGGING - Denylisting with logging enabled
NONE - No default rule is created.

* `sequence_number` - (Required) This field is used to resolve conflicts between security policies
across domains. In order to change the sequence number of a policy
one can fire a POST request on the policy entity with
a query parameter action=revise
The sequence number field will reflect the value of the computed
sequence number upon execution of the above mentioned POST request.
For scenarios where the administrator is using a template to update
several security policies, the only way to set the sequence number is
to explicitly specify the sequence number for each security policy.
If no sequence number is specified in the payload, a value of 0 is
assigned by default. If there are multiple policies with the same
sequence number then their order is not deterministic. If a specific
order of policies is desired, then one has to specify unique sequence
numbers or use the POST request on the policy entity with
a query parameter action=revise to let the framework assign a
sequence number.
The value of sequence number must be between 0 and 999,999.

* `tags` - (Optional) Opaque identifiers meaningful to the API user
  * `scope` - (Optional) Tag searches may optionally be restricted by scope
  * `tag` - (Optional) Identifier meaningful to user with maximum length of 256 characters
* `scope` - (Optional) The list of group paths where the rules in this policy will get
applied. This scope will take precedence over rule level scope.
Supported only for security and redirection policies. In case of
RedirectionPolicy, it is expected only when the policy is NS and
redirecting to service chain.

* `application_connectivity_strategy` - (Optional) This field indicates the application connectivity policy for the security
policy.

  * `logging_enabled` - (Optional) Flag to enable packet logging. Default is disabled.
  * `application_connectivity_strategy` - (Required) App connectivity strategies


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Unique identifier of this resource
* `_revision` -  Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - Indicates the NSX policy path of SecurityPolicy resource.

## Importing

An existing SecurityPolicy can be [imported][docs-import] into this resource, via the following command:

[docs-import]: https://www.terraform.io/cli/import

```hcl
terraform import nsxt_vpc_security_policy.test-securitypolicy ID
```

The above would import NSX `SecurityPolicy` as a resource named test-securitypolicy with the terraform ID `ID`, 
which is the external ID of SecurityPolicy, with value as full policy path of this resource.
