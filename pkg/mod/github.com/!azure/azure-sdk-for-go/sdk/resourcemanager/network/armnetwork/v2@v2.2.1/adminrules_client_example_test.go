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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/NetworkManagerAdminRuleList.json
func ExampleAdminRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAdminRulesClient().NewListPager("rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", &armnetwork.AdminRulesClientListOptions{Top: nil,
		SkipToken: nil,
	})
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
		// page.AdminRuleListResult = armnetwork.AdminRuleListResult{
		// 	Value: []armnetwork.BaseAdminRuleClassification{
		// 		&armnetwork.AdminRule{
		// 			Name: to.Ptr("SampleAdminRule"),
		// 			Type: to.Ptr("Microsoft.Network/networkManagers/securityAdminConfigurations/ruleCollections/rules"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkmanagers/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleAdminRule"),
		// 			Kind: to.Ptr(armnetwork.AdminRuleKindCustom),
		// 			SystemData: &armnetwork.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
		// 				CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 				CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 				LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 			},
		// 			Properties: &armnetwork.AdminPropertiesFormat{
		// 				Description: to.Ptr("This is Sample Admin Rule"),
		// 				Access: to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
		// 				DestinationPortRanges: []*string{
		// 					to.Ptr("22")},
		// 					Destinations: []*armnetwork.AddressPrefixItem{
		// 						{
		// 							AddressPrefix: to.Ptr("*"),
		// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
		// 					}},
		// 					Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
		// 					Priority: to.Ptr[int32](1),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					SourcePortRanges: []*string{
		// 						to.Ptr("0-65535")},
		// 						Sources: []*armnetwork.AddressPrefixItem{
		// 							{
		// 								AddressPrefix: to.Ptr("*"),
		// 								AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
		// 						}},
		// 						Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/NetworkManagerAdminRuleGet.json
func ExampleAdminRulesClient_Get_getsSecurityAdminRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdminRulesClient().Get(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleAdminRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.AdminRulesClientGetResponse{
	// 	                            BaseAdminRuleClassification: &armnetwork.AdminRule{
	// 		Name: to.Ptr("SampleAdminRule"),
	// 		Type: to.Ptr("Microsoft.Network/networkManagers/securityAdminConfigurations/ruleCollections/rules"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleAdminRule"),
	// 		Kind: to.Ptr(armnetwork.AdminRuleKindCustom),
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 		Properties: &armnetwork.AdminPropertiesFormat{
	// 			Description: to.Ptr("This is Sample Admin Rule"),
	// 			Access: to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
	// 			DestinationPortRanges: []*string{
	// 				to.Ptr("22")},
	// 				Destinations: []*armnetwork.AddressPrefixItem{
	// 					{
	// 						AddressPrefix: to.Ptr("*"),
	// 						AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 				}},
	// 				Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 				Priority: to.Ptr[int32](1),
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				SourcePortRanges: []*string{
	// 					to.Ptr("0-65535")},
	// 					Sources: []*armnetwork.AddressPrefixItem{
	// 						{
	// 							AddressPrefix: to.Ptr("*"),
	// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 					}},
	// 					Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 				},
	// 			},
	// 			                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/NetworkManagerDefaultAdminRuleGet.json
func ExampleAdminRulesClient_Get_getsSecurityDefaultAdminRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdminRulesClient().Get(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleDefaultAdminRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.AdminRulesClientGetResponse{
	// 	                            BaseAdminRuleClassification: &armnetwork.DefaultAdminRule{
	// 		Name: to.Ptr("SampleDefaultAdminRule"),
	// 		Type: to.Ptr("Microsoft.Network/networkManagers/securityAdminConfigurations/ruleCollections/rules"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleDefaultAdminRule"),
	// 		Kind: to.Ptr(armnetwork.AdminRuleKindDefault),
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 		Properties: &armnetwork.DefaultAdminPropertiesFormat{
	// 			Description: to.Ptr("This is Sample Default Admin Rule"),
	// 			Access: to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
	// 			DestinationPortRanges: []*string{
	// 				to.Ptr("22")},
	// 				Destinations: []*armnetwork.AddressPrefixItem{
	// 					{
	// 						AddressPrefix: to.Ptr("*"),
	// 						AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 				}},
	// 				Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 				Flag: to.Ptr("AllowVnetInbound"),
	// 				Priority: to.Ptr[int32](1),
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				SourcePortRanges: []*string{
	// 					to.Ptr("0-65535")},
	// 					Sources: []*armnetwork.AddressPrefixItem{
	// 						{
	// 							AddressPrefix: to.Ptr("*"),
	// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 					}},
	// 					Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 				},
	// 			},
	// 			                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/NetworkManagerDefaultAdminRulePut.json
func ExampleAdminRulesClient_CreateOrUpdate_createADefaultAdminRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdminRulesClient().CreateOrUpdate(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleDefaultAdminRule", &armnetwork.DefaultAdminRule{
		Kind: to.Ptr(armnetwork.AdminRuleKindDefault),
		Properties: &armnetwork.DefaultAdminPropertiesFormat{
			Flag: to.Ptr("AllowVnetInbound"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.AdminRulesClientCreateOrUpdateResponse{
	// 	                            BaseAdminRuleClassification: &armnetwork.DefaultAdminRule{
	// 		Name: to.Ptr("SampleDefaultAdminRule"),
	// 		Type: to.Ptr("Microsoft.Network/networkManagers/securityAdminConfigurations/ruleCollections/rules"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleDefaultAdminRule"),
	// 		Kind: to.Ptr(armnetwork.AdminRuleKindDefault),
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 		Properties: &armnetwork.DefaultAdminPropertiesFormat{
	// 			Description: to.Ptr("This is Sample Default Admin Rule"),
	// 			Access: to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
	// 			DestinationPortRanges: []*string{
	// 				to.Ptr("22")},
	// 				Destinations: []*armnetwork.AddressPrefixItem{
	// 					{
	// 						AddressPrefix: to.Ptr("*"),
	// 						AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 				}},
	// 				Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 				Flag: to.Ptr("AllowVnetInbound"),
	// 				Priority: to.Ptr[int32](1),
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				SourcePortRanges: []*string{
	// 					to.Ptr("0-65535")},
	// 					Sources: []*armnetwork.AddressPrefixItem{
	// 						{
	// 							AddressPrefix: to.Ptr("Internet"),
	// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeServiceTag),
	// 					}},
	// 					Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 				},
	// 			},
	// 			                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/NetworkManagerAdminRulePut.json
func ExampleAdminRulesClient_CreateOrUpdate_createAnAdminRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdminRulesClient().CreateOrUpdate(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleAdminRule", &armnetwork.AdminRule{
		Kind: to.Ptr(armnetwork.AdminRuleKindCustom),
		Properties: &armnetwork.AdminPropertiesFormat{
			Description: to.Ptr("This is Sample Admin Rule"),
			Access:      to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
			DestinationPortRanges: []*string{
				to.Ptr("22")},
			Destinations: []*armnetwork.AddressPrefixItem{
				{
					AddressPrefix:     to.Ptr("*"),
					AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
				}},
			Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
			Priority:  to.Ptr[int32](1),
			SourcePortRanges: []*string{
				to.Ptr("0-65535")},
			Sources: []*armnetwork.AddressPrefixItem{
				{
					AddressPrefix:     to.Ptr("Internet"),
					AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeServiceTag),
				}},
			Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.AdminRulesClientCreateOrUpdateResponse{
	// 	                            BaseAdminRuleClassification: &armnetwork.AdminRule{
	// 		Name: to.Ptr("SampleAdminRule"),
	// 		Type: to.Ptr("Microsoft.Network/networkManagers/securityAdminConfigurations/ruleCollections/rules"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/myTestSecurityConfig/ruleCollections/testRuleCollection/rules/SampleAdminRule"),
	// 		Kind: to.Ptr(armnetwork.AdminRuleKindCustom),
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 		Properties: &armnetwork.AdminPropertiesFormat{
	// 			Description: to.Ptr("This is Sample Admin Rule"),
	// 			Access: to.Ptr(armnetwork.SecurityConfigurationRuleAccessDeny),
	// 			DestinationPortRanges: []*string{
	// 				to.Ptr("22")},
	// 				Destinations: []*armnetwork.AddressPrefixItem{
	// 					{
	// 						AddressPrefix: to.Ptr("*"),
	// 						AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeIPPrefix),
	// 				}},
	// 				Direction: to.Ptr(armnetwork.SecurityConfigurationRuleDirectionInbound),
	// 				Priority: to.Ptr[int32](1),
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				SourcePortRanges: []*string{
	// 					to.Ptr("0-65535")},
	// 					Sources: []*armnetwork.AddressPrefixItem{
	// 						{
	// 							AddressPrefix: to.Ptr("Internet"),
	// 							AddressPrefixType: to.Ptr(armnetwork.AddressPrefixTypeServiceTag),
	// 					}},
	// 					Protocol: to.Ptr(armnetwork.SecurityConfigurationRuleProtocolTCP),
	// 				},
	// 			},
	// 			                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/NetworkManagerAdminRuleDelete.json
func ExampleAdminRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAdminRulesClient().BeginDelete(ctx, "rg1", "testNetworkManager", "myTestSecurityConfig", "testRuleCollection", "SampleAdminRule", &armnetwork.AdminRulesClientBeginDeleteOptions{Force: to.Ptr(false)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}