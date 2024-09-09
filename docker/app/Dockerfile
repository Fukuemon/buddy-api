ARG GO_VERSION=1.22
FROM golang:${GO_VERSION} AS go
WORKDIR /app
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY internal/go.mod internal/go.sum ./
RUN go mod download
COPY internal/ ./
COPY internal/docs /app/docs
RUN swag init -g cmd/api-server/main.go
RUN go build -o main ./cmd/api-server/main.go
EXPOSE 8080
CMD ["/app/main"]