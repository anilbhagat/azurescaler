package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// TenantsClient is the all resource groups and resources exist within subscriptions. These operation enable you get
// information about your subscriptions and tenants. A tenant is a dedicated instance of Azure Active Directory (Azure
// AD) for your organization.
type TenantsClient struct {
	BaseClient
}

// NewTenantsClient creates an instance of the TenantsClient client.
func NewTenantsClient() TenantsClient {
	return NewTenantsClientWithBaseURI(DefaultBaseURI)
}

// NewTenantsClientWithBaseURI creates an instance of the TenantsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTenantsClientWithBaseURI(baseURI string) TenantsClient {
	return TenantsClient{NewWithBaseURI(baseURI)}
}

// List gets the tenants for your account.
func (client TenantsClient) List(ctx context.Context) (result TenantListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantsClient.List")
		defer func() {
			sc := -1
			if result.tlr.Response.Response != nil {
				sc = result.tlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.TenantsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.tlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscriptions.TenantsClient", "List", resp, "Failure sending request")
		return
	}

	result.tlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.TenantsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.tlr.hasNextLink() && result.tlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client TenantsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/tenants"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TenantsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TenantsClient) ListResponder(resp *http.Response) (result TenantListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client TenantsClient) listNextResults(ctx context.Context, lastResults TenantListResult) (result TenantListResult, err error) {
	req, err := lastResults.tenantListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "subscriptions.TenantsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "subscriptions.TenantsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.TenantsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client TenantsClient) ListComplete(ctx context.Context) (result TenantListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
