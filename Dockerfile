# syntax=docker/dockerfile:1

FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
ADD internal/. internal/
ADD migrations/. migrations/ 
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/todo-list
EXPOSE 8000
CMD ["/app/todo-list"]