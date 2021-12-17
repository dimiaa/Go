FROM golang:latest

WORKDIR /weather/app
COPY . .

RUN go mod download
RUN go mod tidy

ENTRYPOINT go run .
