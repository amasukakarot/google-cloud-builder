package cloudfunction

import (
	"cloud.google.com/go/functions/apiv2/functionspb"
	"context"
	"fmt"
	"github.com/amasukakarot/google-cloud-builder/internal/config"
	"google.golang.org/api/iterator"
	"log"
	"strings"
	"time"
)

func StartDeployment(ctx context.Context) {
	//loop through each function from config
	//check if function exists
	//if function exists, update
	//if function doeesn't exist, create

	projectId := config.FunctionData.Project.GCPProjectId
	location := config.FunctionData.Project.Location
	for _, function := range config.FunctionData.Function {
		if IfFunctionExists(ctx, function, projectId, location) {
			log.Println("Updating cloud function...")
			UpdateCloudFunction(ctx, function, projectId, location)
		} else {
			log.Println("Creating cloud function...")
			start := time.Now()
			CreateCloudFunction(ctx, function, projectId, location)
			finished := time.Since(start)
			log.Printf("Function took %s", finished)
		}
	}
}

func CreateCloudFunction(ctx context.Context, function config.Function, projectId string, location string) {

	client := createClient(ctx)
	defer client.Close()

	req := &functionspb.CreateFunctionRequest{
		Parent:     fmt.Sprintf("projects/%v/locations/%v", projectId, location),
		Function:   buildFunctionRequest(function, projectId, location),
		FunctionId: function.FunctionName,
	}

	op, err := client.CreateFunction(ctx, req)
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

func buildFunctionRequest(function config.Function, projectId string, location string) *functionspb.Function {

	functionName := function.FunctionName
	functionDescription := function.FunctionDescription

	functionReq := &functionspb.Function{
		Name:          fmt.Sprintf("projects/%v/locations/%v/functions/%v", projectId, location, functionName),
		Description:   functionDescription,
		BuildConfig:   createBuildConfig(function),
		ServiceConfig: createServiceConfig(function),
		Environment:   2, //gen2
	}

	return functionReq
}

func createBuildConfig(function config.Function) *functionspb.BuildConfig {
	runtime := function.Runtime
	entrypoint := function.Entrypoint
	buildConfigReq := &functionspb.BuildConfig{
		Runtime:              runtime,
		EntryPoint:           entrypoint,
		Source:               getFunctionSource(function.SourceFile, function.Bucket),
		EnvironmentVariables: getEnvironmentVariables(),
		//WorkerPool:           "workerPoolHere",
	}

	return buildConfigReq
}

func getFunctionSource(sourceCode string, bucket string) *functionspb.Source {

	bucketInfo := &functionspb.StorageSource{
		Bucket: bucket,
		Object: fmt.Sprintf("%v.zip", sourceCode),
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

func createServiceConfig(function config.Function) *functionspb.ServiceConfig {
	serviceConfig := &functionspb.ServiceConfig{
		//VpcConnector:               "",
		AvailableMemory:  function.AvailableMemory,
		AvailableCpu:     function.AvailableCpu,
		MaxInstanceCount: function.MaxInstanceCount,
		MinInstanceCount: function.MinInstanceCount,
		//VpcConnectorEgressSettings: 2,
		//IngressSettings:            1,
		ServiceAccountEmail: function.ServiceAccountEmail,
	}

	return serviceConfig
}

func UpdateCloudFunction(ctx context.Context, function config.Function, projectId string, location string) {

	functionName := fmt.Sprintf("projects/%v/locations/%v/functions/%v", projectId, location, function.FunctionName)
	log.Printf("Updating %v ", functionName)

	client := createClient(ctx)
	defer client.Close()

	req := &functionspb.UpdateFunctionRequest{
		Function: buildFunctionRequest(function, projectId, location),
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

func IfFunctionExists(ctx context.Context, function config.Function, projectId string, location string) bool {

	client := createClient(ctx)
	defer client.Close()

	functionName := function.FunctionName
	fmt.Println(projectId)
	fmt.Println(location)
	req := &functionspb.ListFunctionsRequest{
		Parent: fmt.Sprintf("projects/%v/locations/%v", projectId, location),
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
		if strings.EqualFold(resp.Name, fmt.Sprintf("projects/%v/locations/%v/functions/%v", projectId, location, functionName)) {
			log.Printf("Function %v found", function)
			return true
		}
	}
	log.Printf("Function %v not found.", functionName)
	return false
}
