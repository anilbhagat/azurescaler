// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package accountapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/datalake/store/mgmt/2015-10-01-preview/account"
	"github.com/Azure/go-autorest/autorest"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, name string, parameters account.DataLakeStoreAccount) (result account.CreateFuture, err error)
	CreateOrUpdateFirewallRule(ctx context.Context, resourceGroupName string, accountName string, name string, parameters account.FirewallRule) (result account.FirewallRule, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result account.DeleteFuture, err error)
	DeleteFirewallRule(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string) (result autorest.Response, err error)
	EnableKeyVault(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result account.DataLakeStoreAccount, err error)
	GetFirewallRule(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string) (result account.FirewallRule, err error)
	List(ctx context.Context, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeStoreAccountListResultPage, err error)
	ListComplete(ctx context.Context, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeStoreAccountListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeStoreAccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeStoreAccountListResultIterator, err error)
	ListFirewallRules(ctx context.Context, resourceGroupName string, accountName string) (result account.DataLakeStoreFirewallRuleListResultPage, err error)
	ListFirewallRulesComplete(ctx context.Context, resourceGroupName string, accountName string) (result account.DataLakeStoreFirewallRuleListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, name string, parameters account.DataLakeStoreAccount) (result account.UpdateFuture, err error)
}

var _ ClientAPI = (*account.Client)(nil)