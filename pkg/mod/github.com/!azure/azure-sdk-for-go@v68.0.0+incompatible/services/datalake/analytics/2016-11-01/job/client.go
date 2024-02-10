// Package job implements the Azure ARM Job service API version 2016-11-01.
//
// Creates an Azure Data Lake Analytics job client.
package job

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultAdlaJobDNSSuffix is the default value for adla job dns suffix
	DefaultAdlaJobDNSSuffix = "azuredatalakeanalytics.net"
)

// BaseClient is the base client for Job.
type BaseClient struct {
	autorest.Client
	AdlaJobDNSSuffix string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithoutDefaults(DefaultAdlaJobDNSSuffix)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(adlaJobDNSSuffix string) BaseClient {
	return BaseClient{
		Client:           autorest.NewClientWithUserAgent(UserAgent()),
		AdlaJobDNSSuffix: adlaJobDNSSuffix,
	}
}