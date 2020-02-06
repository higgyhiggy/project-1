all: hello run
include .env

hello:
	echo "makefile running"
	echo $(portone)
run:
	echo "makefile running script to run both host and a rproxy server"
	./yes.sh