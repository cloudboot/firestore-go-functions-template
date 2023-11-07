package cloudboot

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/googleapis/google-cloudevents-go/cloud/firestoredata"
	"google.golang.org/protobuf/proto"
)

func init() {
	functions.CloudEvent("Main", main)
}

func main(ctx context.Context, event event.Event) error {
	var data firestoredata.DocumentEventData
	if err := proto.Unmarshal(event.Data(), &data); err != nil {
		return fmt.Errorf("proto.Unmarshal: %w", err)
	}

	fmt.Printf("Function triggered by change to: %v\n", event.Source())
	fmt.Printf("Old value: %+v\n", data.GetOldValue())
	fmt.Printf("New value: %+v\n", data.GetValue())
	return nil
}
