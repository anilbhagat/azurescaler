// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package managedblockchainqueryiface provides an interface to enable mocking the Amazon Managed Blockchain Query service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package managedblockchainqueryiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/managedblockchainquery"
)

// ManagedBlockchainQueryAPI provides an interface to enable mocking the
// managedblockchainquery.ManagedBlockchainQuery service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Managed Blockchain Query.
//	func myFunc(svc managedblockchainqueryiface.ManagedBlockchainQueryAPI) bool {
//	    // Make svc.BatchGetTokenBalance request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := managedblockchainquery.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockManagedBlockchainQueryClient struct {
//	    managedblockchainqueryiface.ManagedBlockchainQueryAPI
//	}
//	func (m *mockManagedBlockchainQueryClient) BatchGetTokenBalance(input *managedblockchainquery.BatchGetTokenBalanceInput) (*managedblockchainquery.BatchGetTokenBalanceOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockManagedBlockchainQueryClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ManagedBlockchainQueryAPI interface {
	BatchGetTokenBalance(*managedblockchainquery.BatchGetTokenBalanceInput) (*managedblockchainquery.BatchGetTokenBalanceOutput, error)
	BatchGetTokenBalanceWithContext(aws.Context, *managedblockchainquery.BatchGetTokenBalanceInput, ...request.Option) (*managedblockchainquery.BatchGetTokenBalanceOutput, error)
	BatchGetTokenBalanceRequest(*managedblockchainquery.BatchGetTokenBalanceInput) (*request.Request, *managedblockchainquery.BatchGetTokenBalanceOutput)

	GetAssetContract(*managedblockchainquery.GetAssetContractInput) (*managedblockchainquery.GetAssetContractOutput, error)
	GetAssetContractWithContext(aws.Context, *managedblockchainquery.GetAssetContractInput, ...request.Option) (*managedblockchainquery.GetAssetContractOutput, error)
	GetAssetContractRequest(*managedblockchainquery.GetAssetContractInput) (*request.Request, *managedblockchainquery.GetAssetContractOutput)

	GetTokenBalance(*managedblockchainquery.GetTokenBalanceInput) (*managedblockchainquery.GetTokenBalanceOutput, error)
	GetTokenBalanceWithContext(aws.Context, *managedblockchainquery.GetTokenBalanceInput, ...request.Option) (*managedblockchainquery.GetTokenBalanceOutput, error)
	GetTokenBalanceRequest(*managedblockchainquery.GetTokenBalanceInput) (*request.Request, *managedblockchainquery.GetTokenBalanceOutput)

	GetTransaction(*managedblockchainquery.GetTransactionInput) (*managedblockchainquery.GetTransactionOutput, error)
	GetTransactionWithContext(aws.Context, *managedblockchainquery.GetTransactionInput, ...request.Option) (*managedblockchainquery.GetTransactionOutput, error)
	GetTransactionRequest(*managedblockchainquery.GetTransactionInput) (*request.Request, *managedblockchainquery.GetTransactionOutput)

	ListAssetContracts(*managedblockchainquery.ListAssetContractsInput) (*managedblockchainquery.ListAssetContractsOutput, error)
	ListAssetContractsWithContext(aws.Context, *managedblockchainquery.ListAssetContractsInput, ...request.Option) (*managedblockchainquery.ListAssetContractsOutput, error)
	ListAssetContractsRequest(*managedblockchainquery.ListAssetContractsInput) (*request.Request, *managedblockchainquery.ListAssetContractsOutput)

	ListAssetContractsPages(*managedblockchainquery.ListAssetContractsInput, func(*managedblockchainquery.ListAssetContractsOutput, bool) bool) error
	ListAssetContractsPagesWithContext(aws.Context, *managedblockchainquery.ListAssetContractsInput, func(*managedblockchainquery.ListAssetContractsOutput, bool) bool, ...request.Option) error

	ListTokenBalances(*managedblockchainquery.ListTokenBalancesInput) (*managedblockchainquery.ListTokenBalancesOutput, error)
	ListTokenBalancesWithContext(aws.Context, *managedblockchainquery.ListTokenBalancesInput, ...request.Option) (*managedblockchainquery.ListTokenBalancesOutput, error)
	ListTokenBalancesRequest(*managedblockchainquery.ListTokenBalancesInput) (*request.Request, *managedblockchainquery.ListTokenBalancesOutput)

	ListTokenBalancesPages(*managedblockchainquery.ListTokenBalancesInput, func(*managedblockchainquery.ListTokenBalancesOutput, bool) bool) error
	ListTokenBalancesPagesWithContext(aws.Context, *managedblockchainquery.ListTokenBalancesInput, func(*managedblockchainquery.ListTokenBalancesOutput, bool) bool, ...request.Option) error

	ListTransactionEvents(*managedblockchainquery.ListTransactionEventsInput) (*managedblockchainquery.ListTransactionEventsOutput, error)
	ListTransactionEventsWithContext(aws.Context, *managedblockchainquery.ListTransactionEventsInput, ...request.Option) (*managedblockchainquery.ListTransactionEventsOutput, error)
	ListTransactionEventsRequest(*managedblockchainquery.ListTransactionEventsInput) (*request.Request, *managedblockchainquery.ListTransactionEventsOutput)

	ListTransactionEventsPages(*managedblockchainquery.ListTransactionEventsInput, func(*managedblockchainquery.ListTransactionEventsOutput, bool) bool) error
	ListTransactionEventsPagesWithContext(aws.Context, *managedblockchainquery.ListTransactionEventsInput, func(*managedblockchainquery.ListTransactionEventsOutput, bool) bool, ...request.Option) error

	ListTransactions(*managedblockchainquery.ListTransactionsInput) (*managedblockchainquery.ListTransactionsOutput, error)
	ListTransactionsWithContext(aws.Context, *managedblockchainquery.ListTransactionsInput, ...request.Option) (*managedblockchainquery.ListTransactionsOutput, error)
	ListTransactionsRequest(*managedblockchainquery.ListTransactionsInput) (*request.Request, *managedblockchainquery.ListTransactionsOutput)

	ListTransactionsPages(*managedblockchainquery.ListTransactionsInput, func(*managedblockchainquery.ListTransactionsOutput, bool) bool) error
	ListTransactionsPagesWithContext(aws.Context, *managedblockchainquery.ListTransactionsInput, func(*managedblockchainquery.ListTransactionsOutput, bool) bool, ...request.Option) error
}

var _ ManagedBlockchainQueryAPI = (*managedblockchainquery.ManagedBlockchainQuery)(nil)