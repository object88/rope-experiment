#!/bin/bash

go test -bench . -benchtime 1s -ver 1 > 1.txt
go test -bench . -benchtime 1s -ver 2 > 2.txt

$GOPATH/bin/benchcmp 1.txt 2.txt
