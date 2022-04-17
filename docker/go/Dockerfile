FROM golang:latest
COPY . /app/backend
WORKDIR /app/backend
RUN go build -o main ./cmd/main.go 
CMD ["/app/backend/main"]