FROM golang:1.17-alpine as builder

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ./bin .

FROM alpine:latest
COPY --from=builder /app/bin /app/bin

RUN mkdir -p /app/uploads

EXPOSE 3000

WORKDIR /app

ENTRYPOINT ["/app/bin"]
