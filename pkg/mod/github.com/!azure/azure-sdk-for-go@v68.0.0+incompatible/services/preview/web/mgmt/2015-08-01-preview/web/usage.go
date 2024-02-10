package web

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

// UsageClient is the webSite Management Client
type UsageClient struct {
	BaseClient
}

// NewUsageClient creates an instance of the UsageClient client.
func NewUsageClient(subscriptionID string) UsageClient {
	return NewUsageClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsageClientWithBaseURI creates an instance of the UsageClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUsageClientWithBaseURI(baseURI string, subscriptionID string) UsageClient {
	return UsageClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetUsage sends the get usage request.
// Parameters:
// resourceGroupName - name of resource group
// environmentName - environment name
// lastID - last marker that was returned from the batch
// batchSize - size of the batch to be returned.
func (client UsageClient) GetUsage(ctx context.Context, resourceGroupName string, environmentName string, lastID string, batchSize int32) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageClient.GetUsage")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetUsagePreparer(ctx, resourceGroupName, environmentName, lastID, batchSize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.UsageClient", "GetUsage", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.UsageClient", "GetUsage", resp, "Failure sending request")
		return
	}

	result, err = client.GetUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.UsageClient", "GetUsage", resp, "Failure responding to request")
		return
	}

	return
}

// GetUsagePreparer prepares the GetUsage request.
func (client UsageClient) GetUsagePreparer(ctx context.Context, resourceGroupName string, environmentName string, lastID string, batchSize int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":   autorest.Encode("path", environmentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"batchSize":   autorest.Encode("query", batchSize),
		"lastId":      autorest.Encode("query", lastID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web.Admin/environments/{environmentName}/usage", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUsageSender sends the GetUsage request. The method will close the
// http.Response Body if it receives an error.
func (client UsageClient) GetUsageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetUsageResponder handles the response to the GetUsage request. The method always
// closes the http.Response Body.
func (client UsageClient) GetUsageResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}