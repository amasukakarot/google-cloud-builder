package cloudfunction

import (
	"cloud.google.com/go/functions/apiv2/functionspb"
	"context"
	"log"
)

func UpdateCloudFunction(ctx context.Context, functionName string) {
	log.Printf("Updating %v ", functionName)

	client := createClient(ctx)
	defer client.Close()

	req := &functionspb.UpdateFunctionRequest{
		Function: buildFunctionRequest(),
	}

	op, err := client.UpdateFunction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp

	log.Printf("Function %v updated!", functionName)
}
