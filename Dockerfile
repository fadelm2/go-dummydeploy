# Stage 1: Build Go binary
FROM golang:latest AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Stage 2: Run minimal image
FROM alpine:latest

WORKDIR /root/

# copy binary
COPY --from=builder /app/app .

# 🔥 copy config.json
COPY config.json .

EXPOSE 8086
CMD ["./app"]
