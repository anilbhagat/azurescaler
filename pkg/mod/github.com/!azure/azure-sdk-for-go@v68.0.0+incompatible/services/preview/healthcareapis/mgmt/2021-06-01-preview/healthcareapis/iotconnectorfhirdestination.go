package healthcareapis

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

// IotConnectorFhirDestinationClient is the azure Healthcare APIs Client
type IotConnectorFhirDestinationClient struct {
	BaseClient
}

// NewIotConnectorFhirDestinationClient creates an instance of the IotConnectorFhirDestinationClient client.
func NewIotConnectorFhirDestinationClient(subscriptionID string) IotConnectorFhirDestinationClient {
	return NewIotConnectorFhirDestinationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIotConnectorFhirDestinationClientWithBaseURI creates an instance of the IotConnectorFhirDestinationClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewIotConnectorFhirDestinationClientWithBaseURI(baseURI string, subscriptionID string) IotConnectorFhirDestinationClient {
	return IotConnectorFhirDestinationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an IoT Connector FHIR destination resource with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// workspaceName - the name of workspace resource.
// iotConnectorName - the name of IoT Connector resource.
// fhirDestinationName - the name of IoT Connector FHIR destination resource.
// iotFhirDestination - the parameters for creating or updating an IoT Connector FHIR destination resource.
func (client IotConnectorFhirDestinationClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, fhirDestinationName string, iotFhirDestination IotFhirDestination) (result IotConnectorFhirDestinationCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotConnectorFhirDestinationClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: iotConnectorName,
			Constraints: []validation.Constraint{{Target: "iotConnectorName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "iotConnectorName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: fhirDestinationName,
			Constraints: []validation.Constraint{{Target: "fhirDestinationName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "fhirDestinationName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: iotFhirDestination,
			Constraints: []validation.Constraint{{Target: "iotFhirDestination.IotFhirDestinationProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "iotFhirDestination.IotFhirDestinationProperties.FhirServiceResourceID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "iotFhirDestination.IotFhirDestinationProperties.FhirMapping", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("healthcareapis.IotConnectorFhirDestinationClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, iotConnectorName, fhirDestinationName, iotFhirDestination)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.IotConnectorFhirDestinationClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.IotConnectorFhirDestinationClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IotConnectorFhirDestinationClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, fhirDestinationName string, iotFhirDestination IotFhirDestination) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fhirDestinationName": autorest.Encode("path", fhirDestinationName),
		"iotConnectorName":    autorest.Encode("path", iotConnectorName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}/fhirdestinations/{fhirDestinationName}", pathParameters),
		autorest.WithJSON(iotFhirDestination),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IotConnectorFhirDestinationClient) CreateOrUpdateSender(req *http.Request) (future IotConnectorFhirDestinationCreateOrUpdateFuture, err error) {
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
func (client IotConnectorFhirDestinationClient) CreateOrUpdateResponder(resp *http.Response) (result IotFhirDestination, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an IoT Connector FHIR destination.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// workspaceName - the name of workspace resource.
// iotConnectorName - the name of IoT Connector resource.
// fhirDestinationName - the name of IoT Connector FHIR destination resource.
func (client IotConnectorFhirDestinationClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, fhirDestinationName string) (result IotConnectorFhirDestinationDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotConnectorFhirDestinationClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: iotConnectorName,
			Constraints: []validation.Constraint{{Target: "iotConnectorName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "iotConnectorName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: fhirDestinationName,
			Constraints: []validation.Constraint{{Target: "fhirDestinationName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "fhirDestinationName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("healthcareapis.IotConnectorFhirDestinationClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, iotConnectorName, fhirDestinationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.IotConnectorFhirDestinationClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.IotConnectorFhirDestinationClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IotConnectorFhirDestinationClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, fhirDestinationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fhirDestinationName": autorest.Encode("path", fhirDestinationName),
		"iotConnectorName":    autorest.Encode("path", iotConnectorName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}/fhirdestinations/{fhirDestinationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IotConnectorFhirDestinationClient) DeleteSender(req *http.Request) (future IotConnectorFhirDestinationDeleteFuture, err error) {
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
func (client IotConnectorFhirDestinationClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified Iot Connector FHIR destination.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// workspaceName - the name of workspace resource.
// iotConnectorName - the name of IoT Connector resource.
// fhirDestinationName - the name of IoT Connector FHIR destination resource.
func (client IotConnectorFhirDestinationClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, fhirDestinationName string) (result IotFhirDestination, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotConnectorFhirDestinationClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: iotConnectorName,
			Constraints: []validation.Constraint{{Target: "iotConnectorName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "iotConnectorName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: fhirDestinationName,
			Constraints: []validation.Constraint{{Target: "fhirDestinationName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "fhirDestinationName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("healthcareapis.IotConnectorFhirDestinationClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, iotConnectorName, fhirDestinationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.IotConnectorFhirDestinationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "healthcareapis.IotConnectorFhirDestinationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "healthcareapis.IotConnectorFhirDestinationClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client IotConnectorFhirDestinationClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, fhirDestinationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fhirDestinationName": autorest.Encode("path", fhirDestinationName),
		"iotConnectorName":    autorest.Encode("path", iotConnectorName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}/fhirdestinations/{fhirDestinationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IotConnectorFhirDestinationClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IotConnectorFhirDestinationClient) GetResponder(resp *http.Response) (result IotFhirDestination, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}