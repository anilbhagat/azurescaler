//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/EndpointServicesList.json
func ExampleAvailableEndpointServicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailableEndpointServicesClient().NewListPager("westus", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.EndpointServicesListResult = armnetwork.EndpointServicesListResult{
		// 	Value: []*armnetwork.EndpointServiceResult{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/virtualNetworkEndpointServices/Microsoft.Storage"),
		// 			Name: to.Ptr("Microsoft.Storage"),
		// 			Type: to.Ptr("Microsoft.Network/virtualNetworkEndpointServices"),
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/virtualNetworkEndpointServices/Microsoft.Sql"),
		// 			Name: to.Ptr("Microsoft.Sql"),
		// 			Type: to.Ptr("Microsoft.Network/virtualNetworkEndpointServices"),
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/virtualNetworkEndpointServices/Microsoft.AzureActiveDirectory"),
		// 			Name: to.Ptr("Microsoft.AzureActiveDirectory"),
		// 			Type: to.Ptr("Microsoft.Network/virtualNetworkEndpointServices"),
		// 	}},
		// }
	}
}