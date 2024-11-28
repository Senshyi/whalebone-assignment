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
# preferably something like `scratch` should be used but because
# CSO_ENABLE needs to be set to 1 for sqLite it's not possible.
# For the sake of this demo app, this should do but normally it could
# and should be size optimized 
FROM golang:1.23-bookworm AS runner

WORKDIR /app

COPY --from=builder /build/whalebone-assignment ./

EXPOSE 8000

CMD ["./whalebone-assignment"]
