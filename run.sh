#!/bin/bash

goimports -w -d $(find . -type f -name '*.go' -not -path "./vendor/*")
go mod vendor