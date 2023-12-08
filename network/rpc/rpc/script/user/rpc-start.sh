#!/usr/bin/env bash

:<<BLOCK

BLOCK


function rpcStart() {
   cd /Users/xulei/jungle/golangworkspace/holy-go-lib/liblearn/rpc/server && \
   go run server.go
}

rpcStart