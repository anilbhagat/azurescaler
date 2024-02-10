package storagedatalakeapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/storage/datalake/2018-06-17/storagedatalake"
	"github.com/Azure/go-autorest/autorest"
	"io"
)

// FilesystemClientAPI contains the set of methods on the FilesystemClient type.
type FilesystemClientAPI interface {
	Create(ctx context.Context, filesystem string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
	Delete(ctx context.Context, filesystem string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
	GetProperties(ctx context.Context, filesystem string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
	List(ctx context.Context, prefix string, continuation string, maxResults *int32, xMsClientRequestID string, timeout *int32, xMsDate string) (result storagedatalake.FilesystemList, err error)
	SetProperties(ctx context.Context, filesystem string, xMsProperties string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
}

var _ FilesystemClientAPI = (*storagedatalake.FilesystemClient)(nil)

// PathClientAPI contains the set of methods on the PathClient type.
type PathClientAPI interface {
	Create(ctx context.Context, filesystem string, pathParameter string, resource storagedatalake.PathResourceType, continuation string, mode storagedatalake.PathRenameMode, cacheControl string, contentEncoding string, contentLanguage string, contentDisposition string, xMsCacheControl string, xMsContentType string, xMsContentEncoding string, xMsContentLanguage string, xMsContentDisposition string, xMsRenameSource string, xMsLeaseID string, xMsProposedLeaseID string, xMsSourceLeaseID string, xMsProperties string, xMsPermissions string, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, xMsSourceIfMatch string, xMsSourceIfNoneMatch string, xMsSourceIfModifiedSince string, xMsSourceIfUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
	Delete(ctx context.Context, filesystem string, pathParameter string, recursive *bool, continuation string, xMsLeaseID string, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
	GetProperties(ctx context.Context, filesystem string, pathParameter string, action storagedatalake.PathGetPropertiesAction, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
	Lease(ctx context.Context, xMsLeaseAction storagedatalake.PathLeaseAction, filesystem string, pathParameter string, xMsLeaseDuration *int32, xMsLeaseBreakPeriod *int32, xMsLeaseID string, xMsProposedLeaseID string, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
	List(ctx context.Context, recursive bool, filesystem string, directory string, continuation string, maxResults *int32, xMsClientRequestID string, timeout *int32, xMsDate string) (result storagedatalake.PathList, err error)
	Read(ctx context.Context, filesystem string, pathParameter string, rangeParameter string, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result storagedatalake.ReadCloser, err error)
	Update(ctx context.Context, action storagedatalake.PathUpdateAction, filesystem string, pathParameter string, position *int64, retainUncommittedData *bool, contentLength string, xMsLeaseAction storagedatalake.PathUpdateLeaseAction, xMsLeaseID string, xMsCacheControl string, xMsContentType string, xMsContentDisposition string, xMsContentEncoding string, xMsContentLanguage string, xMsProperties string, xMsOwner string, xMsGroup string, xMsPermissions string, xMsACL string, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, requestBody io.ReadCloser, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
}

var _ PathClientAPI = (*storagedatalake.PathClient)(nil)