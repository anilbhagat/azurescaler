package receive

import (
    "fmt"
    log "log"
    time "time"
    context "context"
    eventhub "github.com/Azure/azure-event-hubs-go"

)



func ReceiveMessage( stage string) {
    fmt.Println("Started received event from Event Hubs")

	hub, err := eventhub.NewHubFromConnectionString("Endpoint=sb://testscaler.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=JcmzRS3FMoHgehT8wrG2dOiwPx8dMAnm0+AEhPYTC9M=")
    if err != nil {
        log.Fatalf("failed to create hub from connection string: %s\n", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    // Define a handler function for processing incoming events
    handler := func(ctx context.Context, event *eventhub.Event) error {
        fmt.Printf("Received: %s\n", string(event.Data))
        return nil
    }

    // Start receiving events
    //eventhub.ConsumerGroupDefault = "$Default"
    _, err = hub.Receive(ctx, "$Default", handler)
    if err != nil {
        log.Fatalf("failed to receive event: %s\n", err)
    }

    err = hub.Close(context.Background())
    if err != nil {
        log.Fatalf("failed to close hub: %s\n", err)
    }
    

    fmt.Println("Successfully received event from Event Hubs")
}