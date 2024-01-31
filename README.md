# google-cloud-builder

A cli program built in Go to provision infrastructure in a GCP project.

Current capability:
* Create a gen2 Cloud Function - source = GCS bucket
* Update a gen2 Cloud Function - source = GCS bucket


# config

```gcp:
  projectId: "groovy-iris-412518"
  location: "europe-west2"
function:
  - functionName: "my-test-function"
    functionDescription: "my-function-description1"
    runtime: "python312"
    entrypoint: "hello_world"
    bucketName: "gcb-functions"
    sourceFile: "hello_world_function"
    availableCpu: 1
    availableMemory: "256M"
    minInstanceCount: 0
    maxInstanceCount: 1
    serviceAccountEmail: "gcb-dev@groovy-iris-412518.iam.gserviceaccount.com"

  - functionName: "my-test-function2"
    functionDescription: "my-function-description2"
    runtime: "python312"
    entrypoint: "hello_world"
    bucketName: "gcb-functions"
    sourceFile: "hello_world_function"
    availableCpu: 1
    availableMemory: "256M"
    minInstanceCount: 0
    maxInstanceCount: 1
    serviceAccountEmail: "gcb-dev@groovy-iris-412518.iam.gserviceaccount.com" ```

    

# todo

* Use viper, get function from config and use any env vars
* list functions first, check if it exists
* if it exists, update, if not create
* return useful info in command response
* Learn how to write tests in go
* delete function capability



