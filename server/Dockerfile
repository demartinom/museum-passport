ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
# Build the binary
RUN go build -v -o /run-app .

# Final Stage
FROM debian:bookworm-slim

# Install ca-certificates AND curl
RUN apt-get update && apt-get install -y ca-certificates curl && rm -rf /var/lib/apt/lists/*

COPY --from=builder /run-app /usr/local/bin/

RUN chmod +x /usr/local/bin/run-app

CMD ["run-app"]