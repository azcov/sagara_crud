# Build Stage
FROM golang:1.16.5-alpine3.13 AS builder
WORKDIR /docker/app
COPY . .

COPY go.mod ./ 
COPY go.sum ./
COPY app_config.json ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8081
CMD ["/app/main"]
