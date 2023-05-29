FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app server.go \
    && chmod +x app


FROM alpine
WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT ["./app"]
