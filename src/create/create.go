package create

import (
	"azurescaler/v3/common"
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func CreateAzure(stage string) {
	// create a clieecho $GOPATHnt

	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		fmt.Println(err)
		return
	}
	var logger = common.NewLogger("start")
	client := resources.NewGroupsClient("19da374f-4530-462b-b8fb-324cfc2cdc80")
	client.Authorizer = authorizer

	// list all resource groups
	groupListResult, err := client.List(context.Background(), "", nil)
	if err != nil {
		logger.Error(err)
		return
	}

	for _, group := range groupListResult.Values() {
		fmt.Println(*group.Name)
	}
}
