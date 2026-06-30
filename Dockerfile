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
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum backend/Makefile ./
RUN make dep
COPY backend/ .
COPY --from=frontend-builder /app/backend/internal/webassets/dist /app/backend/internal/webassets/dist
RUN make build

# =========================
# STAGE 3: runtime
# =========================
FROM alpine:3.22
RUN addgroup -g 1000 -S appgroup && adduser -u 1000 -S appuser -G appgroup
RUN mkdir /app && chown appuser:appgroup /app
WORKDIR /app/backend
COPY --from=builder /app/backend/bin/keto-granola /app/keto-granola
USER appuser
ENTRYPOINT ["/app/keto-granola"]