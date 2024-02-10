package compute

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

// CommunityGalleriesClient is the compute Client
type CommunityGalleriesClient struct {
	BaseClient
}

// NewCommunityGalleriesClient creates an instance of the CommunityGalleriesClient client.
func NewCommunityGalleriesClient(subscriptionID string) CommunityGalleriesClient {
	return NewCommunityGalleriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCommunityGalleriesClientWithBaseURI creates an instance of the CommunityGalleriesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewCommunityGalleriesClientWithBaseURI(baseURI string, subscriptionID string) CommunityGalleriesClient {
	return CommunityGalleriesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a community gallery by gallery public name.
// Parameters:
// location - resource location.
// publicGalleryName - the public name of the community gallery.
func (client CommunityGalleriesClient) Get(ctx context.Context, location string, publicGalleryName string) (result CommunityGallery, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunityGalleriesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, location, publicGalleryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleriesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleriesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CommunityGalleriesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CommunityGalleriesClient) GetPreparer(ctx context.Context, location string, publicGalleryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":          autorest.Encode("path", location),
		"publicGalleryName": autorest.Encode("path", publicGalleryName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-03"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/communityGalleries/{publicGalleryName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CommunityGalleriesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CommunityGalleriesClient) GetResponder(resp *http.Response) (result CommunityGallery, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
