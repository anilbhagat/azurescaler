// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package netappapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/netapp/mgmt/2020-05-01/netapp"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result netapp.OperationListResult, err error)
}

var _ OperationsClientAPI = (*netapp.OperationsClient)(nil)

// ResourceClientAPI contains the set of methods on the ResourceClient type.
type ResourceClientAPI interface {
	CheckFilePathAvailability(ctx context.Context, body netapp.ResourceNameAvailabilityRequest, location string) (result netapp.CheckAvailabilityResponse, err error)
	CheckNameAvailability(ctx context.Context, body netapp.ResourceNameAvailabilityRequest, location string) (result netapp.CheckAvailabilityResponse, err error)
	CheckQuotaAvailability(ctx context.Context, body netapp.QuotaAvailabilityRequest, location string) (result netapp.CheckAvailabilityResponse, err error)
}

var _ ResourceClientAPI = (*netapp.ResourceClient)(nil)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CreateOrUpdate(ctx context.Context, body netapp.Account, resourceGroupName string, accountName string) (result netapp.AccountsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result netapp.AccountsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result netapp.Account, err error)
	List(ctx context.Context, resourceGroupName string) (result netapp.AccountList, err error)
	Update(ctx context.Context, body netapp.AccountPatch, resourceGroupName string, accountName string) (result netapp.AccountsUpdateFuture, err error)
}

var _ AccountsClientAPI = (*netapp.AccountsClient)(nil)

// PoolsClientAPI contains the set of methods on the PoolsClient type.
type PoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, body netapp.CapacityPool, resourceGroupName string, accountName string, poolName string) (result netapp.PoolsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result netapp.PoolsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result netapp.CapacityPool, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result netapp.CapacityPoolList, err error)
	Update(ctx context.Context, body netapp.CapacityPoolPatch, resourceGroupName string, accountName string, poolName string) (result netapp.PoolsUpdateFuture, err error)
}

var _ PoolsClientAPI = (*netapp.PoolsClient)(nil)

// VolumesClientAPI contains the set of methods on the VolumesClient type.
type VolumesClientAPI interface {
	AuthorizeReplication(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, body netapp.AuthorizeRequest) (result netapp.VolumesAuthorizeReplicationFuture, err error)
	BreakReplication(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, body *netapp.BreakReplicationRequest) (result netapp.VolumesBreakReplicationFuture, err error)
	CreateOrUpdate(ctx context.Context, body netapp.Volume, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.VolumesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.VolumesDeleteFuture, err error)
	DeleteReplication(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.VolumesDeleteReplicationFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.Volume, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result netapp.VolumeList, err error)
	PoolChange(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, body netapp.PoolChangeRequest) (result netapp.VolumesPoolChangeFuture, err error)
	ReInitializeReplication(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.VolumesReInitializeReplicationFuture, err error)
	ReplicationStatusMethod(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.ReplicationStatus, err error)
	ResyncReplication(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.VolumesResyncReplicationFuture, err error)
	Revert(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, body netapp.VolumeRevert) (result netapp.VolumesRevertFuture, err error)
	Update(ctx context.Context, body netapp.VolumePatch, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.VolumesUpdateFuture, err error)
}

var _ VolumesClientAPI = (*netapp.VolumesClient)(nil)

// SnapshotsClientAPI contains the set of methods on the SnapshotsClient type.
type SnapshotsClientAPI interface {
	Create(ctx context.Context, body netapp.Snapshot, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string) (result netapp.SnapshotsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string) (result netapp.SnapshotsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string) (result netapp.Snapshot, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.SnapshotsList, err error)
	Update(ctx context.Context, body interface{}, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string) (result netapp.SnapshotsUpdateFuture, err error)
}

var _ SnapshotsClientAPI = (*netapp.SnapshotsClient)(nil)

// SnapshotPoliciesClientAPI contains the set of methods on the SnapshotPoliciesClient type.
type SnapshotPoliciesClientAPI interface {
	Create(ctx context.Context, body netapp.SnapshotPolicy, resourceGroupName string, accountName string, snapshotPolicyName string) (result netapp.SnapshotPolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string) (result netapp.SnapshotPoliciesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string) (result netapp.SnapshotPolicy, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result netapp.SnapshotPoliciesList, err error)
	ListVolumes(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string) (result netapp.SnapshotPolicyVolumeList, err error)
	Update(ctx context.Context, body netapp.SnapshotPolicyPatch, resourceGroupName string, accountName string, snapshotPolicyName string) (result netapp.SnapshotPolicy, err error)
}

var _ SnapshotPoliciesClientAPI = (*netapp.SnapshotPoliciesClient)(nil)

// AccountBackupsClientAPI contains the set of methods on the AccountBackupsClient type.
type AccountBackupsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, accountName string, backupName string) (result netapp.AccountBackupsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, backupName string) (result netapp.Backup, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result netapp.BackupsList, err error)
}

var _ AccountBackupsClientAPI = (*netapp.AccountBackupsClient)(nil)

// BackupsClientAPI contains the set of methods on the BackupsClient type.
type BackupsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body netapp.Backup) (result netapp.BackupsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string) (result netapp.BackupsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string) (result netapp.Backup, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.BackupsList, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body *netapp.BackupPatch) (result netapp.Backup, err error)
}

var _ BackupsClientAPI = (*netapp.BackupsClient)(nil)

// BackupPoliciesClientAPI contains the set of methods on the BackupPoliciesClient type.
type BackupPoliciesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, backupPolicyName string, body netapp.BackupPolicy) (result netapp.BackupPoliciesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, backupPolicyName string) (result netapp.BackupPoliciesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, backupPolicyName string) (result netapp.BackupPolicy, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result netapp.BackupPoliciesList, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, backupPolicyName string, body netapp.BackupPolicyPatch) (result netapp.BackupPolicy, err error)
}

var _ BackupPoliciesClientAPI = (*netapp.BackupPoliciesClient)(nil)

// VaultsClientAPI contains the set of methods on the VaultsClient type.
type VaultsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, accountName string) (result netapp.VaultList, err error)
}

var _ VaultsClientAPI = (*netapp.VaultsClient)(nil)
