//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHostGroup_CreateOrUpdate_WithUltraSSD.json
func ExampleDedicatedHostGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewDedicatedHostGroupsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"myDedicatedHostGroup",
		armcompute.DedicatedHostGroup{
			Location: to.Ptr("westus"),
			Tags: map[string]*string{
				"department": to.Ptr("finance"),
			},
			Properties: &armcompute.DedicatedHostGroupProperties{
				AdditionalCapabilities: &armcompute.DedicatedHostGroupPropertiesAdditionalCapabilities{
					UltraSSDEnabled: to.Ptr(true),
				},
				PlatformFaultDomainCount:  to.Ptr[int32](3),
				SupportAutomaticPlacement: to.Ptr(true),
			},
			Zones: []*string{
				to.Ptr("1")},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHostGroups_Update_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewDedicatedHostGroupsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"rgcompute",
		"aaaa",
		armcompute.DedicatedHostGroupUpdate{
			Tags: map[string]*string{
				"key9921": to.Ptr("aaaaaaaaaa"),
			},
			Properties: &armcompute.DedicatedHostGroupProperties{
				InstanceView: &armcompute.DedicatedHostGroupInstanceView{
					Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
						{
							AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
								AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
									{
										Count:  to.Ptr[float64](26),
										VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
									}},
							},
							Statuses: []*armcompute.InstanceViewStatus{
								{
									Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
									DisplayStatus: to.Ptr("aaaaaa"),
									Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
									Message:       to.Ptr("a"),
									Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
								}},
						}},
				},
				PlatformFaultDomainCount:  to.Ptr[int32](3),
				SupportAutomaticPlacement: to.Ptr(true),
			},
			Zones: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHostGroups_Delete_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewDedicatedHostGroupsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"rgcompute",
		"a",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHostGroup_Get.json
func ExampleDedicatedHostGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewDedicatedHostGroupsClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"myResourceGroup",
		"myDedicatedHostGroup",
		&armcompute.DedicatedHostGroupsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHostGroups_ListByResourceGroup_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewDedicatedHostGroupsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("rgcompute",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHostGroups_ListBySubscription_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewDedicatedHostGroupsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySubscriptionPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}