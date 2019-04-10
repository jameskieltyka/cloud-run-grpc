# CLOUD RUN GRPC
## Overview

This project is to test grpc on cloud run.

Micro Services:
- Entry-GRPC
- Internal-GRPC

The entry-GRPC service is to act as the entry point into the knative environment and will be used to invoke a secondary GRPC service (internal-GRPC) to see if service to service calls work on GCP Cloud Run.

The client will be used to test the deployed endpoint

## Requirements
- Go 1.11+
- Protobuf tooling (protoc, etc)
- Go specific GRPC tools: https://grpc.io/docs/quickstart/go.html

- Containers to be build with google cloud build and stored on GCR