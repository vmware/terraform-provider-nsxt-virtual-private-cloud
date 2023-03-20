---
layout: "nsxt"
page_title: "Provider: NSXT VPC"
sidebar_current: "docs-nsxt-index"
description: |-
  VMware NSX-T VPC Terraform Provider
  This provider is used to interact with the VPC resources supported by NSX-T. The provider needs to be configured with the proper credentials before it can be used.
---

<!--
    Copyright 2023 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->

# NSX-T VPC Terraform Provider

The NSX-T VPC Terraform provider gives the NSX VPC administrator a way to automate NSX's Virtual Private Cloud to
provide virtualized networking and security services.

More information on NSX can be found on the [NSX Product
Page](https://www.vmware.com/products/nsx.html)

Documentation on the NSX platform can be found on the [NSX Documentation
Page](https://docs.vmware.com/en/VMware-NSX-T/index.html)

Please use the navigation to the left to read about available data sources and
resources.

## Basic Configuration of the NSX-T VPC Terraform Provider

In order to use the NSX-T VPC Terraform provider you must first configure the
provider to communicate with the VMware NSX manager. The NSX manager is the
system which serves the NSX REST API and provides a way to configure the
desired state of the NSX system. The configuration of this provider requires
the IP address of the NSX manager.

The NSX-T VPC provider offers several ways to authenticate to the NSX manager.
Credentials can be provided statically or provided as environment variables. In
addition, client certificates can be used for authentication. For
authentication with certificates Terraform will require a certificate file and
private key file in PEM format. To use client certificates the client
certificate needs to be registered with NSX-T manager prior to invoking
Terraform.

The provider also can accept both signed and self-signed server certificates.
It is recommended that in production environments you only use certificates
signed by a certificate authority. NSX ships by default with a self-signed
server certificates as the hostname of the NSX manager is not known until the
NSX administrator determines what name or IP to use.

Setting the `allow_unverified_ssl` parameter to `true` will direct the
Terraform client to skip server certificate verification. This is not
recommended in production deployments as it is recommended that you use trusted
connection using certificates signed by a certificate authority.

With the `ca_file` parameter you can also specify a file that contains your
certificate authority certificate in PEM format to verify certificates with a
certificate authority.

There are also a number of other parameters that can be set to tune how the
provider connects to the NSX REST API. It is recommended you leave these to the
defaults unless you experience issues in which case they can be tuned to
optimize the system in your environment.

Note that with terraform 0.14 onwards, `terraform` block should be added to your
configuration:

```hcl
terraform {
  required_providers {
    nsxt = {
      source = "vmware/nsxt-vpc"
      version = "1.0.0"
    }
  }
}
```

Note that in all of the examples you will need to update attributes such as
`host`, `username`, `password`, `org`, `project`, `vpc` to match your NSX deployment configuration.

### Example of Configuration with Credentials

```hcl
provider "nsxt" {
  host                 = "192.168.110.41"
  username             = "admin"
  password             = "Admin!23Admin"
  org                  = "default"
  project              = "Dev_project"
	vpc                  = "dev_vpc"
  allow_unverified_ssl = true
}

```


### Example of Setting Environment Variables

```
export NSXT_MANAGER_HOST = "192.168.110.41"
export NSXT_USERNAME     = "admin"
export NSXT_PASSWORD     = "Admin!23Admin"
export NSXT_ORG          = "default"
export NSXT_PROJECT      = "Dev_project"
export NSXT_VPC          = "dev_vpc"
```

### Example using a Client Certificate

```hcl
provider "nsxt" {
  host                  = "192.168.110.41"
  client_auth_cert_file = "mycert.pem"
  client_auth_key_file  = "mykey.pem"
  allow_unverified_ssl  = true
  org                   = "default"
  project               = "Dev_project"
	vpc                   = "dev_vpc"
}

```

### Example with Certificate Authority Certificate

```hcl
provider "nsxt" {
  host        = "10.160.94.11"
  username    = "admin"
  password    = "qwerty"
  org         = "default"
  project     = "Dev_project"
	vpc         = "dev_vpc"
  ca_file     = "myca.pem"
}

```

## Argument Reference

The following arguments are used to configure the VMware NSX-T VPC Provider:

* `host` - (Required) The host name or IP address of the NSX-T manager. Can also
  be specified with the `NSXT_MANAGER_HOST` environment variable. Do not include
  `http://` or `https://` in the host.
* `username` - (Required) The user name to connect to the NSX-T manager as. Can
  also be specified with the `NSXT_USERNAME` environment variable.
* `password` - (Required) The password for the NSX-T manager user. Can also be
  specified with the `NSXT_PASSWORD` environment variable.
* `org` - (Required) Organisation identifier. Can also 
  be specified with the `NSXT_ORG` environment variable.
* `project` - (Required) Identifier for the project under the organisation.
  Can also be specified with the `NSXT_PROJECT` environment variable.
* `vpc` - (Required) Identifier for the VPC under the project of the organisation.
  Can also be specified with the `NSXT_VPC` environment variable.
* `connection_timeout` - (Optional) Maximum time in milliseconds for connection to
  wait for a TLS handshake. Zero means no timeout. Default: `60` 
  Can also be specified with the `NSXT_CONNECTON_TIMEOUT` environment variable.
* `max_retries` - (Optional) The maximum number of retires before failing an API
  request. Default: `4` Can also be specified with the `NSXT_MAX_RETRIES`
  environment variable.
* `client_auth_cert_file` - (Optional) The path to a certificate file for client
  certificate authorization. Can also be specified with the
  `NSXT_CLIENT_AUTH_CERT_FILE` environment variable.
* `client_auth_key_file` - (Optional) The path to a private key file for the
  certificate supplied to `client_auth_cert_file`. Can also be specified with
  the `NSXT_CLIENT_AUTH_KEY_FILE` environment variable.
* `client_auth_cert` - (Optional) Client certificate string.
  Can also be specified with the `NSXT_CLIENT_AUTH_CERT` environment variable.
* `client_auth_key` - (Optional) Client certificate private key string.
  Can also be specified with the `NSXT_CLIENT_AUTH_KEY` environment variable.
* `allow_unverified_ssl` - (Optional) Boolean that can be set to true to disable
  SSL certificate verification. This should be used with care as it could allow
  an attacker to intercept your auth token. If omitted, default value is
  `false`. Can also be specified with the `NSXT_ALLOW_UNVERIFIED_SSL`
  environment variable.
* `ca_file` - (Optional) The path to an optional CA certificate file for SSL
  validation. Can also be specified with the `NSXT_CA_FILE` environment
  variable.
* `ca` - (Optional) CA certificate string for SSL validation.
  Can also be specified with the `NSXT_CA` environment variable.

## NSX VPC Logical Networking

This release of the NSX-T VPC Terraform Provider extends to cover VPC APIs
in NSX. The Provider has resources and data sources covering VPC, VpcSubnet,
VpcSubnetPort, SecurityPolicy, SecurityPolicyRule, GatewayPolicy, GatewayPolicyRule,
VpcGroup, PolicyNAT, PolicyNATRule, VpcIpAddressAllocation, VpcSubnetIpAllocation,
VpcSubnetDhcpStaticBinding, VpcStaticRoutes, VpcSubnetBridgeProfile, Tag-operations.
This typically allows VPC admin to create and manage VPC topologies. The resources
and data sources have _vpc_ in their name. All these resources
and data sources are fully documented on the NSX-T Terraform Provider
page:•https://registry.terraform.io/providers/vmware/nsxt-vpc/latest/docs For more details on the
NSX-T Policy API usage, please look at NSX-T documentation.

### Logical Networking and Security Example Usage

The following example demonstrates using the NSX Terraform provider to create
VPC, Security policy, and security policy rule.

#### Example variables.tf File

This file allows you to define some variables that can be reused in multiple
.tf files.

```hcl
variable "nsxt_manager" {}
variable "nsxt_username" {}
variable "nsxt_password" {}
variable "nsxt_org" {}
variable "nsxt_project" {}
variable "nsxt_vpc" {}
```
#### Example terraform.tfvars File

This file allows you to set some variables that can be reused in multiple .tf
files.

```hcl
nsxt_manager  = "192.168.110.41"

nsxt_username = "admin"

nsxt_password = "default"

nsxt_org      = "default"

nsxt_project  = "project-prod"

nsxt_vpc      = "vpc-prod"
```


#### Example nsx.tf file

```hcl

################################################################################
#
# This configuration file is an example of creating a securitypolicy,
# a rule on VPC using Terraform.
#
# It creates the following objects:
#   - A Security Policy
#   - A Rule
#
# The config below requires the following to be pre-created
#   - Edge Cluster
#   - IP address block
#   - Tier-0 Gateway
#   - Org
#   - Project
#   - VPC
#
################################################################################


#
# The first step is to configure the VMware NSX provider to connect to the NSX
# REST API running on the NSX manager.
#
provider "nsxt" {
  host                  = var.nsxt_manager
  username              = var.nsxt_username
  password              = var.nsxt_password
  org                   = var.nsxt_org
  project               = var.nsxt_project
  vpc                   = var.nsxt_vpc
  allow_unverified_ssl  = true
  max_retries           = 10
}

#
# This part of the example shows some data sources we will need to refer to
# later in the .tf file. They include the tier 0 router and ip address block.
# The Tier-0 (T0) Gateway is considered a "provider" router that is pre-created
# by the NSX Admin. A T0 Gateway is used for north/south connectivity between
# the logical networking space and the physical networking space. Many Tier1
# Gateways will be connected to the T0 Gateway
#
data "nsxt_shared_project_infra_ip_address_block" "projectInfraIpAddressBlock" {
  display_name = "IP-address-block inside project"
}

data "nsxt_shared_infra_ip_address_block" "infraIpAddressBlock" {
  display_name = "IP-address-block inside infra"
}

#
# In this part of the example, we are creating the SecurityPolicy and a Rule on the VPC.
# In Rule, since we need group policy path for its properties, we have created group and referred to it in Rule.
#
  resource "nsxt_vpc_security_policy" "testSecurityPolicy" {
	nsx_id = "securitypolicy-test"
	display_name = "SecurityPolicy-test"
	description = "SecurityPolicy description"
}

resource "nsxt_vpc_security_policy_rule" "testSecurityPolicyRule" {
  nsx_id = "test-rule-abc"
	parent_path = nsxt_vpc_security_policy.testSecurityPolicy.path
  display_name = "Security policy rule - test"
	description = "Rule description"
	action = "ALLOW"
  source_groups = [nsxt_vpc_group.testGroup.path]
	destination_groups = [nsxt_vpc_group.testGroup.path]
	services = ["ANY"]
}

resource "nsxt_vpc_group" "testGroup" {
	nsx_id = "test-group-abc"
	display_name = "test-group-abc"
	description = "Group description"
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
}

```

## Feature Requests, Bug Reports, and Contributing

For more information how how to submit feature requests, bug reports, or
details on how to make your own contributions to the provider, see the [NSX-T
provider project page][nsxt-vpc-provider-project-page].

[nsxt-vpc-provider-project-page]: https://github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud
