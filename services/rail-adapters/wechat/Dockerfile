# syntax=docker/dockerfile:1

# ── Build stage ────────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o /service .

# ── Runtime stage ──────────────────────────────────────────────
FROM scratch

COPY --from=builder /service /service

EXPOSE 8094
ENV PORT=8094

ENTRYPOINT ["/service"]
