#!/bin/sh
export LD_LIBRARY_PATH=/usr/local/modsecurity/lib/
go clean -cache
go run main.go