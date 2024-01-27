package cloudfunction

import (
	functions "cloud.google.com/go/functions/apiv2"
	"context"
)

func createClient(ctx context.Context) *functions.FunctionClient {
	client, err := functions.NewFunctionClient(ctx)
	if err != nil {
		panic(err)
	}

	return client
}
