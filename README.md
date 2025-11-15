# Go NATS Demo

این پروژه یک دمو ساده برای کار با **NATS و JetStream** در Golang است.
در این پروژه دو سرویس داریم:

- **Publisher**: پیام‌ها را به JetStream ارسال می‌کند.
- **Worker**: پیام‌ها را دریافت و پردازش می‌کند.
- **NATS**: سرور پیام‌رسان با JetStream و داشبورد مدیریتی.

---

## ویژگی‌ها

- اتصال به NATS با retry (صبر می‌کند تا سرور بالا بیاید)
- Durable subscription در worker
- استفاده از JetStream برای ذخیره و پردازش پیام‌ها
- داشبورد مدیریتی NATS فعال بر روی پورت `8222`

---

## ساختار پروژه

```
go-nats-demo/
│
├── docker-compose.yml
├── publisher/
│   ├── Dockerfile
│   └── main.go
├── worker/
│   ├── Dockerfile
│   └── main.go
└── README.md
```

---

## راه‌اندازی پروژه

1. کلون کردن پروژه:

```bash
git clone git@github.com:h-agharezaei/go-nats-demo.git
cd go-nats-demo
```

2. اجرای Docker Compose:

```bash
docker compose up --build
```

3. دسترسی به داشبورد NATS:

```
http://localhost:8222/
```

---

## نکات

- پیام‌ها توسط publisher هر 2 ثانیه منتشر می‌شوند و worker آنها را دریافت می‌کند.
- اگر NATS هنوز بالا نیامده باشد، سرویس‌ها **منتظر اتصال می‌مانند**.
- Durable subscription باعث می‌شود worker بعد از ریستارت هم پیام‌های از دست رفته را دریافت کند.

---

## نیازمندی‌ها

- Docker 20+
- Docker Compose 1.29+
- Golang 1.25+ (برای ساخت Docker image ها)