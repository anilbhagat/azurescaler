package kubernetesconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ExtensionsClient is the kubernetesConfiguration Client
type ExtensionsClient struct {
	BaseClient
}

// NewExtensionsClient creates an instance of the ExtensionsClient client.
func NewExtensionsClient(subscriptionID string) ExtensionsClient {
	return NewExtensionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExtensionsClientWithBaseURI creates an instance of the ExtensionsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExtensionsClientWithBaseURI(baseURI string, subscriptionID string) ExtensionsClient {
	return ExtensionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a new Kubernetes Cluster Extension.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterRp - the Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes,
// Microsoft.HybridContainerService.
// clusterResourceName - the Kubernetes cluster resource name - i.e. managedClusters, connectedClusters,
// provisionedClusters.
// clusterName - the name of the kubernetes cluster.
// extensionName - name of the Extension.
// extension - properties necessary to Create an Extension.
func (client ExtensionsClient) Create(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, extension Extension) (result ExtensionsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kubernetesconfiguration.ExtensionsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName, extension)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ExtensionsClient) CreatePreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, extension Extension) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":         autorest.Encode("path", clusterName),
		"clusterResourceName": autorest.Encode("path", clusterResourceName),
		"clusterRp":           autorest.Encode("path", clusterRp),
		"extensionName":       autorest.Encode("path", extensionName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions/{extensionName}", pathParameters),
		autorest.WithJSON(extension),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionsClient) CreateSender(req *http.Request) (future ExtensionsCreateFuture, err error) {
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ExtensionsClient) CreateResponder(resp *http.Response) (result Extension, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a Kubernetes Cluster Extension. This will cause the Agent to Uninstall the extension from the cluster.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterRp - the Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes,
// Microsoft.HybridContainerService.
// clusterResourceName - the Kubernetes cluster resource name - i.e. managedClusters, connectedClusters,
// provisionedClusters.
// clusterName - the name of the kubernetes cluster.
// extensionName - name of the Extension.
// forceDelete - delete the extension resource in Azure - not the normal asynchronous delete.
func (client ExtensionsClient) Delete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, forceDelete *bool) (result ExtensionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kubernetesconfiguration.ExtensionsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName, forceDelete)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExtensionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, forceDelete *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":         autorest.Encode("path", clusterName),
		"clusterResourceName": autorest.Encode("path", clusterResourceName),
		"clusterRp":           autorest.Encode("path", clusterRp),
		"extensionName":       autorest.Encode("path", extensionName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if forceDelete != nil {
		queryParameters["forceDelete"] = autorest.Encode("query", *forceDelete)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions/{extensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionsClient) DeleteSender(req *http.Request) (future ExtensionsDeleteFuture, err error) {
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

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExtensionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets Kubernetes Cluster Extension.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterRp - the Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes,
// Microsoft.HybridContainerService.
// clusterResourceName - the Kubernetes cluster resource name - i.e. managedClusters, connectedClusters,
// provisionedClusters.
// clusterName - the name of the kubernetes cluster.
// extensionName - name of the Extension.
func (client ExtensionsClient) Get(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string) (result Extension, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kubernetesconfiguration.ExtensionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExtensionsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":         autorest.Encode("path", clusterName),
		"clusterResourceName": autorest.Encode("path", clusterResourceName),
		"clusterRp":           autorest.Encode("path", clusterRp),
		"extensionName":       autorest.Encode("path", extensionName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions/{extensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExtensionsClient) GetResponder(resp *http.Response) (result Extension, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all Extensions in the cluster.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterRp - the Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes,
// Microsoft.HybridContainerService.
// clusterResourceName - the Kubernetes cluster resource name - i.e. managedClusters, connectedClusters,
// provisionedClusters.
// clusterName - the name of the kubernetes cluster.
func (client ExtensionsClient) List(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (result ExtensionsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.List")
		defer func() {
			sc := -1
			if result.el.Response.Response != nil {
				sc = result.el.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kubernetesconfiguration.ExtensionsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.el.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "List", resp, "Failure sending request")
		return
	}

	result.el, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.el.hasNextLink() && result.el.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ExtensionsClient) ListPreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":         autorest.Encode("path", clusterName),
		"clusterResourceName": autorest.Encode("path", clusterResourceName),
		"clusterRp":           autorest.Encode("path", clusterRp),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExtensionsClient) ListResponder(resp *http.Response) (result ExtensionsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ExtensionsClient) listNextResults(ctx context.Context, lastResults ExtensionsList) (result ExtensionsList, err error) {
	req, err := lastResults.extensionsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExtensionsClient) ListComplete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (result ExtensionsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName)
	return
}

// Update patch an existing Kubernetes Cluster Extension.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterRp - the Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes,
// Microsoft.HybridContainerService.
// clusterResourceName - the Kubernetes cluster resource name - i.e. managedClusters, connectedClusters,
// provisionedClusters.
// clusterName - the name of the kubernetes cluster.
// extensionName - name of the Extension.
// patchExtension - properties to Patch in an existing Extension.
func (client ExtensionsClient) Update(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, patchExtension PatchExtension) (result ExtensionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtensionsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kubernetesconfiguration.ExtensionsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName, patchExtension)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.ExtensionsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ExtensionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, patchExtension PatchExtension) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":         autorest.Encode("path", clusterName),
		"clusterResourceName": autorest.Encode("path", clusterResourceName),
		"clusterRp":           autorest.Encode("path", clusterRp),
		"extensionName":       autorest.Encode("path", extensionName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions/{extensionName}", pathParameters),
		autorest.WithJSON(patchExtension),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ExtensionsClient) UpdateSender(req *http.Request) (future ExtensionsUpdateFuture, err error) {
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

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ExtensionsClient) UpdateResponder(resp *http.Response) (result Extension, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
