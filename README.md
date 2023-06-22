# Restful API Go

### Điều kiện cần:

- Mysql
- Go
- Makefile

### Hướng dẫn khởi chạy ứng dụng trên localhost

Thay đổi file .env theo đúng định dạng, cấu hình để kết nối tới mysql local

Cài đặt dependencies:

- air

```
go install github.com/cosmtrek/air@latest
```

**Ubuntu**

Build command:

```
make build
```

**Windows**

Build command:

```
make win-build
```

### Cách thức hoạt động API

[API Docs](https://app.swaggerhub.com/apis-docs/batthanhvan/apiservice/2.0)