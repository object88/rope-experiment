#!/bin/bash

go test -bench . -benchtime 500ms -ver 1 > 1.txt
go test -bench . -benchtime 500ms -ver 2 > 2.txt
go test -bench . -benchtime 500ms -ver 3 > 3.txt

echo "Comparing v1 and v2"
$GOPATH/bin/benchcmp 1.txt 2.txt

echo "Comparing v1 and v3"
$GOPATH/bin/benchcmp 1.txt 3.txt
