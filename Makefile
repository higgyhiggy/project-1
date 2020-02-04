all: hello run
include .env

hello:
	echo "hello"
	echo $(portone)
run:
	echo " this is run"
	./yes.sh