/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/vmware/terraform-provider-nsxt-virtual-private-cloud/nsxt"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return nsxt.Provider()
		},
	})
}
