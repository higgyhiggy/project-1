

include .env
hello:
	echo "makefile running"
	echo $(portone)
net:
	echo "makefile running script to run both host and a rproxy server"
	./yes.sh
docker:
	docker run -p 9091:8000 --rm -d host
	docker run -p 9092:8000 --rm -d host
	docker run -p 9093:8000 --rm rproxy