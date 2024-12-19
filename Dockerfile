FROM golang:1.23.3-alpine3.20

COPY . /app

WORKDIR /app

RUN go mod download && CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o ./bin/shuttle ./cmd/

EXPOSE 9182
CMD ./bin/shuttle run