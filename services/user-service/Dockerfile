# مرحله اول: Build (کامپایل و ساخت)
FROM golang:1.23.5-alpine3.21 AS builder

WORKDIR /app

# نصب Swagger CLI
# RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
# کپی فایل‌های پروژه و وابستگی‌ها
COPY go.mod go.sum ./
RUN go mod download

COPY . .


# RUN swag init --dir ./internal --output ./docs


RUN swag init --dir ./internal --output ./docs

# ساخت فایل اجرایی

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go
# مرحله دوم: Runtime (اجرای نهایی)

FROM alpine:latest

WORKDIR /app/

# کپی فایل اجرایی و مستندات Swagger
COPY --from=builder /app/main .
# COPY --from=builder /app/docs ./docs

# پورت مورد نیاز را باز کنید
EXPOSE 8001

# اجرای فایل اجرایی
CMD ["./main"]
