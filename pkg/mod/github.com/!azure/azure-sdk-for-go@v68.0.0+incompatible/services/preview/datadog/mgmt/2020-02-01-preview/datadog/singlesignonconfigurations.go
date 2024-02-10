package datadog

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

// SingleSignOnConfigurationsClient is the client for the SingleSignOnConfigurations methods of the Datadog service.
type SingleSignOnConfigurationsClient struct {
	BaseClient
}

// NewSingleSignOnConfigurationsClient creates an instance of the SingleSignOnConfigurationsClient client.
func NewSingleSignOnConfigurationsClient(subscriptionID string) SingleSignOnConfigurationsClient {
	return NewSingleSignOnConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSingleSignOnConfigurationsClientWithBaseURI creates an instance of the SingleSignOnConfigurationsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewSingleSignOnConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) SingleSignOnConfigurationsClient {
	return SingleSignOnConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate sends the create or update request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Datadog resource belongs.
// monitorName - monitor resource name
// configurationName - configuration name
func (client SingleSignOnConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, body *SingleSignOnResource) (result SingleSignOnConfigurationsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SingleSignOnConfigurationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, monitorName, configurationName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.SingleSignOnConfigurationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.SingleSignOnConfigurationsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SingleSignOnConfigurationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, body *SingleSignOnResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", configurationName),
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Name = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/singleSignOnConfigurations/{configurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SingleSignOnConfigurationsClient) CreateOrUpdateSender(req *http.Request) (future SingleSignOnConfigurationsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SingleSignOnConfigurationsClient) CreateOrUpdateResponder(resp *http.Response) (result SingleSignOnResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Datadog resource belongs.
// monitorName - monitor resource name
// configurationName - configuration name
func (client SingleSignOnConfigurationsClient) Get(ctx context.Context, resourceGroupName string, monitorName string, configurationName string) (result SingleSignOnResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SingleSignOnConfigurationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, monitorName, configurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.SingleSignOnConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datadog.SingleSignOnConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.SingleSignOnConfigurationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SingleSignOnConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, monitorName string, configurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", configurationName),
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/singleSignOnConfigurations/{configurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SingleSignOnConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SingleSignOnConfigurationsClient) GetResponder(resp *http.Response) (result SingleSignOnResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Datadog resource belongs.
// monitorName - monitor resource name
func (client SingleSignOnConfigurationsClient) List(ctx context.Context, resourceGroupName string, monitorName string) (result SingleSignOnResourceListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SingleSignOnConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.ssorlr.Response.Response != nil {
				sc = result.ssorlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, monitorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.SingleSignOnConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ssorlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datadog.SingleSignOnConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result.ssorlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.SingleSignOnConfigurationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.ssorlr.hasNextLink() && result.ssorlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SingleSignOnConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, monitorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/singleSignOnConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SingleSignOnConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SingleSignOnConfigurationsClient) ListResponder(resp *http.Response) (result SingleSignOnResourceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SingleSignOnConfigurationsClient) listNextResults(ctx context.Context, lastResults SingleSignOnResourceListResponse) (result SingleSignOnResourceListResponse, err error) {
	req, err := lastResults.singleSignOnResourceListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datadog.SingleSignOnConfigurationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datadog.SingleSignOnConfigurationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.SingleSignOnConfigurationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SingleSignOnConfigurationsClient) ListComplete(ctx context.Context, resourceGroupName string, monitorName string) (result SingleSignOnResourceListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SingleSignOnConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, monitorName)
	return
}