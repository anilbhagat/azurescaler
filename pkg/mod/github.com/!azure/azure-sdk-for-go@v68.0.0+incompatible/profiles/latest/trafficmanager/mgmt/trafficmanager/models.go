//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package trafficmanager

import original "github.com/Azure/azure-sdk-for-go/services/trafficmanager/mgmt/2018-08-01/trafficmanager"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AllowedEndpointRecordType = original.AllowedEndpointRecordType

const (
	AllowedEndpointRecordTypeAny         AllowedEndpointRecordType = original.AllowedEndpointRecordTypeAny
	AllowedEndpointRecordTypeDomainName  AllowedEndpointRecordType = original.AllowedEndpointRecordTypeDomainName
	AllowedEndpointRecordTypeIPv4Address AllowedEndpointRecordType = original.AllowedEndpointRecordTypeIPv4Address
	AllowedEndpointRecordTypeIPv6Address AllowedEndpointRecordType = original.AllowedEndpointRecordTypeIPv6Address
)

type EndpointMonitorStatus = original.EndpointMonitorStatus

const (
	EndpointMonitorStatusCheckingEndpoint EndpointMonitorStatus = original.EndpointMonitorStatusCheckingEndpoint
	EndpointMonitorStatusDegraded         EndpointMonitorStatus = original.EndpointMonitorStatusDegraded
	EndpointMonitorStatusDisabled         EndpointMonitorStatus = original.EndpointMonitorStatusDisabled
	EndpointMonitorStatusInactive         EndpointMonitorStatus = original.EndpointMonitorStatusInactive
	EndpointMonitorStatusOnline           EndpointMonitorStatus = original.EndpointMonitorStatusOnline
	EndpointMonitorStatusStopped          EndpointMonitorStatus = original.EndpointMonitorStatusStopped
)

type EndpointStatus = original.EndpointStatus

const (
	EndpointStatusDisabled EndpointStatus = original.EndpointStatusDisabled
	EndpointStatusEnabled  EndpointStatus = original.EndpointStatusEnabled
)

type MonitorProtocol = original.MonitorProtocol

const (
	MonitorProtocolHTTP  MonitorProtocol = original.MonitorProtocolHTTP
	MonitorProtocolHTTPS MonitorProtocol = original.MonitorProtocolHTTPS
	MonitorProtocolTCP   MonitorProtocol = original.MonitorProtocolTCP
)

type ProfileMonitorStatus = original.ProfileMonitorStatus

const (
	ProfileMonitorStatusCheckingEndpoints ProfileMonitorStatus = original.ProfileMonitorStatusCheckingEndpoints
	ProfileMonitorStatusDegraded          ProfileMonitorStatus = original.ProfileMonitorStatusDegraded
	ProfileMonitorStatusDisabled          ProfileMonitorStatus = original.ProfileMonitorStatusDisabled
	ProfileMonitorStatusInactive          ProfileMonitorStatus = original.ProfileMonitorStatusInactive
	ProfileMonitorStatusOnline            ProfileMonitorStatus = original.ProfileMonitorStatusOnline
)

type ProfileStatus = original.ProfileStatus

const (
	ProfileStatusDisabled ProfileStatus = original.ProfileStatusDisabled
	ProfileStatusEnabled  ProfileStatus = original.ProfileStatusEnabled
)

type TrafficRoutingMethod = original.TrafficRoutingMethod

const (
	TrafficRoutingMethodGeographic  TrafficRoutingMethod = original.TrafficRoutingMethodGeographic
	TrafficRoutingMethodMultiValue  TrafficRoutingMethod = original.TrafficRoutingMethodMultiValue
	TrafficRoutingMethodPerformance TrafficRoutingMethod = original.TrafficRoutingMethodPerformance
	TrafficRoutingMethodPriority    TrafficRoutingMethod = original.TrafficRoutingMethodPriority
	TrafficRoutingMethodSubnet      TrafficRoutingMethod = original.TrafficRoutingMethodSubnet
	TrafficRoutingMethodWeighted    TrafficRoutingMethod = original.TrafficRoutingMethodWeighted
)

type TrafficViewEnrollmentStatus = original.TrafficViewEnrollmentStatus

const (
	TrafficViewEnrollmentStatusDisabled TrafficViewEnrollmentStatus = original.TrafficViewEnrollmentStatusDisabled
	TrafficViewEnrollmentStatusEnabled  TrafficViewEnrollmentStatus = original.TrafficViewEnrollmentStatusEnabled
)

type BaseClient = original.BaseClient
type CheckTrafficManagerRelativeDNSNameAvailabilityParameters = original.CheckTrafficManagerRelativeDNSNameAvailabilityParameters
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type DNSConfig = original.DNSConfig
type DeleteOperationResult = original.DeleteOperationResult
type Endpoint = original.Endpoint
type EndpointProperties = original.EndpointProperties
type EndpointPropertiesCustomHeadersItem = original.EndpointPropertiesCustomHeadersItem
type EndpointPropertiesSubnetsItem = original.EndpointPropertiesSubnetsItem
type EndpointsClient = original.EndpointsClient
type GeographicHierarchiesClient = original.GeographicHierarchiesClient
type GeographicHierarchy = original.GeographicHierarchy
type GeographicHierarchyProperties = original.GeographicHierarchyProperties
type HeatMapClient = original.HeatMapClient
type HeatMapEndpoint = original.HeatMapEndpoint
type HeatMapModel = original.HeatMapModel
type HeatMapProperties = original.HeatMapProperties
type MonitorConfig = original.MonitorConfig
type MonitorConfigCustomHeadersItem = original.MonitorConfigCustomHeadersItem
type MonitorConfigExpectedStatusCodeRangesItem = original.MonitorConfigExpectedStatusCodeRangesItem
type NameAvailability = original.NameAvailability
type Profile = original.Profile
type ProfileListResult = original.ProfileListResult
type ProfileProperties = original.ProfileProperties
type ProfilesClient = original.ProfilesClient
type ProxyResource = original.ProxyResource
type QueryExperience = original.QueryExperience
type Region = original.Region
type Resource = original.Resource
type TrackedResource = original.TrackedResource
type TrafficFlow = original.TrafficFlow
type UserMetricsKeysClient = original.UserMetricsKeysClient
type UserMetricsModel = original.UserMetricsModel
type UserMetricsProperties = original.UserMetricsProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return original.NewEndpointsClient(subscriptionID)
}
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewGeographicHierarchiesClient(subscriptionID string) GeographicHierarchiesClient {
	return original.NewGeographicHierarchiesClient(subscriptionID)
}
func NewGeographicHierarchiesClientWithBaseURI(baseURI string, subscriptionID string) GeographicHierarchiesClient {
	return original.NewGeographicHierarchiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewHeatMapClient(subscriptionID string) HeatMapClient {
	return original.NewHeatMapClient(subscriptionID)
}
func NewHeatMapClientWithBaseURI(baseURI string, subscriptionID string) HeatMapClient {
	return original.NewHeatMapClientWithBaseURI(baseURI, subscriptionID)
}
func NewProfilesClient(subscriptionID string) ProfilesClient {
	return original.NewProfilesClient(subscriptionID)
}
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return original.NewProfilesClientWithBaseURI(baseURI, subscriptionID)
}
func NewUserMetricsKeysClient(subscriptionID string) UserMetricsKeysClient {
	return original.NewUserMetricsKeysClient(subscriptionID)
}
func NewUserMetricsKeysClientWithBaseURI(baseURI string, subscriptionID string) UserMetricsKeysClient {
	return original.NewUserMetricsKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAllowedEndpointRecordTypeValues() []AllowedEndpointRecordType {
	return original.PossibleAllowedEndpointRecordTypeValues()
}
func PossibleEndpointMonitorStatusValues() []EndpointMonitorStatus {
	return original.PossibleEndpointMonitorStatusValues()
}
func PossibleEndpointStatusValues() []EndpointStatus {
	return original.PossibleEndpointStatusValues()
}
func PossibleMonitorProtocolValues() []MonitorProtocol {
	return original.PossibleMonitorProtocolValues()
}
func PossibleProfileMonitorStatusValues() []ProfileMonitorStatus {
	return original.PossibleProfileMonitorStatusValues()
}
func PossibleProfileStatusValues() []ProfileStatus {
	return original.PossibleProfileStatusValues()
}
func PossibleTrafficRoutingMethodValues() []TrafficRoutingMethod {
	return original.PossibleTrafficRoutingMethodValues()
}
func PossibleTrafficViewEnrollmentStatusValues() []TrafficViewEnrollmentStatus {
	return original.PossibleTrafficViewEnrollmentStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
