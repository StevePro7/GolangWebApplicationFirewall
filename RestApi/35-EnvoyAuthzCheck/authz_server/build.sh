#!/bin/sh
sudo fuser -k 8080/tcp
go clean -cache
go run main.go