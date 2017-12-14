#!/bin/bash

protoc3 -I . ./test.proto --go_out=plugins=grpc:.
