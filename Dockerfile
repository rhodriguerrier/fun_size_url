FROM golang:1.16.4

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./docker_basics_app .

ENV PORT=:8080

EXPOSE 8080

CMD ["./docker_basics_app"]
