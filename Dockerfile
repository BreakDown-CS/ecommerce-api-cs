# =========================================================
# 🏗  Stage 1: Build Stage
# ใช้ image golang สำหรับ compile source code ให้เป็น binary
# =========================================================
FROM golang:1.19 AS build-stage

WORKDIR /app

# Copy go module files ก่อน เพื่อใช้ Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy source code ทั้งหมด
COPY . .

# Build application binary จาก ./cmd
# ได้ไฟล์ชื่อ "app" อยู่ที่ /app/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd



# =========================================================
# 🚀 Stage 2: Release Stage (Production Image)
# ใช้ distroless image เพื่อความปลอดภัยและขนาดเล็ก
# =========================================================
FROM gcr.io/distroless/base-debian11

WORKDIR /app

# Copy เฉพาะ binary จาก stage แรก
COPY --from=build-stage /app/app .

EXPOSE 8080

# รันด้วย non-root user เพื่อความปลอดภัย
USER nonroot:nonroot

# เริ่มรันโปรแกรมเมื่อ container start
ENTRYPOINT ["./app"]