FROM golang:latest
WORKDIR /app/go
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /main
EXPOSE 8080
CMD ["/main"]