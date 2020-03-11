 
FROM golang:1.14 as httpservice

ENV GOPATH /home/myrto/go
ENV PATH ${GOPATH}/bin:$PATH

WORKDIR /home/myrto/endocode/go_training/httpserver
COPY . .

RUN go get /home/myrto/go/src/github.com/fatih/camelcase

RUN GIT_COMMIT=$(git rev-list -1 HEAD) && \
  go build -ldflags "-X main.GitCommit=$GIT_COMMIT" httpservice.go

EXPOSE 8080

CMD ["./httpservice"]