# =========================================================
# 🏗  Stage 1: Build Stage
# =========================================================
FROM golang:1.19 AS build-stage

WORKDIR /app

# รับค่า VERSION จาก docker build
ARG VERSION=dev

# Copy go module files ก่อน เพื่อใช้ Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy source code ทั้งหมด
COPY . .

# Build binary พร้อมฝัง version เข้าไปในตัวแอป
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-X main.version=${VERSION}" \
    -o app ./cmd



# =========================================================
# 🚀 Stage 2: Release Stage
# =========================================================
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=build-stage /app/app .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./app"]