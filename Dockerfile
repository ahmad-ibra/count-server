FROM golang:1.21

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/main server/main.go
EXPOSE $PORT

CMD ["/app/main"]
