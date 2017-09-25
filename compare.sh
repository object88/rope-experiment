#!/bin/bash

# go test -bench . -benchtime 500ms -ver 1 > 1.txt
go test -bench . -benchtime 500ms -ver 2 > 2.txt
go test -bench . -benchtime 500ms -ver 3 > 3.txt

# $GOPATH/bin/benchcmp 1.txt 2.txt

$GOPATH/bin/benchcmp 2.txt 3.txt
