package logic

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

// IntegrationAccountPartnersClient is the REST API for Azure Logic Apps.
type IntegrationAccountPartnersClient struct {
	BaseClient
}

// NewIntegrationAccountPartnersClient creates an instance of the IntegrationAccountPartnersClient client.
func NewIntegrationAccountPartnersClient(subscriptionID string) IntegrationAccountPartnersClient {
	return NewIntegrationAccountPartnersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationAccountPartnersClientWithBaseURI creates an instance of the IntegrationAccountPartnersClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewIntegrationAccountPartnersClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountPartnersClient {
	return IntegrationAccountPartnersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an integration account partner.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// partnerName - the integration account partner name.
// partner - the integration account partner.
func (client IntegrationAccountPartnersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, partnerName string, partner IntegrationAccountPartner) (result IntegrationAccountPartner, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountPartnersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, integrationAccountName, partnerName, partner)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IntegrationAccountPartnersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, partnerName string, partner IntegrationAccountPartner) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"partnerName":            autorest.Encode("path", partnerName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", pathParameters),
		autorest.WithJSON(partner),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountPartnersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IntegrationAccountPartnersClient) CreateOrUpdateResponder(resp *http.Response) (result IntegrationAccountPartner, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an integration account partner.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// partnerName - the integration account partner name.
func (client IntegrationAccountPartnersClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, partnerName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountPartnersClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, integrationAccountName, partnerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IntegrationAccountPartnersClient) DeletePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, partnerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"partnerName":            autorest.Encode("path", partnerName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountPartnersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IntegrationAccountPartnersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an integration account partner.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// partnerName - the integration account partner name.
func (client IntegrationAccountPartnersClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, partnerName string) (result IntegrationAccountPartner, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountPartnersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, integrationAccountName, partnerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client IntegrationAccountPartnersClient) GetPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, partnerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"partnerName":            autorest.Encode("path", partnerName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountPartnersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IntegrationAccountPartnersClient) GetResponder(resp *http.Response) (result IntegrationAccountPartner, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of integration account partners.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// top - the number of items to be included in the result.
// filter - the filter to apply on the operation.
func (client IntegrationAccountPartnersClient) List(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (result IntegrationAccountPartnerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountPartnersClient.List")
		defer func() {
			sc := -1
			if result.iaplr.Response.Response != nil {
				sc = result.iaplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, integrationAccountName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.iaplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "List", resp, "Failure sending request")
		return
	}

	result.iaplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.iaplr.hasNextLink() && result.iaplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client IntegrationAccountPartnersClient) ListPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountPartnersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IntegrationAccountPartnersClient) ListResponder(resp *http.Response) (result IntegrationAccountPartnerListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IntegrationAccountPartnersClient) listNextResults(ctx context.Context, lastResults IntegrationAccountPartnerListResult) (result IntegrationAccountPartnerListResult, err error) {
	req, err := lastResults.integrationAccountPartnerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountPartnersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IntegrationAccountPartnersClient) ListComplete(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (result IntegrationAccountPartnerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountPartnersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, integrationAccountName, top, filter)
	return
}
