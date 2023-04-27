/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

package clients

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	nsxtsession "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/session"
)

type Configuration struct {
	BasePath             string
	Scheme               string
	UserName             string
	Password             string
	OrgID                string
	ProjectID            string
	VpcID                string
	RemoteAuth           bool
	UserAgent            string
	ClientAuthCertFile   string
	ClientAuthKeyFile    string
	CAFile               string
	ClientAuthCertString string
	ClientAuthKeyString  string
	CAString             string
	Insecure             bool
	SkipSessionAuth      bool
}

// NsxtClient -- an API Client for Nsxt manager
type NsxtClient struct {
	Config      *Configuration
	NsxtSession *nsxtsession.NsxtSession
}

// NewNsxtClient initiates an NsxtSession and returns an NsxtClient wrapping that session
func NewNsxtClient(host string, username string, orgID string, projectID string, vpcID string, clientAuthCertFile string,
	clientAuthCertStr string, clientAuthKeyFile string, clientAuthKeyStr string, caFile string, caStr string, securityContextNeeded bool, options ...func(*nsxtsession.NsxtSession) error) (*NsxtClient, error) {
	nsxtClient := NsxtClient{}
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	if !securityContextNeeded {
		// Set client certificate if securityContextNeeded is false
		// Set TLS config with certificate in session's transport object
		if len(clientAuthCertFile) > 0 {
			cert, err := tls.LoadX509KeyPair(clientAuthCertFile, clientAuthKeyFile)
			if err != nil {
				return nil, err
			}
			tlsConfig.GetClientCertificate = func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
				return &cert, nil
			}
		}

		if len(clientAuthCertStr) > 0 {
			cert, err := tls.X509KeyPair([]byte(clientAuthCertStr),
				[]byte(clientAuthKeyStr))
			if err != nil {
				return nil, err
			}
			tlsConfig.GetClientCertificate = func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
				return &cert, nil
			}
		}

		if len(caFile) > 0 {
			caCert, err := ioutil.ReadFile(caFile)
			if err != nil {
				return nil, err
			}
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)
			tlsConfig.RootCAs = caCertPool
		}

		if len(caStr) > 0 {
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM([]byte(caStr))
			tlsConfig.RootCAs = caCertPool
		}
	}

	nsxtSession, err := nsxtsession.NewNsxtSession(host, username, tlsConfig, securityContextNeeded, options...)
	if err != nil {
		return &nsxtClient, err
	}
	nsxtClient.NsxtSession = nsxtSession
	// set org and project ID config in client
	config, err1 := NewConfiguration(orgID, projectID, vpcID)
	if err1 != nil {
		return &nsxtClient, err1
	}
	nsxtClient.Config = config
	return &nsxtClient, nil
}

func NewConfiguration(orgID string, projectID string, vpcID string) (*Configuration, error) {
	cfg := &Configuration{
		OrgID:     orgID,
		ProjectID: projectID,
		VpcID:     vpcID,
		BasePath:  "policy/api/v1",
	}
	return cfg, nil
}
