// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package subscriptionapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/Azure/go-autorest/autorest"
)

// SubscriptionsClientAPI contains the set of methods on the SubscriptionsClient type.
type SubscriptionsClientAPI interface {
	Get(ctx context.Context, subscriptionID string) (result subscription.Model, err error)
	List(ctx context.Context) (result subscription.ListResultPage, err error)
	ListComplete(ctx context.Context) (result subscription.ListResultIterator, err error)
	ListLocations(ctx context.Context, subscriptionID string) (result subscription.LocationListResult, err error)
}

var _ SubscriptionsClientAPI = (*subscription.SubscriptionsClient)(nil)

// TenantsClientAPI contains the set of methods on the TenantsClient type.
type TenantsClientAPI interface {
	List(ctx context.Context) (result subscription.TenantListResultPage, err error)
	ListComplete(ctx context.Context) (result subscription.TenantListResultIterator, err error)
}

var _ TenantsClientAPI = (*subscription.TenantsClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Cancel(ctx context.Context, subscriptionID string) (result subscription.CanceledSubscriptionID, err error)
	Enable(ctx context.Context, subscriptionID string) (result subscription.EnabledSubscriptionID, err error)
	Rename(ctx context.Context, subscriptionID string, body subscription.Name) (result subscription.RenamedSubscriptionID, err error)
}

var _ ClientAPI = (*subscription.Client)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result subscription.OperationListResult, err error)
}

var _ OperationsClientAPI = (*subscription.OperationsClient)(nil)

// AliasClientAPI contains the set of methods on the AliasClient type.
type AliasClientAPI interface {
	Create(ctx context.Context, aliasName string, body subscription.PutAliasRequest) (result subscription.AliasCreateFuture, err error)
	Delete(ctx context.Context, aliasName string) (result autorest.Response, err error)
	Get(ctx context.Context, aliasName string) (result subscription.PutAliasResponse, err error)
	List(ctx context.Context) (result subscription.PutAliasListResult, err error)
}

var _ AliasClientAPI = (*subscription.AliasClient)(nil)
