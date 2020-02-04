#!/bin/sh
x-terminal-emulator -e go run host/host.go 9091&
x-terminal-emulator -e go run host/host.go 9092&
x-terminal-emulator -e go run rproxy/rproxy.go&
x-terminal-emulator&




 


