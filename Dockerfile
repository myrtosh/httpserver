FROM golang:1.18

WORKDIR /go/src/github.com/myrtosh/httpservice
COPY .git     .
COPY httpservice.go   .
COPY go.mod    .

RUN GIT_COMMIT=$(git rev-list -1 HEAD) && \
  go build -ldflags "-X main.GitCommit=$GIT_COMMIT"

RUN go build -o /httpservice

EXPOSE 8080

CMD ["./httpservice"]