# ==========================================
# 1. BUILD STAGE
# ==========================================
FROM golang:1.25.0-alpine3.22 AS builder

WORKDIR /http

ARG GIT_VERSION

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binary
RUN current_time=$(date +"%Y-%m-%dT%H:%M:%SZ") && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.buildTime=$current_time -X main.version=$GIT_VERSION" -o /out/main ./

# ==========================================
# 2. RUNTIME STAGE
# ==========================================
FROM alpine:3.22 

RUN apk --no-cache add curl tzdata

ENV TZ=Asia/Ho_Chi_Minh

WORKDIR /http

RUN addgroup -S appgroup && adduser -S appuser -G appgroup \
    && chown -R appuser:appgroup /http

COPY --from=builder /out/main /http/main
COPY entry-point.sh /http/entry-point.sh
RUN chmod +x /http/entry-point.sh && chown appuser:appgroup /http/entry-point.sh

USER appuser

EXPOSE 8080

ENTRYPOINT ["/http/entry-point.sh"]