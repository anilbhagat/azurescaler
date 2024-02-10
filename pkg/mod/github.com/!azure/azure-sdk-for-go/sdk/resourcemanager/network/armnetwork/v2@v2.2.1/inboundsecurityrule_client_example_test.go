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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/InboundSecurityRulePut.json
func ExampleInboundSecurityRuleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInboundSecurityRuleClient().BeginCreateOrUpdate(ctx, "rg1", "nva", "rule1", armnetwork.InboundSecurityRule{
		Properties: &armnetwork.InboundSecurityRuleProperties{
			Rules: []*armnetwork.InboundSecurityRules{
				{
					DestinationPortRange: to.Ptr[int32](22),
					SourceAddressPrefix:  to.Ptr("50.20.121.5/32"),
					Protocol:             to.Ptr(armnetwork.InboundSecurityRulesProtocolTCP),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InboundSecurityRule = armnetwork.InboundSecurityRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva/InboundSecurityRules/rule1"),
	// 	Name: to.Ptr("rule1"),
	// 	Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
	// 	Properties: &armnetwork.InboundSecurityRuleProperties{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		Rules: []*armnetwork.InboundSecurityRules{
	// 			{
	// 				DestinationPortRange: to.Ptr[int32](22),
	// 				SourceAddressPrefix: to.Ptr("50.20.121.5/32"),
	// 				Protocol: to.Ptr(armnetwork.InboundSecurityRulesProtocolTCP),
	// 		}},
	// 	},
	// }
}