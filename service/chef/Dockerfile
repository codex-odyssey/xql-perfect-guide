FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY --link go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY --link . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o chef

FROM alpine:3.18
RUN apk --no-cache add ca-certificates
WORKDIR /wd/
COPY --link --from=builder /app/chef .
CMD ["./chef"]
