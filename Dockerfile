FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build /app/cmd/main.go

EXPOSE 8080

ENV GIN_MODE=release

CMD ["/app/main"]