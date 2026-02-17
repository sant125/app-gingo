# ── Build stage ───────────────────────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install swag CLI for doc generation
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Generate Swagger docs
RUN swag init -g cmd/main.go -o docs

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o gin-tattoo ./cmd/main.go

# ── Runtime stage ──────────────────────────────────────────────────────────────
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/gin-tattoo .
COPY --from=builder /app/static ./static

EXPOSE 8080

ENTRYPOINT ["./gin-tattoo"]

# ─────────────────────────────────────────────────────────────────────────────
# ArgoCD / Kubernetes — manifests will be applied manually via ArgoCD.
# See k8s/ folder for commented deployment and service manifests.
# ─────────────────────────────────────────────────────────────────────────────
