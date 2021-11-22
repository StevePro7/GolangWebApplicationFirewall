#!/bin/sh
#export LD_LIBRARY_PATH=/usr/local/modsecurity/lib/
sudo fuser -k 3080/tcp
go clean -cache
go run main.go