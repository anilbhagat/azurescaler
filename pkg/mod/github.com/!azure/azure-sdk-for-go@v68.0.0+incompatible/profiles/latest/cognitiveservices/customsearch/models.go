//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package customsearch

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/customsearch"

const (
	DefaultEndpoint = original.DefaultEndpoint
)

type ErrorCode = original.ErrorCode

const (
	InsufficientAuthorization ErrorCode = original.InsufficientAuthorization
	InvalidAuthorization      ErrorCode = original.InvalidAuthorization
	InvalidRequest            ErrorCode = original.InvalidRequest
	None                      ErrorCode = original.None
	RateLimitExceeded         ErrorCode = original.RateLimitExceeded
	ServerError               ErrorCode = original.ServerError
)

type ErrorSubCode = original.ErrorSubCode

const (
	AuthorizationDisabled   ErrorSubCode = original.AuthorizationDisabled
	AuthorizationExpired    ErrorSubCode = original.AuthorizationExpired
	AuthorizationMissing    ErrorSubCode = original.AuthorizationMissing
	AuthorizationRedundancy ErrorSubCode = original.AuthorizationRedundancy
	Blocked                 ErrorSubCode = original.Blocked
	HTTPNotAllowed          ErrorSubCode = original.HTTPNotAllowed
	NotImplemented          ErrorSubCode = original.NotImplemented
	ParameterInvalidValue   ErrorSubCode = original.ParameterInvalidValue
	ParameterMissing        ErrorSubCode = original.ParameterMissing
	ResourceError           ErrorSubCode = original.ResourceError
	UnexpectedError         ErrorSubCode = original.UnexpectedError
)

type SafeSearch = original.SafeSearch

const (
	Moderate SafeSearch = original.Moderate
	Off      SafeSearch = original.Off
	Strict   SafeSearch = original.Strict
)

type TextFormat = original.TextFormat

const (
	HTML TextFormat = original.HTML
	Raw  TextFormat = original.Raw
)

type Type = original.Type

const (
	TypeAnswer              Type = original.TypeAnswer
	TypeCreativeWork        Type = original.TypeCreativeWork
	TypeErrorResponse       Type = original.TypeErrorResponse
	TypeIdentifiable        Type = original.TypeIdentifiable
	TypeResponse            Type = original.TypeResponse
	TypeResponseBase        Type = original.TypeResponseBase
	TypeSearchResponse      Type = original.TypeSearchResponse
	TypeSearchResultsAnswer Type = original.TypeSearchResultsAnswer
	TypeThing               Type = original.TypeThing
	TypeWebPage             Type = original.TypeWebPage
	TypeWebWebAnswer        Type = original.TypeWebWebAnswer
)

type Answer = original.Answer
type BaseClient = original.BaseClient
type BasicAnswer = original.BasicAnswer
type BasicCreativeWork = original.BasicCreativeWork
type BasicIdentifiable = original.BasicIdentifiable
type BasicResponse = original.BasicResponse
type BasicResponseBase = original.BasicResponseBase
type BasicSearchResultsAnswer = original.BasicSearchResultsAnswer
type BasicThing = original.BasicThing
type CreativeWork = original.CreativeWork
type CustomInstanceClient = original.CustomInstanceClient
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type Identifiable = original.Identifiable
type Query = original.Query
type QueryContext = original.QueryContext
type Response = original.Response
type ResponseBase = original.ResponseBase
type SearchResponse = original.SearchResponse
type SearchResultsAnswer = original.SearchResultsAnswer
type Thing = original.Thing
type WebMetaTag = original.WebMetaTag
type WebPage = original.WebPage
type WebWebAnswer = original.WebWebAnswer

func New() BaseClient {
	return original.New()
}
func NewCustomInstanceClient() CustomInstanceClient {
	return original.NewCustomInstanceClient()
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleErrorCodeValues() []ErrorCode {
	return original.PossibleErrorCodeValues()
}
func PossibleErrorSubCodeValues() []ErrorSubCode {
	return original.PossibleErrorSubCodeValues()
}
func PossibleSafeSearchValues() []SafeSearch {
	return original.PossibleSafeSearchValues()
}
func PossibleTextFormatValues() []TextFormat {
	return original.PossibleTextFormatValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
