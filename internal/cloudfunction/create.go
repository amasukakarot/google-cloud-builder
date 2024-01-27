package cloudfunction

import (
	"cloud.google.com/go/functions/apiv2/functionspb"
	"context"
	"fmt"
)

func CreateCloudFunction() {
	ctx := context.Background()

	client := createClient(ctx)
	defer client.Close()

	req := &functionspb.CreateFunctionRequest{
		Parent:     "projects/groovy-iris-412518/locations/europe-west2",
		Function:   buildFunctionRequest(),
		FunctionId: "my-gcb-function-970601",
	}

	op, err := client.CreateFunction(ctx, req)
	fmt.Println("Creating CloudFunction...")
	if err != nil {
		panic(err)
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		panic(err)
	}
	// TODO: Use resp.
	_ = resp
}

func buildFunctionRequest() *functionspb.Function {
	functionReq := &functionspb.Function{
		Name:          "projects/groovy-iris-412518/locations/europe-west2/functions/my-gcb-function-970601",
		Description:   "my-function-description",
		BuildConfig:   createBuildConfig(),
		ServiceConfig: createServiceConfig(),
		Environment:   2, //gen2
	}

	return functionReq
}

func createBuildConfig() *functionspb.BuildConfig {
	buildConfigReq := &functionspb.BuildConfig{
		Runtime:              "python312",
		EntryPoint:           "hello_world",
		Source:               getFunctionSource(),
		EnvironmentVariables: getEnvironmentVariables(),
		//WorkerPool:           "workerPoolHere",
	}

	return buildConfigReq
}

func getFunctionSource() *functionspb.Source {
	bucketInfo := &functionspb.StorageSource{
		Bucket: "gcb-functions",
		Object: "hello_world_function.zip",
	}

	storageSource := &functionspb.Source_StorageSource{
		StorageSource: bucketInfo,
	}

	source := &functionspb.Source{
		Source: storageSource,
	}

	return source
}

func getEnvironmentVariables() map[string]string {
	envVars := make(map[string]string)
	envVars["TEST"] = "TEST123"

	return envVars
}

func createServiceConfig() *functionspb.ServiceConfig {
	serviceConfig := &functionspb.ServiceConfig{
		//VpcConnector:               "",
		AvailableMemory:  "256M",
		AvailableCpu:     "1",
		MaxInstanceCount: 1,
		MinInstanceCount: 0,
		//VpcConnectorEgressSettings: 2,
		//IngressSettings:            1,
		ServiceAccountEmail: "gcb-dev@groovy-iris-412518.iam.gserviceaccount.com",
	}

	return serviceConfig
}
