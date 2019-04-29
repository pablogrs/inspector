#/bin/bash

protoc --go_out=plugins=grpc:. commands/commands.proto
protoc -I . --go_out=plugins=grpc:. results/*.proto
