FROM golang:alpine3.20 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build  /app/cmd/main.go

EXPOSE 8080

ENV GIN_MODE=release

CMD ["/app/main"]

FROM alpine:latest

WORKDIR /app

ENV GIN_MODE=release

COPY --from=build /app/main /app/main

COPY ./index/ ./index/

EXPOSE 8080

ENTRYPOINT ["/app/main"]