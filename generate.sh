#!/bin/bash

protoc resizeimagemspb/resizeimagems.proto --go_out=plugins=grpc:.

# Node.js Client
# protoc calculator/calculatorpb/calculator.proto --js_out=library=calculator/calculator_client_node,binary:calculator/calculator_client_node/build/gen

# Development Docker
docker build -f development.Dockerfile -t resizeimage_go_grpc_server:v1 .