package delete

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

var (
	subscriptionID     = "19da374f-4530-462b-b8fb-324cfc2cdc80"
	location           = "eastus"
	resourceGroupName  = "sample-resource-group"
	virtualNetworkName = "sample-virtual-network"
	subnetName         = "sample-subnet"
	vmScaleSetName     = "sample-vm-scale-set"
)

func DeleteAzure(stage string) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to create Azure credential: %v", err)
	}

	client, err := armcompute.NewVirtualMachinesClient(subscriptionID, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create VM client: %v", err)
	}

	resp, err := client.BeginDelete(context.Background(), resourceGroupName, stage, nil)
	if err != nil {
		log.Fatalf("Failed to start deleting VM: %v", err)
	}

	_, err = resp.Result(context.Background())
	if err != nil {
		log.Fatalf("Failed to finish deleting VM: %v", err)
	}

	log.Println("Successfully deleted VM")
}
