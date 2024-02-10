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

// FluxConfigurationsClient is the kubernetesConfiguration Client
type FluxConfigurationsClient struct {
	BaseClient
}

// NewFluxConfigurationsClient creates an instance of the FluxConfigurationsClient client.
func NewFluxConfigurationsClient(subscriptionID string) FluxConfigurationsClient {
	return NewFluxConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFluxConfigurationsClientWithBaseURI creates an instance of the FluxConfigurationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewFluxConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) FluxConfigurationsClient {
	return FluxConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create a new Kubernetes Flux Configuration.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterRp - the Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes,
// Microsoft.HybridContainerService.
// clusterResourceName - the Kubernetes cluster resource name - i.e. managedClusters, connectedClusters,
// provisionedClusters.
// clusterName - the name of the kubernetes cluster.
// fluxConfigurationName - name of the Flux Configuration.
// fluxConfiguration - properties necessary to Create a FluxConfiguration.
func (client FluxConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfiguration FluxConfiguration) (result FluxConfigurationsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FluxConfigurationsClient.CreateOrUpdate")
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
		return result, validation.NewError("kubernetesconfiguration.FluxConfigurationsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName, fluxConfiguration)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client FluxConfigurationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfiguration FluxConfiguration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":           autorest.Encode("path", clusterName),
		"clusterResourceName":   autorest.Encode("path", clusterResourceName),
		"clusterRp":             autorest.Encode("path", clusterRp),
		"fluxConfigurationName": autorest.Encode("path", fluxConfigurationName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/{fluxConfigurationName}", pathParameters),
		autorest.WithJSON(fluxConfiguration),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client FluxConfigurationsClient) CreateOrUpdateSender(req *http.Request) (future FluxConfigurationsCreateOrUpdateFuture, err error) {
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
func (client FluxConfigurationsClient) CreateOrUpdateResponder(resp *http.Response) (result FluxConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete this will delete the YAML file used to set up the Flux Configuration, thus stopping future sync from the
// source repo.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterRp - the Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes,
// Microsoft.HybridContainerService.
// clusterResourceName - the Kubernetes cluster resource name - i.e. managedClusters, connectedClusters,
// provisionedClusters.
// clusterName - the name of the kubernetes cluster.
// fluxConfigurationName - name of the Flux Configuration.
// forceDelete - delete the extension resource in Azure - not the normal asynchronous delete.
func (client FluxConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, forceDelete *bool) (result FluxConfigurationsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FluxConfigurationsClient.Delete")
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
		return result, validation.NewError("kubernetesconfiguration.FluxConfigurationsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName, forceDelete)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FluxConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, forceDelete *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":           autorest.Encode("path", clusterName),
		"clusterResourceName":   autorest.Encode("path", clusterResourceName),
		"clusterRp":             autorest.Encode("path", clusterRp),
		"fluxConfigurationName": autorest.Encode("path", fluxConfigurationName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/{fluxConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FluxConfigurationsClient) DeleteSender(req *http.Request) (future FluxConfigurationsDeleteFuture, err error) {
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
func (client FluxConfigurationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets details of the Flux Configuration.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterRp - the Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes,
// Microsoft.HybridContainerService.
// clusterResourceName - the Kubernetes cluster resource name - i.e. managedClusters, connectedClusters,
// provisionedClusters.
// clusterName - the name of the kubernetes cluster.
// fluxConfigurationName - name of the Flux Configuration.
func (client FluxConfigurationsClient) Get(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string) (result FluxConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FluxConfigurationsClient.Get")
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
		return result, validation.NewError("kubernetesconfiguration.FluxConfigurationsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client FluxConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":           autorest.Encode("path", clusterName),
		"clusterResourceName":   autorest.Encode("path", clusterResourceName),
		"clusterRp":             autorest.Encode("path", clusterRp),
		"fluxConfigurationName": autorest.Encode("path", fluxConfigurationName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/{fluxConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FluxConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FluxConfigurationsClient) GetResponder(resp *http.Response) (result FluxConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all Flux Configurations.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterRp - the Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes,
// Microsoft.HybridContainerService.
// clusterResourceName - the Kubernetes cluster resource name - i.e. managedClusters, connectedClusters,
// provisionedClusters.
// clusterName - the name of the kubernetes cluster.
func (client FluxConfigurationsClient) List(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (result FluxConfigurationsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FluxConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.fcl.Response.Response != nil {
				sc = result.fcl.Response.Response.StatusCode
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
		return result, validation.NewError("kubernetesconfiguration.FluxConfigurationsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.fcl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result.fcl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.fcl.hasNextLink() && result.fcl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client FluxConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (*http.Request, error) {
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client FluxConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FluxConfigurationsClient) ListResponder(resp *http.Response) (result FluxConfigurationsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client FluxConfigurationsClient) listNextResults(ctx context.Context, lastResults FluxConfigurationsList) (result FluxConfigurationsList, err error) {
	req, err := lastResults.fluxConfigurationsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client FluxConfigurationsClient) ListComplete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (result FluxConfigurationsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FluxConfigurationsClient.List")
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

// Update update an existing Kubernetes Flux Configuration.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterRp - the Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes,
// Microsoft.HybridContainerService.
// clusterResourceName - the Kubernetes cluster resource name - i.e. managedClusters, connectedClusters,
// provisionedClusters.
// clusterName - the name of the kubernetes cluster.
// fluxConfigurationName - name of the Flux Configuration.
// fluxConfigurationPatch - properties to Patch in an existing Flux Configuration.
func (client FluxConfigurationsClient) Update(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfigurationPatch FluxConfigurationPatch) (result FluxConfigurationsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FluxConfigurationsClient.Update")
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
		return result, validation.NewError("kubernetesconfiguration.FluxConfigurationsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName, fluxConfigurationPatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigurationsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client FluxConfigurationsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, fluxConfigurationPatch FluxConfigurationPatch) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":           autorest.Encode("path", clusterName),
		"clusterResourceName":   autorest.Encode("path", clusterResourceName),
		"clusterRp":             autorest.Encode("path", clusterRp),
		"fluxConfigurationName": autorest.Encode("path", fluxConfigurationName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/{fluxConfigurationName}", pathParameters),
		autorest.WithJSON(fluxConfigurationPatch),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client FluxConfigurationsClient) UpdateSender(req *http.Request) (future FluxConfigurationsUpdateFuture, err error) {
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
func (client FluxConfigurationsClient) UpdateResponder(resp *http.Response) (result FluxConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}