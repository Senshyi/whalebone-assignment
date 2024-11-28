# ============================================
# BUILDER
# ============================================
FROM golang:1.23-bookworm AS base

FROM base AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o whalebone-assignment cmd/api/main.go

# ============================================
# PRODUCTION
# ============================================
FROM golang:1.23-bookworm AS production

WORKDIR /prod

COPY --from=builder /build/whalebone-assignment ./

EXPOSE 8000

CMD ["./whalebone-assignment"]
