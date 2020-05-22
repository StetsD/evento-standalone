#!/usr/bin/env bash

protoc -I $GOPATH/src --go_out=plugins=grpc:$GOPATH/src $GOPATH/src/evento-standalone/internal/proto/service/event-service.proto
protoc -I $GOPATH/src --go_out=$GOPATH/src $GOPATH/src/evento-standalone/internal/proto/domain/event.proto