package powerplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/powerplatform/mgmt/2020-10-30/powerplatform"

// EnterprisePolicy definition of the EnterprisePolicy.
type EnterprisePolicy struct {
	autorest.Response `json:"-"`
	// Identity - The identity of the EnterprisePolicy.
	Identity *EnterprisePolicyIdentity `json:"identity,omitempty"`
	// Properties - The properties that define configuration for the enterprise policy
	*Properties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// ID - READ-ONLY; ARM resource id of the EnterprisePolicy.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the EnterprisePolicy.
	Name *string `json:"name,omitempty"`
	// Location - READ-ONLY; Region where the EnterprisePolicy is located.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for EnterprisePolicy.
func (ep EnterprisePolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ep.Identity != nil {
		objectMap["identity"] = ep.Identity
	}
	if ep.Properties != nil {
		objectMap["properties"] = ep.Properties
	}
	if ep.Tags != nil {
		objectMap["tags"] = ep.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for EnterprisePolicy struct.
func (ep *EnterprisePolicy) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "identity":
			if v != nil {
				var identity EnterprisePolicyIdentity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				ep.Identity = &identity
			}
		case "properties":
			if v != nil {
				var properties Properties
				err = json.Unmarshal(*v, &properties)
				if err != nil {
					return err
				}
				ep.Properties = &properties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ep.Tags = tags
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ep.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ep.Name = &name
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				ep.Location = &location
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ep.Type = &typeVar
			}
		}
	}

	return nil
}

// EnterprisePolicyIdentity the identity of the EnterprisePolicy.
type EnterprisePolicyIdentity struct {
	// SystemAssignedIdentityPrincipalID - READ-ONLY; The principal id of EnterprisePolicy identity.
	SystemAssignedIdentityPrincipalID *string `json:"systemAssignedIdentityPrincipalId,omitempty"`
	// TenantID - READ-ONLY; The tenant id associated with the EnterprisePolicy.
	TenantID *string `json:"tenantId,omitempty"`
	// Type - The type of identity used for the EnterprisePolicy. Currently, the only supported type is 'SystemAssigned', which implicitly creates an identity. Possible values include: 'SystemAssigned', 'None'
	Type ResourceIdentityType `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for EnterprisePolicyIdentity.
func (epi EnterprisePolicyIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if epi.Type != "" {
		objectMap["type"] = epi.Type
	}
	return json.Marshal(objectMap)
}

// EnterprisePolicyList the response of the list EnterprisePolicy operation.
type EnterprisePolicyList struct {
	autorest.Response `json:"-"`
	// Value - Result of the list EnterprisePolicy operation.
	Value *[]EnterprisePolicy `json:"value,omitempty"`
}

// ErrorResponse ARM error response body.
type ErrorResponse struct {
	// Error - Details about the error.
	Error *ErrorResponseBody `json:"error,omitempty"`
}

// ErrorResponseBody an error response from the PowerPlatform service.
type ErrorResponseBody struct {
	// Code - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`
	// Message - A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
}

// KeyProperties url and version of the KeyVault Secret
type KeyProperties struct {
	// Name - The identifier of the key vault key used to encrypt data.
	Name *string `json:"name,omitempty"`
	// Version - The version of the identity which will be used to access key vault.
	Version *string `json:"version,omitempty"`
}

// KeyVaultProperties settings concerning key vault encryption for a configuration store.
type KeyVaultProperties struct {
	// ID - Uri of KeyVault
	ID *string `json:"id,omitempty"`
	// Key - Identity of the secret that includes name and version.
	Key *KeyProperties `json:"key,omitempty"`
	// Status - The state of onboarding, which only appears in the response. Possible values include: 'Enabled', 'Disabled', 'NotConfigured'
	Status Status `json:"status,omitempty"`
}

// Operation powerPlatform REST API operation
type Operation struct {
	// Name - Operation name: For ex. providers/Microsoft.PowerPlatform/enterprisePolicies/write or read
	Name *string `json:"name,omitempty"`
	// IsDataAction - Indicates whether the operation is a data action
	IsDataAction *string `json:"isDataAction,omitempty"`
	// Display - Provider, Resource, Operation and description values.
	Display *OperationDisplay `json:"display,omitempty"`
	// OperationProperties - Provider, Resource, Operation and description values.
	*OperationProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.Name != nil {
		objectMap["name"] = o.Name
	}
	if o.IsDataAction != nil {
		objectMap["isDataAction"] = o.IsDataAction
	}
	if o.Display != nil {
		objectMap["display"] = o.Display
	}
	if o.OperationProperties != nil {
		objectMap["properties"] = o.OperationProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Operation struct.
func (o *Operation) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				o.Name = &name
			}
		case "isDataAction":
			if v != nil {
				var isDataAction string
				err = json.Unmarshal(*v, &isDataAction)
				if err != nil {
					return err
				}
				o.IsDataAction = &isDataAction
			}
		case "display":
			if v != nil {
				var display OperationDisplay
				err = json.Unmarshal(*v, &display)
				if err != nil {
					return err
				}
				o.Display = &display
			}
		case "properties":
			if v != nil {
				var operationProperties OperationProperties
				err = json.Unmarshal(*v, &operationProperties)
				if err != nil {
					return err
				}
				o.OperationProperties = &operationProperties
			}
		}
	}

	return nil
}

// OperationDisplay provider, Resource, Operation and description values.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.PowerPlatform
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description about operation.
	Description *string `json:"description,omitempty"`
}

// OperationList the response model for the list of PowerPlatform operations
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - List of PowerPlatform operations supported by the PowerPlatform resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// OperationProperties provider, Resource, Operation and description values.
type OperationProperties struct {
	// StatusCode - Service provider: Microsoft.PowerPlatform
	StatusCode *string `json:"statusCode,omitempty"`
}

// PrivateEndpoint the Private Endpoint resource.
type PrivateEndpoint struct {
	// ID - READ-ONLY; The ARM identifier for Private Endpoint
	ID *string `json:"id,omitempty"`
}

// MarshalJSON is the custom marshaler for PrivateEndpoint.
func (peVar PrivateEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// PrivateEndpointConnection a private endpoint connection
type PrivateEndpointConnection struct {
	autorest.Response `json:"-"`
	// PrivateEndpointConnectionProperties - Resource properties.
	*PrivateEndpointConnectionProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; ARM resource id of the EnterprisePolicy.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the EnterprisePolicy.
	Name *string `json:"name,omitempty"`
	// Location - READ-ONLY; Region where the EnterprisePolicy is located.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for PrivateEndpointConnection.
func (pec PrivateEndpointConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pec.PrivateEndpointConnectionProperties != nil {
		objectMap["properties"] = pec.PrivateEndpointConnectionProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PrivateEndpointConnection struct.
func (pec *PrivateEndpointConnection) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var privateEndpointConnectionProperties PrivateEndpointConnectionProperties
				err = json.Unmarshal(*v, &privateEndpointConnectionProperties)
				if err != nil {
					return err
				}
				pec.PrivateEndpointConnectionProperties = &privateEndpointConnectionProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				pec.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				pec.Name = &name
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				pec.Location = &location
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				pec.Type = &typeVar
			}
		}
	}

	return nil
}

// PrivateEndpointConnectionListResult a list of private endpoint connections
type PrivateEndpointConnectionListResult struct {
	autorest.Response `json:"-"`
	// Value - Array of private endpoint connections
	Value *[]PrivateEndpointConnection `json:"value,omitempty"`
}

// PrivateEndpointConnectionProperties properties of the PrivateEndpointConnectProperties.
type PrivateEndpointConnectionProperties struct {
	// PrivateEndpoint - The resource of private end point.
	PrivateEndpoint *PrivateEndpoint `json:"privateEndpoint,omitempty"`
	// PrivateLinkServiceConnectionState - A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`
	// ProvisioningState - The provisioning state of the private endpoint connection resource. Possible values include: 'Succeeded', 'Creating', 'Deleting', 'Failed'
	ProvisioningState PrivateEndpointConnectionProvisioningState `json:"provisioningState,omitempty"`
}

// PrivateEndpointConnectionsCreateOrUpdateFuture an abstraction for monitoring and retrieving the results
// of a long-running operation.
type PrivateEndpointConnectionsCreateOrUpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(PrivateEndpointConnectionsClient) (PrivateEndpointConnection, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *PrivateEndpointConnectionsCreateOrUpdateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for PrivateEndpointConnectionsCreateOrUpdateFuture.Result.
func (future *PrivateEndpointConnectionsCreateOrUpdateFuture) result(client PrivateEndpointConnectionsClient) (pec PrivateEndpointConnection, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerplatform.PrivateEndpointConnectionsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		pec.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("powerplatform.PrivateEndpointConnectionsCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if pec.Response.Response, err = future.GetResult(sender); err == nil && pec.Response.Response.StatusCode != http.StatusNoContent {
		pec, err = client.CreateOrUpdateResponder(pec.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "powerplatform.PrivateEndpointConnectionsCreateOrUpdateFuture", "Result", pec.Response.Response, "Failure responding to request")
		}
	}
	return
}

// PrivateEndpointConnectionsDeleteFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type PrivateEndpointConnectionsDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(PrivateEndpointConnectionsClient) (autorest.Response, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *PrivateEndpointConnectionsDeleteFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for PrivateEndpointConnectionsDeleteFuture.Result.
func (future *PrivateEndpointConnectionsDeleteFuture) result(client PrivateEndpointConnectionsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerplatform.PrivateEndpointConnectionsDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		ar.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("powerplatform.PrivateEndpointConnectionsDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// PrivateLinkResource a private link resource
type PrivateLinkResource struct {
	autorest.Response `json:"-"`
	// PrivateLinkResourceProperties - Resource properties.
	*PrivateLinkResourceProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; ARM resource id of the EnterprisePolicy.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the EnterprisePolicy.
	Name *string `json:"name,omitempty"`
	// Location - READ-ONLY; Region where the EnterprisePolicy is located.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for PrivateLinkResource.
func (plr PrivateLinkResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if plr.PrivateLinkResourceProperties != nil {
		objectMap["properties"] = plr.PrivateLinkResourceProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PrivateLinkResource struct.
func (plr *PrivateLinkResource) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var privateLinkResourceProperties PrivateLinkResourceProperties
				err = json.Unmarshal(*v, &privateLinkResourceProperties)
				if err != nil {
					return err
				}
				plr.PrivateLinkResourceProperties = &privateLinkResourceProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				plr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				plr.Name = &name
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				plr.Location = &location
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				plr.Type = &typeVar
			}
		}
	}

	return nil
}

// PrivateLinkResourceListResult a list of private link resources
type PrivateLinkResourceListResult struct {
	autorest.Response `json:"-"`
	// Value - Array of private link resources
	Value *[]PrivateLinkResource `json:"value,omitempty"`
}

// PrivateLinkResourceProperties properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// GroupID - READ-ONLY; The private link resource group id.
	GroupID *string `json:"groupId,omitempty"`
	// RequiredMembers - READ-ONLY; The private link resource required member names.
	RequiredMembers *[]string `json:"requiredMembers,omitempty"`
	// SystemData - Metadata for the private link resource.
	SystemData *SystemData `json:"systemData,omitempty"`
}

// MarshalJSON is the custom marshaler for PrivateLinkResourceProperties.
func (plrp PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if plrp.SystemData != nil {
		objectMap["systemData"] = plrp.SystemData
	}
	return json.Marshal(objectMap)
}

// PrivateLinkServiceConnectionState a collection of information about the state of the connection between
// service consumer and provider.
type PrivateLinkServiceConnectionState struct {
	// Status - Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service. Possible values include: 'Pending', 'Approved', 'Rejected'
	Status PrivateEndpointServiceConnectionStatus `json:"status,omitempty"`
	// Description - The reason for approval/rejection of the connection.
	Description *string `json:"description,omitempty"`
	// ActionsRequired - A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `json:"actionsRequired,omitempty"`
}

// Properties the properties that define configuration for the enterprise policy.
type Properties struct {
	// Lockbox - Settings concerning lockbox.
	Lockbox *PropertiesLockbox `json:"lockbox,omitempty"`
	// Encryption - The encryption settings for a configuration store.
	Encryption *PropertiesEncryption `json:"encryption,omitempty"`
	// SystemData - Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty"`
}

// PropertiesEncryption the encryption settings for a configuration store.
type PropertiesEncryption struct {
	// KeyVaultProperties - Key vault properties.
	KeyVaultProperties *KeyVaultProperties `json:"keyVaultProperties,omitempty"`
}

// PropertiesLockbox settings concerning lockbox.
type PropertiesLockbox struct {
	// Status - lockbox configuration. Possible values include: 'Enabled', 'Disabled', 'NotConfigured'
	Status Status `json:"status,omitempty"`
}

// ProxyResource ARM proxy resource.
type ProxyResource struct {
	// ID - READ-ONLY; ARM resource id of the EnterprisePolicy.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the EnterprisePolicy.
	Name *string `json:"name,omitempty"`
	// Location - READ-ONLY; Region where the EnterprisePolicy is located.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ProxyResource.
func (pr ProxyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Resource the core properties of ARM resources
type Resource struct {
	// ID - READ-ONLY; ARM resource id of the EnterprisePolicy.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the EnterprisePolicy.
	Name *string `json:"name,omitempty"`
	// Location - READ-ONLY; Region where the EnterprisePolicy is located.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Subnet a subnet
type Subnet struct {
	autorest.Response `json:"-"`
	// SubnetProperties - Resource properties.
	*SubnetProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; ARM resource id of the EnterprisePolicy.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the EnterprisePolicy.
	Name *string `json:"name,omitempty"`
	// Location - READ-ONLY; Region where the EnterprisePolicy is located.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Subnet.
func (s Subnet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.SubnetProperties != nil {
		objectMap["properties"] = s.SubnetProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Subnet struct.
func (s *Subnet) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var subnetProperties SubnetProperties
				err = json.Unmarshal(*v, &subnetProperties)
				if err != nil {
					return err
				}
				s.SubnetProperties = &subnetProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				s.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				s.Name = &name
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				s.Location = &location
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				s.Type = &typeVar
			}
		}
	}

	return nil
}

// SubnetEndpointProperty endpoint of the subnet.
type SubnetEndpointProperty struct {
	// ID - Resource id of the subnet.
	ID *string `json:"id,omitempty"`
}

// SubnetListResult a list of subnets
type SubnetListResult struct {
	autorest.Response `json:"-"`
	// Value - Array of subnets
	Value *[]Subnet `json:"value,omitempty"`
}

// SubnetProperties properties of a subnet.
type SubnetProperties struct {
	// Subnet - Endpoint of the subnet.
	Subnet *SubnetEndpointProperty `json:"subnet,omitempty"`
	// Status - Connection State of the subnet. Possible values include: 'Enabled', 'Disabled', 'NotConfigured'
	Status Status `json:"status,omitempty"`
	// SystemData - Metadata for the subnet.
	SystemData *SystemData `json:"systemData,omitempty"`
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The type of identity that last modified the resource.
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// ID - READ-ONLY; ARM resource id of the EnterprisePolicy.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the EnterprisePolicy.
	Name *string `json:"name,omitempty"`
	// Location - READ-ONLY; Region where the EnterprisePolicy is located.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	return json.Marshal(objectMap)
}