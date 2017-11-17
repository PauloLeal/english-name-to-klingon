#!/usr/bin/env bash

cover () {
    t="/tmp/go-cover.$$.tmp"
    testOut="/tmp/go-test-$$.test.out"
    go test -coverprofile=$t $@
    go tool cover -html=$t
    unlink $t
}

cover github.com/PauloLeal/english-name-to-klingon/stapiclient
cover github.com/PauloLeal/english-name-to-klingon/

