package engagementfabric

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/engagementfabric/mgmt/2018-09-01/engagementfabric"

// Account the EngagementFabric account
type Account struct {
	autorest.Response `json:"-"`
	// Location - The location of the resource
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource
	Tags map[string]*string `json:"tags"`
	// Sku - The SKU of the resource
	Sku *SKU `json:"sku,omitempty"`
	// ID - READ-ONLY; The ID of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The fully qualified type of the resource
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	if a.Sku != nil {
		objectMap["sku"] = a.Sku
	}
	return json.Marshal(objectMap)
}

// AccountList the list of the EngagementFabric accounts
type AccountList struct {
	autorest.Response `json:"-"`
	// Value - EngagementFabric accounts
	Value *[]Account `json:"value,omitempty"`
}

// AccountPatch the patch of EngagementFabric account
type AccountPatch struct {
	// Tags - The tags of the resource
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for AccountPatch.
func (ap AccountPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ap.Tags != nil {
		objectMap["tags"] = ap.Tags
	}
	return json.Marshal(objectMap)
}

// Channel the EngagementFabric channel
type Channel struct {
	autorest.Response `json:"-"`
	// ChannelProperties - The properties of the channel
	*ChannelProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The ID of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The fully qualified type of the resource
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Channel.
func (c Channel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if c.ChannelProperties != nil {
		objectMap["properties"] = c.ChannelProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Channel struct.
func (c *Channel) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var channelProperties ChannelProperties
				err = json.Unmarshal(*v, &channelProperties)
				if err != nil {
					return err
				}
				c.ChannelProperties = &channelProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				c.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				c.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				c.Type = &typeVar
			}
		}
	}

	return nil
}

// ChannelList the list of the EngagementFabric channels
type ChannelList struct {
	autorest.Response `json:"-"`
	// Value - EngagementFabric channels
	Value *[]Channel `json:"value,omitempty"`
}

// ChannelProperties the EngagementFabric channel properties
type ChannelProperties struct {
	// ChannelType - The channel type
	ChannelType *string `json:"channelType,omitempty"`
	// ChannelFunctions - The functions to be enabled for the channel
	ChannelFunctions *[]string `json:"channelFunctions,omitempty"`
	// Credentials - The channel credentials
	Credentials map[string]*string `json:"credentials"`
}

// MarshalJSON is the custom marshaler for ChannelProperties.
func (cp ChannelProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cp.ChannelType != nil {
		objectMap["channelType"] = cp.ChannelType
	}
	if cp.ChannelFunctions != nil {
		objectMap["channelFunctions"] = cp.ChannelFunctions
	}
	if cp.Credentials != nil {
		objectMap["credentials"] = cp.Credentials
	}
	return json.Marshal(objectMap)
}

// ChannelTypeDescription engagementFabric channel description
type ChannelTypeDescription struct {
	// ChannelType - Channel type
	ChannelType *string `json:"channelType,omitempty"`
	// ChannelDescription - Text description for the channel
	ChannelDescription *string `json:"channelDescription,omitempty"`
	// ChannelFunctions - All the available functions for the channel
	ChannelFunctions *[]string `json:"channelFunctions,omitempty"`
}

// ChannelTypeDescriptionList list of the EngagementFabric channel descriptions
type ChannelTypeDescriptionList struct {
	autorest.Response `json:"-"`
	// Value - Channel descriptions
	Value *[]ChannelTypeDescription `json:"value,omitempty"`
}

// CheckNameAvailabilityParameter the parameter for name availability check
type CheckNameAvailabilityParameter struct {
	// Name - The name to be checked
	Name *string `json:"name,omitempty"`
	// Type - The fully qualified resource type for the name to be checked
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResult the result of name availability check
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	// NameAvailable - READ-ONLY; The name to be checked
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - READ-ONLY; The reason if name is unavailable. Possible values include: 'Invalid', 'AlreadyExists'
	Reason CheckNameUnavailableReason `json:"reason,omitempty"`
	// Message - READ-ONLY; The message if name is unavailable
	Message *string `json:"message,omitempty"`
}

// MarshalJSON is the custom marshaler for CheckNameAvailabilityResult.
func (cnar CheckNameAvailabilityResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// CloudError the default error response
type CloudError struct {
	// Error - Content of the error
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody content of the default error response
type CloudErrorBody struct {
	// Code - The error code
	Code *string `json:"code,omitempty"`
	// Message - The error message
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error
	Target *string `json:"target,omitempty"`
	// Details - The list of additional details
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// KeyDescription the description of the EngagementFabric account key
type KeyDescription struct {
	autorest.Response `json:"-"`
	// Name - READ-ONLY; The name of the key
	Name *string `json:"name,omitempty"`
	// Rank - READ-ONLY; The rank of the key. Possible values include: 'PrimaryKey', 'SecondaryKey'
	Rank KeyRank `json:"rank,omitempty"`
	// Value - READ-ONLY; The value of the key
	Value *string `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for KeyDescription.
func (kd KeyDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// KeyDescriptionList the list of the EngagementFabric account keys
type KeyDescriptionList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Account keys
	Value *[]KeyDescription `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for KeyDescriptionList.
func (kdl KeyDescriptionList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Operation the EngagementFabric operation
type Operation struct {
	// Name - READ-ONLY; The name of the EngagementFabric operation
	Name *string `json:"name,omitempty"`
	// Display - READ-ONLY; The display content of the EngagementFabric operation
	Display *OperationDisplay `json:"display,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// OperationDisplay the display information of the EngagementFabric operation
type OperationDisplay struct {
	// Provider - READ-ONLY; The resource provider namespace of the EngagementFabric operation
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; The resource type of the EngagementFabric operation
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; The name of the EngagementFabric operation
	Operation *string `json:"operation,omitempty"`
	// Description - READ-ONLY; The description of the EngagementFabric operation
	Description *string `json:"description,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationDisplay.
func (od OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// OperationList the list of the EngagementFabric operations
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The EngagementFabric operations
	Value *[]Operation `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationList.
func (ol OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ProxyOnlyResource the base model for the proxy-only Azure resource
type ProxyOnlyResource struct {
	// ID - READ-ONLY; The ID of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The fully qualified type of the resource
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ProxyOnlyResource.
func (por ProxyOnlyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// RegenerateKeyParameter the parameter to regenerate single EngagementFabric account key
type RegenerateKeyParameter struct {
	// Name - The name of key to be regenerated
	Name *string `json:"name,omitempty"`
	// Rank - The rank of the key to be regenerated. Possible values include: 'PrimaryKey', 'SecondaryKey'
	Rank KeyRank `json:"rank,omitempty"`
}

// Resource the base model for Azure resource
type Resource struct {
	// ID - READ-ONLY; The ID of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The fully qualified type of the resource
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// SKU the EngagementFabric SKU
type SKU struct {
	// Name - The name of the SKU
	Name *string `json:"name,omitempty"`
	// Tier - The price tier of the SKU
	Tier *string `json:"tier,omitempty"`
}

// SkuDescription the EngagementFabric SKU description of given resource type
type SkuDescription struct {
	// ResourceType - READ-ONLY; The fully qualified resource type
	ResourceType *string `json:"resourceType,omitempty"`
	// Name - READ-ONLY; The name of the SKU
	Name *string `json:"name,omitempty"`
	// Tier - READ-ONLY; The price tier of the SKU
	Tier *string `json:"tier,omitempty"`
	// Locations - READ-ONLY; The set of locations that the SKU is available
	Locations *[]string `json:"locations,omitempty"`
	// LocationInfo - READ-ONLY; Locations and zones
	LocationInfo *[]SkuLocationInfoItem `json:"locationInfo,omitempty"`
	// Restrictions - READ-ONLY; The restrictions because of which SKU cannot be used
	Restrictions *[]interface{} `json:"restrictions,omitempty"`
}

// MarshalJSON is the custom marshaler for SkuDescription.
func (sd SkuDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// SkuDescriptionList the list of the EngagementFabric SKU descriptions
type SkuDescriptionList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; SKU descriptions
	Value *[]SkuDescription `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for SkuDescriptionList.
func (sdl SkuDescriptionList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// SkuLocationInfoItem the Locations and zones info for SKU
type SkuLocationInfoItem struct {
	// Location - The available location of the SKU
	Location *string `json:"location,omitempty"`
	// Zones - The available zone of the SKU
	Zones *[]string `json:"zones,omitempty"`
}

// TrackedResource the base model for the tracked Azure resource
type TrackedResource struct {
	// Location - The location of the resource
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource
	Tags map[string]*string `json:"tags"`
	// Sku - The SKU of the resource
	Sku *SKU `json:"sku,omitempty"`
	// ID - READ-ONLY; The ID of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The fully qualified type of the resource
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Sku != nil {
		objectMap["sku"] = tr.Sku
	}
	return json.Marshal(objectMap)
}
