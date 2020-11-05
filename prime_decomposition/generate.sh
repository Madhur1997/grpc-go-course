#!/bin/bash

protoc primedecompositionpb/prime_decomposition.proto --go_out=plugins=grpc:.
