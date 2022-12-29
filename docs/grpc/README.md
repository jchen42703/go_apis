# GRPC <!-- omit in toc -->

## Table of Contents <!-- omit in toc -->

- [Intro to gRPC](#intro-to-grpc)
  - [Workflow](#workflow)
  - [Example Protocol Buffer](#example-protocol-buffer)
- [gRPC Core Concepts](#grpc-core-concepts)
  - [1. Service Definition](#1-service-definition)
  - [2. Code Generation](#2-code-generation)
  - [3. Implementation + Big Picture](#3-implementation--big-picture)
- [gRPC Quickstart](#grpc-quickstart)

# Intro to gRPC

https://www.ionos.com/digitalguide/server/know-how/an-introduction-to-grpc/

## Workflow

1. **Define service contract.** Client and server define the protocol buffers (contract).

   - The client stub must find a match from the server to be able to send a request.

2. **Generate gRPC code from the proto file.**

   - Typically use something like `protoc` to generate gRPC code from a `proto` file for a target language

3. **Implement server.**

4. **Implement client stub that calls the service contract.**

## Example Protocol Buffer

Search for product in inventory.

```protobuf
syntax = "proto3";
package gRPC_service;
import "google/protobuf/wrappers.proto";
// Interface we want to call and implement
service InventoryService {
	rpc getItemByName(google.protobuf.StringValue) returns (Items);
	rpc getItemByID(google.protobuf.StringValue) returns (Item);
	 rpc addItem(Item) returns (google.protobuf.BoolValue);
}

message Items {
  string itemDesc = 1;
  repeated Item items = 2;
}

message Item {
	string id = 1;
	string name = 2;
	string description = 3;
}
```

---

# gRPC Core Concepts

https://grpc.io/docs/what-is-grpc/core-concepts/

## 1. Service Definition

In a `proto` file:

```protobuf
service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string greeting = 1;
}

message HelloResponse {
  string reply = 1;
}
```

1. Specifies methods that can be called remotely, their parameters and return types
2. Describes service interface and payload structure

## 2. Code Generation

Using the `proto` file, you can use `protoc` to generate gRPC code.

- Clients will typically call this code, but servers need to actually implement the API.
  - AKA the servers implement what the called service methods actually do (logging, calling ML model, calling DB etc.)

## 3. Implementation + Big Picture

1. **Server**
   1. implements methods defined by service.
   2. runs gRPC server to handle client-side calls
   3. Decodes incoming requests, executes service methods, and encodes service responses
2. **Client**
   1. Has a **stub** (implements the same methods as the service)
   2. The stub abstracts the RPC process, so that clients can call the service as if the server doesn't exist.
      1. Each method in the `stub` wraps the parameters into a protobuf, sends the request to the server, and returns the server's protobuf response.

# gRPC Quickstart

https://grpc.io/docs/languages/go/quickstart/
