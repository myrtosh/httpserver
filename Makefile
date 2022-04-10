build:
	go build -mod vendor -o out/bin/httpservice httpservice.go

run:
	go run httpservice.go

clean:
	rm -fr ./bin
	rm -fr ./out

docker-build:
	sudo docker build -t httpservice .

docker-run:
	sudo docker run -p 8080:8080 -t httpservice
