#!/bin/sh
x-terminal-emulator -e go run host/host.go -p 9091&
sleep 5
x-terminal-emulator -e go run host/host.go -p 9092&
x-terminal-emulator -e go run rproxy/rproxy.go&
x-terminal-emulator&




 


