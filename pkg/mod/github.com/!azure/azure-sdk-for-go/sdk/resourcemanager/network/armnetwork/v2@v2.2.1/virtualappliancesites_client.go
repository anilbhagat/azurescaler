//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualApplianceSitesClient contains the methods for the VirtualApplianceSites group.
// Don't use this type directly, use NewVirtualApplianceSitesClient() instead.
type VirtualApplianceSitesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualApplianceSitesClient creates a new instance of VirtualApplianceSitesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualApplianceSitesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualApplianceSitesClient, error) {
	cl, err := arm.NewClient(moduleName+".VirtualApplianceSitesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualApplianceSitesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the specified Network Virtual Appliance Site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of the Network Virtual Appliance.
//   - siteName - The name of the site.
//   - parameters - Parameters supplied to the create or update Network Virtual Appliance Site operation.
//   - options - VirtualApplianceSitesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualApplianceSitesClient.BeginCreateOrUpdate
//     method.
func (client *VirtualApplianceSitesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite, options *VirtualApplianceSitesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualApplianceSitesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkVirtualApplianceName, siteName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualApplianceSitesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualApplianceSitesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates the specified Network Virtual Appliance Site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VirtualApplianceSitesClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite, options *VirtualApplianceSitesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, siteName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualApplianceSitesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite, options *VirtualApplianceSitesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified site from a Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of the Network Virtual Appliance.
//   - siteName - The name of the site.
//   - options - VirtualApplianceSitesClientBeginDeleteOptions contains the optional parameters for the VirtualApplianceSitesClient.BeginDelete
//     method.
func (client *VirtualApplianceSitesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesClientBeginDeleteOptions) (*runtime.Poller[VirtualApplianceSitesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkVirtualApplianceName, siteName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualApplianceSitesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualApplianceSitesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified site from a Virtual Appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VirtualApplianceSitesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, siteName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualApplianceSitesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Virtual Appliance Site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of the Network Virtual Appliance.
//   - siteName - The name of the site.
//   - options - VirtualApplianceSitesClientGetOptions contains the optional parameters for the VirtualApplianceSitesClient.Get
//     method.
func (client *VirtualApplianceSitesClient) Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesClientGetOptions) (VirtualApplianceSitesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, siteName, options)
	if err != nil {
		return VirtualApplianceSitesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualApplianceSitesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualApplianceSitesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualApplianceSitesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualApplianceSitesClient) getHandleResponse(resp *http.Response) (VirtualApplianceSitesClientGetResponse, error) {
	result := VirtualApplianceSitesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualApplianceSite); err != nil {
		return VirtualApplianceSitesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all Network Virtual Appliance Sites in a Network Virtual Appliance resource.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of the Network Virtual Appliance.
//   - options - VirtualApplianceSitesClientListOptions contains the optional parameters for the VirtualApplianceSitesClient.NewListPager
//     method.
func (client *VirtualApplianceSitesClient) NewListPager(resourceGroupName string, networkVirtualApplianceName string, options *VirtualApplianceSitesClientListOptions) *runtime.Pager[VirtualApplianceSitesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualApplianceSitesClientListResponse]{
		More: func(page VirtualApplianceSitesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualApplianceSitesClientListResponse) (VirtualApplianceSitesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualApplianceSitesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualApplianceSitesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualApplianceSitesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualApplianceSitesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualApplianceSitesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualApplianceSitesClient) listHandleResponse(resp *http.Response) (VirtualApplianceSitesClientListResponse, error) {
	result := VirtualApplianceSitesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualApplianceSiteListResult); err != nil {
		return VirtualApplianceSitesClientListResponse{}, err
	}
	return result, nil
}
