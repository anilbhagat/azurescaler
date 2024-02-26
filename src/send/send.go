package send

import (
	"context"
    "fmt"
    "log"
   // time "time"

    eventhub "github.com/Azure/azure-event-hubs-go"

)


func SendMessage(stage string) {
    fmt.Println("Successfully caaled send api call")

	hub, err := eventhub.NewHubFromConnectionString("Endpoint=sb://testscaler.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=JcmzRS3FMoHgehT8wrG2dOiwPx8dMAnm0+AEhPYTC9M=")
    if err != nil {
        log.Fatalf("failed to create hub from connection string: %s\n", err)
    }
	//ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
   // defer cancel()
// Create an event to send
event := eventhub.NewEventFromString("Hello, Event Hubs!")
  // Send the event
  err = hub.Send(context.Background(), event)
  if err != nil {
      log.Fatalf("failed to send event: %s\n", err)
  }
  err = hub.Close(context.Background())
  if err != nil {
      log.Fatalf("failed to close hub: %s\n", err)
  }

  log.Println("Successfully sent event to Event Hubs")
  fmt.Println("Successfully received event from Event Hubs")


}