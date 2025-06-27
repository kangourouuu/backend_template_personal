# 🎉 Code Generator

**Một công cụ tự động sinh mã nguồn Go với API tiện dụng — giúp bạn tạo nhanh dự án và mã CRUD, lưu trữ tại ổ đĩa `C:/`.**

---

## 📚 Mục lục

* [🔰 Giới thiệu](#-giới-thiệu)
* [✨ Tính năng](#-tính-năng)
* [⚙️ Yêu cầu](#️-yêu-cầu)
* [📦 Cài đặt](#-cài-đặt)
* [🚀 Cách sử dụng](#-cách-sử-dụng)
* [🧩 API](#-api)

  * [1. Tạo dự án](#1-tạo-dự-án)
  * [2. Tạo mã CRUD](#2-tạo-mã-crud)
* [📁 Cấu trúc thư mục sau khi sinh ra](#-cấu-trúc-thư-mục-sau-khi-sinh-ra)

---

## 🔰 Giới thiệu

**Code Generator** là một công cụ giúp tự động hóa việc tạo mã nguồn cho các dự án Go.
Bạn chỉ cần gọi API để sinh:

* Cấu trúc dự án Go chuẩn
* Mã CRUD cho các thực thể (entity)

📍 **Mã sẽ được sinh tại:** `C:/ProjectName/`

Công cụ lý tưởng cho lập trình viên muốn tiết kiệm thời gian khởi tạo và phát triển tính năng cơ bản.

---

## ✨ Tính năng

* ✅ Tự động tạo cấu trúc thư mục cho dự án Go
* ✅ Sinh mã CRUD với trường tùy chỉnh
* ✅ API đơn giản, dễ tích hợp
* ✅ Lưu mã tại `C:/` với cấu trúc rõ ràng

---

## ⚙️ Yêu cầu

* Go **1.16** trở lên
* Hệ điều hành: **Windows** (mã được sinh tại `C:/`)
* Công cụ gửi request: Postman, `cURL`, Insomnia, v.v.

---

## 📦 Cài đặt

```bash
# Clone repository
git clone https://github.com/kangourouuu/backend_template_personal.git

# Di chuyển vào thư mục
cd backend_template_personal

# Cài đặt các gói phụ thuộc
go mod tidy

# Chạy ứng dụng
go run main.go
```

Ứng dụng sẽ chạy tại: [http://localhost:9000](http://localhost:9000)

---

## 🚀 Cách sử dụng

Sau khi ứng dụng khởi chạy, bạn có thể:

1. Tạo dự án Go mới
2. Sinh mã CRUD cho một thực thể

📌 **Lưu ý:**

* Đảm bảo ổ `C:/` có quyền ghi
* Sử dụng định dạng JSON hợp lệ trong các request
* Gửi request qua Postman hoặc bất kỳ HTTP client nào

---

## 🧩 API

### 1. Tạo dự án

* **Endpoint:** `POST /api/v1/generate-project`
* **Mô tả:** Tạo một dự án mới với cấu trúc thư mục đầy đủ tại `C:/ProjectName/`

#### 📨 Request (JSON)

```json
{
  "project_name": "MyNewProject",
  "port": "8080"
}
```

📂 **Kết quả:** Thư mục `C:/MyNewProject/` được tạo

---

### 2. Tạo mã CRUD

* **Endpoint:** `POST /api/v1/generate-crud`
* **Mô tả:** Sinh mã CRUD cho một thực thể với các trường tự định nghĩa

#### 📨 Request (JSON)

```json
{
  "entity_name": "User",
  "entity_name_lower": "user",
  "fields": [
    {
      "name": "ID",
      "type": "int"
    },
    {
      "name": "Name",
      "type": "string"
    }
  ]
}
```

📂 **Kết quả:** Mã CRUD được tạo trong thư mục `C:/MyNewProject/`

---

## 📁 Cấu trúc thư mục sau khi sinh ra

```
C:/ProjectName/
├── main.go             # File chính khởi chạy ứng dụng
├── Dockerfile             # Build container Docker
├── docker-compose.yaml    # Build service 
├── api/
│   └── v2/             # API versioning
├── build/              # Script build hoặc pipeline
├── common/             # Logic dùng chung
│   ├── api_response/   # Phản hồi chuẩn API
│   ├── err_response/   # Xử lý lỗi
│   ├── log/            # Ghi log
│   └── limiter/        # Cấu hình giới hạn request
├── configs/            # Cấu hình ứng dụng
├── constant/           # Các hằng số dùng chung
├── internal/           
│   ├── sqlclient/      # Kết nối cơ sở dữ liệu
│   └── redis/          # Logic Redis
├── middleware/         # Middleware xử lý request
├── model/              # Định nghĩa các entity/model
├── repository/         # Truy cập dữ liệu
├── server/             
│   └── http/           # Khởi tạo HTTP server
├── service/            # Business logic
├── tmp/                # Tạm thời (temporary files)
├── dto/                # Data Transfer Object
└── docs/               # Tài liệu dự án
```

