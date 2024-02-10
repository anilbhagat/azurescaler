//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package accesscontrol

import original "github.com/Azure/azure-sdk-for-go/services/preview/synapse/2020-08-01-preview/accesscontrol"

type BaseClient = original.BaseClient
type CheckAccessDecision = original.CheckAccessDecision
type CheckPrincipalAccessRequest = original.CheckPrincipalAccessRequest
type CheckPrincipalAccessResponse = original.CheckPrincipalAccessResponse
type ErrorContract = original.ErrorContract
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type ListString = original.ListString
type ListSynapseRoleDefinition = original.ListSynapseRoleDefinition
type RequiredAction = original.RequiredAction
type RoleAssignmentDetails = original.RoleAssignmentDetails
type RoleAssignmentDetailsList = original.RoleAssignmentDetailsList
type RoleAssignmentRequest = original.RoleAssignmentRequest
type RoleAssignmentsClient = original.RoleAssignmentsClient
type RoleDefinitionsClient = original.RoleDefinitionsClient
type SubjectInfo = original.SubjectInfo
type SynapseRbacPermission = original.SynapseRbacPermission
type SynapseRoleDefinition = original.SynapseRoleDefinition

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewRoleAssignmentsClient(endpoint string) RoleAssignmentsClient {
	return original.NewRoleAssignmentsClient(endpoint)
}
func NewRoleDefinitionsClient(endpoint string) RoleDefinitionsClient {
	return original.NewRoleDefinitionsClient(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}