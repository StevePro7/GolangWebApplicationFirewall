#!/bin/sh
sudo fuser -k 3080/tcp
go clean -cache
go run main.go