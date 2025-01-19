# مرحله اول: Build (کامپایل و ساخت)
FROM golang:1.20 AS builder

WORKDIR /app

# نصب Swagger CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

# کپی فایل‌های پروژه و وابستگی‌ها
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# تولید مستندات Swagger
RUN swag init

# ساخت فایل اجرایی
RUN go build -ldflags="-s -w" -o main cmd/main.go

# مرحله دوم: Runtime (اجرای نهایی)
FROM gcr.io/distroless/base-debian11:nonroot

WORKDIR /root/

# کپی فایل اجرایی و مستندات Swagger
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# پورت مورد نیاز را باز کنید
EXPOSE 8001

# اجرای فایل اجرایی
CMD ["./main"]
