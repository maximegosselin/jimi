#!/bin/sh

GOOS=js GOARCH=wasm go build -o jimi.wasm ../.
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .