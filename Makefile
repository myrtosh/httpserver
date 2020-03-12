all:build run
build:
	docker build --no-cache -t httpservice .
run:
	docker run -p 8080:8080 httpservice