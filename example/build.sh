#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=arm go build main.go
scp main pi@192.168.199.103:~