# HNSECRETS

## Development Prereqs

* Go: The language this app is written in  - https://golang.org/dl
    * Git is a requirement of Go (to use `go get`)
* Protoc: The Protocol Buffers compiler - https://github.com/protocolbuffers/protobuf/releases/latest
* make: The build tool to build this app - TODO
* docker/moby-engine/podman/etc...: A container runtime - TODO

## Building

To build this project, perform the following:

**Building/Running Locally**

1. clone this repository
2. `cd hnsecrets`
3. `make && make run`

This set of commands will test & build the binaries locally (in the respetive `svcs/<service>/bin` directory) and then runs the binaries for local testing.

**Building/Deploying Containers**

1. clone this repository
2. `cd hnsecrets`
3. `docker login`
4. `make test && make build && make deploy`

This set of commands will test the code, build the containers, and deploy the containers to the `j4ng5y/hnsecrets` dockerhub repository.

**Checking test coverage**

To check the current test coverage, simply run `make test && make coverage`. This will open the HTML test coverage reports for both services.
