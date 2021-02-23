FROM golang:latest AS builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/github.com/Kungfucoding23/Golang-Final-Proyect
COPY go.mod .
RUN go mod download
COPY . .
RUN go install

FROM scratch
COPY --from=builder /go/bin/github.com/Kungfucoding23/Golang-Final-Proyect .
ENTRYPOINT ["./main"]

#CMD ["./main"]
# docker build -t myapp . 