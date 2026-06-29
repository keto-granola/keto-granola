# ====================================
# Stage 1: Build frontend assets
# ====================================
FROM node:24-alpine AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci
COPY frontend/ ./
RUN npm run build

# =========================
# Stage 2: Build the binary
# =========================
FROM golang:1.26 AS builder
WORKDIR /app
COPY go.mod go.sum Makefile ./
RUN make dep
COPY . .
COPY --from=frontend-builder /app/frontend/dist /app/internal/webassets/dist
RUN make build

# =========================
# STAGE 3: runtime
# =========================
FROM alpine:3.22
RUN addgroup -g 1000 -S appgroup && adduser -u 1000 -S appuser -G appgroup
RUN mkdir /app && chown appuser:appgroup /app
WORKDIR /app
COPY --from=builder /app/keto-granola-server /app/keto-granola-server
USER appuser
ENTRYPOINT ["/app/keto-granola-server"]