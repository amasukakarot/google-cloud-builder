package cloudfunction

import (
	"context"
	"log"

	functionspb "cloud.google.com/go/functions/apiv2/functionspb"
	"google.golang.org/api/iterator"
)

func IfFunctionExists(ctx context.Context, functionName string) bool {

	client := createClient(ctx)
	defer client.Close()

	req := &functionspb.ListFunctionsRequest{
		Parent: "projects/groovy-iris-412518/locations/europe-west2",
	}
	it := client.ListFunctions(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			log.Println("Finished iterating through functions...")
			break
		}
		if err != nil {
			panic(err)
		}
		if resp.Name == functionName {
			log.Printf("Function %v found.", functionName)
			return true
		}
	}
	log.Printf("Function %v not found.", functionName)
	return false
}
