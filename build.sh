#!/bin/bash
set -xe
GOPATH=`pwd` go test -v -c sse42
GOPATH=`pwd` go test -v sse42 -bench=.
