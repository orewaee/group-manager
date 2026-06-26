FROM golang:1.26.4-alpine3.24 AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /build

COPY . .

RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o /build/api ./cmd/api

FROM alpine:3.24

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /build/api /api

EXPOSE 50000

CMD ["/api"]
