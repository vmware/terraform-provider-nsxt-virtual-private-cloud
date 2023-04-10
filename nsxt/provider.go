/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

// Auto generated code. DO NOT EDIT.

//nolint
package nsxt

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
	nsxtsession "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/session"
)

var defaultRetryOnStatusCodes = []int{400, 409, 429, 500, 503, 504, 603}

// Provider for VMWare NSX-T
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"allow_unverified_ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSXT_ALLOW_UNVERIFIED_SSL", false),
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSXT_MANAGER_HOST", nil),
				Description: "The hostname or IP address of the NSX manager.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSXT_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSXT_PASSWORD", nil),
				Sensitive:   true,
			},
			"org": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organisation identifier",
				DefaultFunc: schema.EnvDefaultFunc("NSXT_ORG", "default"),
			},
			"project": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identifier for the project under the organisation",
				DefaultFunc: schema.EnvDefaultFunc("NSXT_PROJECT", nil),
			},
			"vpc": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identifier for the VPC under the project of an organisation",
				DefaultFunc: schema.EnvDefaultFunc("NSXT_VPC", nil),
			},
			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum number of HTTP client retries",
				DefaultFunc: schema.EnvDefaultFunc("NSXT_MAX_RETRIES", 3),
			},
			"retry_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Delay interval in milliseconds between retries of a request",
				DefaultFunc: schema.EnvDefaultFunc("NSXT_RETRY_INTERVAL", 500),
			},
			"retry_on_status_codes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "HTTP replies status codes to retry on",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"connection_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum time in milliseconds for connection to wait for a TLS handshake. Zero means no timeout",
				DefaultFunc: schema.EnvDefaultFunc("NSXT_CONNECTON_TIMEOUT", 60),
			},
			"enforcement_point": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enforcement Point for NSXT Policy",
				DefaultFunc: schema.EnvDefaultFunc("NSXT_POLICY_ENFORCEMENT_POINT", "default"),
			},
			"remote_auth": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSXT_REMOTE_AUTH", false),
			},
			"client_auth_cert": {
				Type:        schema.TypeString,
				Description: "Client certificate passed as string",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSXT_CLIENT_AUTH_CERT", nil),
			},
			"client_auth_cert_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSXT_CLIENT_AUTH_CERT_FILE", nil),
			},
			"client_auth_key": {
				Type:        schema.TypeString,
				Description: "Client certificate key passed as string",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSXT_CLIENT_AUTH_KEY", nil),
			},
			"client_auth_key_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSXT_CLIENT_AUTH_KEY_FILE", nil),
			},
			"ca": {
				Type:        schema.TypeString,
				Description: "CA certificate passed as string",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSXT_CA", nil),
			},
			"ca_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSXT_CA_FILE", nil),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"nsxt_shared_project_infra_ip_address_block":       dataSourceNsxtSharedProjectInfraIpAddressBlock(),
			"nsxt_shared_infra_ip_address_block":               dataSourceNsxtSharedInfraIpAddressBlock(),
			"nsxt_vpc_gateway_policy":                          dataSourceNsxtVpcGatewayPolicy(),
			"nsxt_vpc_static_routes":                           dataSourceNsxtVpcStaticRoutes(),
			"nsxt_vpc_policy_nat_rule":                         dataSourceNsxtVpcPolicyNatRule(),
			"nsxt_vpc_security_policy_rule":                    dataSourceNsxtVpcSecurityPolicyRule(),
			"nsxt_shared_infra_l2_bridge_endpoint_profile":     dataSourceNsxtSharedInfraL2BridgeEndpointProfile(),
			"nsxt_vpc_security_policy":                         dataSourceNsxtVpcSecurityPolicy(),
			"nsxt_shared_project_infra_service":                dataSourceNsxtSharedProjectInfraService(),
			"nsxt_shared_project_infra_ip_address_pool":        dataSourceNsxtSharedProjectInfraIpAddressPool(),
			"nsxt_vpc_subnet":                                  dataSourceNsxtVpcSubnet(),
			"nsxt_vpc_subnet_ip_address_allocation":            dataSourceNsxtVpcSubnetIpAddressAllocation(),
			"nsxt_shared_project_infra_group":                  dataSourceNsxtSharedProjectInfraGroup(),
			"nsxt_vpc_gateway_policy_rule":                     dataSourceNsxtVpcGatewayPolicyRule(),
			"nsxt_shared_project_infra_policy_context_profile": dataSourceNsxtSharedProjectInfraPolicyContextProfile(),
			"nsxt_vpc_subnet_port":                             dataSourceNsxtVpcSubnetPort(),
			"nsxt_shared_infra_service":                        dataSourceNsxtSharedInfraService(),
			"nsxt_shared_infra_group":                          dataSourceNsxtSharedInfraGroup(),
			"nsxt_vpc_group":                                   dataSourceNsxtVpcGroup(),
			"nsxt_shared_infra_policy_context_profile":         dataSourceNsxtSharedInfraPolicyContextProfile(),
			"nsxt_vpc_ip_address_allocation":                   dataSourceNsxtVpcIpAddressAllocation(),
			"nsxt_shared_infra_ip_address_pool":                dataSourceNsxtSharedInfraIpAddressPool(),
			"nsxt_vpc_dhcp_v4_static_binding_config":           dataSourceNsxtVpcDhcpV4StaticBindingConfig(),
			"nsxt_vpc_dhcp_v6_static_binding_config":           dataSourceNsxtVpcDhcpV6StaticBindingConfig(),
			"nsxt_vpc_vm":                                      dataSourceNsxtVpcVM(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"nsxt_vpc_gateway_policy":                resourceNsxtVpcGatewayPolicy(),
			"nsxt_vpc_static_routes":                 resourceNsxtVpcStaticRoutes(),
			"nsxt_vpc_policy_nat_rule":               resourceNsxtVpcPolicyNatRule(),
			"nsxt_vpc_security_policy_rule":          resourceNsxtVpcSecurityPolicyRule(),
			"nsxt_vpc_security_policy":               resourceNsxtVpcSecurityPolicy(),
			"nsxt_vpc_subnet":                        resourceNsxtVpcSubnet(),
			"nsxt_vpc_subnet_ip_address_allocation":  resourceNsxtVpcSubnetIpAddressAllocation(),
			"nsxt_vpc_gateway_policy_rule":           resourceNsxtVpcGatewayPolicyRule(),
			"nsxt_vpc_subnet_port":                   resourceNsxtVpcSubnetPort(),
			"nsxt_vpc_group":                         resourceNsxtVpcGroup(),
			"nsxt_vpc_ip_address_allocation":         resourceNsxtVpcIpAddressAllocation(),
			"nsxt_vpc_dhcp_v4_static_binding_config": resourceNsxtVpcDhcpV4StaticBindingConfig(),
			"nsxt_vpc_dhcp_v6_static_binding_config": resourceNsxtVpcDhcpV6StaticBindingConfig(),
			"nsxt_vpc_vm_tags":                       resourceNsxtVpcVmTags(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func isNSXTManagerAccessible(nsxtClient interface{}) error {
	// Do GET on /api/v1/node API and check for response status code whether server is accessible or not
	var response interface{}
	uri := "api/v1/node"
	err := nsxtClient.(*nsxtclient.NsxtClient).NsxtSession.Get(uri, &response)
	if err != nil {
		return fmt.Errorf("failed to access NSX-T manager (%s). Please check connectivity and authentication settings of the provider", err)
	}
	return nil
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	log.Printf("start of providerConfigure")
	config := Configuration{
		NsxManagerHost:       d.Get("host").(string),
		Username:             d.Get("username").(string),
		Password:             d.Get("password").(string),
		OrgId:                d.Get("org").(string),
		ProjectId:            d.Get("project").(string),
		VpcId:                d.Get("vpc").(string),
		Insecure:             d.Get("allow_unverified_ssl").(bool),
		MaxRetries:           d.Get("max_retries").(int),
		RetryInterval:        d.Get("retry_interval").(int),
		RetryStausCodes:      d.Get("retry_on_status_codes").([]interface{}),
		ConnectionTimeOut:    d.Get("connection_timeout").(int),
		EnforcementPoint:     d.Get("enforcement_point").(string),
		RemoteAuth:           d.Get("remote_auth").(bool),
		ClientAuthCertFile:   d.Get("client_auth_cert_file").(string),
		ClientAuthCertString: d.Get("client_auth_cert").(string),
		ClientAuthKeyFile:    d.Get("client_auth_key_file").(string),
		ClientAuthKeyString:  d.Get("client_auth_key").(string),
		CAFile:               d.Get("ca_file").(string),
		CAString:             d.Get("ca").(string),
	}
	var nsxtClient interface{}
	var err error

	securityContextNeeded := true
	clientAuthDefined := (len(config.ClientAuthCertFile) > 0) || (len(config.ClientAuthCertString) > 0)
	if clientAuthDefined && !config.RemoteAuth {
		securityContextNeeded = false
	}

	retryStatuses := make([]int, 0, len(config.RetryStausCodes))
	for _, s := range config.RetryStausCodes {
		retryStatuses = append(retryStatuses, s.(int))
	}

	if len(retryStatuses) == 0 {
		// Set to the defaults if empty
		retryStatuses = append(retryStatuses, defaultRetryOnStatusCodes...)
	}

	nsxtClient, err = nsxtclient.NewNsxtClient(config.NsxManagerHost, config.Username, config.OrgId, config.ProjectId,
		config.VpcId,
		config.ClientAuthCertFile, config.ClientAuthCertString,
		config.ClientAuthKeyFile, config.ClientAuthKeyString,
		config.CAFile, config.CAString,
		securityContextNeeded,
		nsxtsession.SetPassword(config.Password), nsxtsession.SetInsecure(config.Insecure),
		nsxtsession.SetEnforcementPoint(config.EnforcementPoint),
		nsxtsession.SetMaxAPIRetries(config.MaxRetries),
		nsxtsession.SetAPIRetryInterval(config.RetryInterval),
		nsxtsession.SetRetryStatusCodes(retryStatuses),
		nsxtsession.SetTimeout(time.Duration(config.ConnectionTimeOut*int(time.Second))))
	log.Printf("Nsxt client and session created")

	err = isNSXTManagerAccessible(nsxtClient)
	return nsxtClient, err
}

type Configuration struct {
	NsxManagerHost       string
	Username             string
	Password             string
	OrgId                string
	ProjectId            string
	VpcId                string
	Insecure             bool
	MaxRetries           int
	RetryInterval        int
	RetryStausCodes      []interface{}
	ConnectionTimeOut    int
	EnforcementPoint     string
	ClientAuthCertFile   string
	ClientAuthCertString string
	ClientAuthKeyFile    string
	ClientAuthKeyString  string
	CAFile               string
	CAString             string
	RemoteAuth           bool
}
