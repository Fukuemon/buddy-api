ARG GO_VERSION=1.22
FROM golang:${GO_VERSION} AS go
WORKDIR /app
COPY internal/go.mod internal/go.sum ./
RUN go mod download
COPY internal/ ./
RUN go build -o main ./cmd/api-server/main.go
EXPOSE 8080
CMD ["/app/main"]