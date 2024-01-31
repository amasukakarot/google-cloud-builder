# google-cloud-builder

A cli program built using **Cobra**, **Viper** and the **Google Go Cloud Client Library** to create resources in a GCP project. 

### Features
* Create **N** number of **Gen2 Cloud Functions** using a GCS bucket as the source of the function.
* Update **N** number of **Gen2 Cloud Functions** using a GCS bucket as the source of the function.


### Commands

```google-cloud-builder deploy --config=values.yaml```

### config

```
gcp:
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
    serviceAccountEmail: "example@project.iam.gserviceaccount.com"

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
    serviceAccountEmail: "example@project.iam.gserviceaccount.com"
```

    

# todo

* Dockerfile
* Error handling  / clean up



