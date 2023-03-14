# terraform-provider-for-vmware-nsxt-virtual-private-cloud

This is the repository for the Terraform NSX-T VPC Provider, which one can use with
Terraform to work with [VMware NSX-T][vmware-nsxt].

[vmware-nsxt]: https://www.vmware.com/products/nsx.html

For general information about Terraform, visit the [official
website][tf-website] and the [GitHub project page][tf-github].

[tf-website]: https://terraform.io/
[tf-github]: https://github.com/hashicorp/terraform

Documentation on the NSX platform can be found at the [NSX-T Documentation page](https://docs.vmware.com/en/VMware-NSX-T/index.html)

# Using the Provider

The latest version of this provider requires Terraform v0.12 or higher to run.

Note that you need to run `terraform init` to fetch the provider before
deploying.

## Full Provider Documentation

The provider is documented in full on the Terraform website and can be found
[here](https://registry.terraform.io/providers/vmware/nsxt-vpc/latest). Check the provider documentation for details on entering your connection information and how to get started with writing configuration for policy resources.

### Controlling the provider version

Note that you can also control the provider version. This requires the use of a `provider` block in your Terraform configuration if you have not added one already.

The syntax is as follows:

```hcl
provider "nsxt" {
  version = "~> 1.0.0"
  ...
}
```

[Read more][provider-vc] on provider version control.

[provider-vc]: https://www.terraform.io/docs/configuration/providers.html#provider-versions

# Automated Installation (Recommended)

Download and initialization of Terraform providers is with the “terraform init” command. This applies to the NSX-T provider as well. Once the provider block for the NSX-T provider is specified in your .tf file, “terraform init” will detect a need for the provider and download it to your environment.
You can list versions of providers installed in your environment by running “terraform version” command:

```hcl
$ ./terraform version
Terraform v1.2.1
on linux_amd64
+ provider registry.terraform.io/vmware/nsxt-vpc v1.0.0
```

# Manual Installation

**NOTE:** Unless you are [developing](#developing-the-provider) or require a
pre-release bugfix or feature, you will want to use the officially released
version of the provider (see [the section above](#using-the-provider)).

**NOTE:** Recommended way to compile the provider is using [Go Modules](https://blog.golang.org/using-go-modules).

**NOTE:** For terraform 0.13, please refer to [provider installation configuration][install-013] in order to use custom provider.

[install-013]: https://www.terraform.io/docs/commands/cli-config.html#provider-installation

## Cloning the Project

First, you will want to clone the repository to
`$GOPATH/src/github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud`:

```sh
mkdir -p $GOPATH/src/github.com/vmware
cd $GOPATH/src/github.com/vmware
git clone https://github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud.git
```

## Building and Installing the Provider

Recommended golang version is go1.18 onwards.
After the clone has been completed, you can enter the provider directory and build the provider.

```sh
cd $GOPATH/src/github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud
make
```

After the build is complete, copy the provider executable `terraform-provider-for-vmware-nsxt-virtual-private-cloud` into location specified in your provider installation configuration. Make sure to delete provider lock files that might exist in your working directory due to prior provider usage. Run `terraform init`.
For developing, consider using [dev overrides configuration][dev-overrides]. Please note that `terraform init` should not be used with dev overrides.

[dev-overrides]: https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers

# Developing the Provider

If you wish to work on the provider, you'll first need [Go][go-website]
installed on your machine (version 1.14+ is recommended). You'll also need to
correctly setup a [GOPATH][gopath], as well as adding `$GOPATH/bin` to your
`$PATH`.

[go-website]: https://golang.org/
[gopath]: http://golang.org/doc/code.html#GOPATH

See [Manual Installation](#manual-installation) for details on building the
provider.

# Testing the Provider

**NOTE:** Testing the NSX-T provider is currently a complex operation as it
requires having a NSX-T manager endpoint to test against, which should be
hosting a standard configuration for a NSX-T cluster.

## Configuring Environment Variables

Most of the tests in this provider require a comprehensive list of environment
variables to run. See the individual `*_test.go` files in the [`nsxt/`](nsxt/)
directory for more details on the tunables that can be
used to specify the locations of certain pre-created resources that the tests
require.

Minimum environment variable :
```sh
$ export NSXT_MANAGER_HOST="10.92.104.56"
$ export NSXT_USERNAME="admin"
$ export NSXT_PASSWORD="u2.OpqZc0Ptd"
$ export NSXT_ALLOW_UNVERIFIED_SSL=true
```

## Running the Acceptance Tests

After this is done, you can run the acceptance tests by running:

```sh
$ make testacc
```

If you want to run against a specific set of tests, run `make testacc` with the
`TESTARGS` parameter containing the run mask as per below:

```sh
make testacc TESTARGS="-run=TestNSXTStaticRoutesBasic"
```

This following example would run all of the acceptance tests matching
`TestNSXTStaticRoutesBasic`. Change this for the specific tests you want
to run.

# Interoperability

The following versions of NSX are supported:

 * NSX-T 1.0.0

# Support

The NSX policy Terraform provider is now VMware supported as well as community supported. For bugs and feature requests please open a Github Issue and label it appropriately or contact VMware support.

# License

Copyright © 2022-2023 VMware, Inc. All Rights Reserved.

The NSX VPC Terraform provider is available under [MPL2.0 license](https://github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/blob/main/LICENSE).