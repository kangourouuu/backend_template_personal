🎉 Code Generator
Một công cụ tự động sinh mã nguồn dựa trên Go, cung cấp các API để tạo dự án và mã CRUD cho các thực thể, với mã được sinh ra tại ổ C:/.
Mục lục

Giới thiệu
Tính năng
Yêu cầu
Cài đặt
Cách sử dụng
API
Cấu trúc thư mục

Giới thiệu
Dự án Code Generator là một công cụ giúp tự động hóa việc tạo mã nguồn cho các dự án Go. Nó cung cấp hai API để sinh cấu trúc dự án và mã CRUD dựa trên các tham số đầu vào. Mã được sinh ra sẽ được lưu tại ổ C:/ với thư mục chính là tên dự án (C:/ProjectName/). Dự án này phù hợp cho các nhà phát triển muốn tiết kiệm thời gian khi khởi tạo dự án hoặc tạo các đoạn mã CRUD cơ bản.
Tính năng

✅ Tự động tạo cấu trúc thư mục cho dự án Go.
✅ Sinh mã CRUD cho các thực thể với các trường tùy chỉnh.
✅ API đơn giản, dễ sử dụng với đầu vào linh hoạt.
✅ Lưu mã nguồn tại ổ C:/ với cấu trúc rõ ràng.

Yêu cầu

Go 1.16 hoặc cao hơn
Hệ điều hành: Windows (do mã được sinh ra tại ổ C:/)
Công cụ gọi API (ví dụ: Postman, cURL, hoặc bất kỳ HTTP client nào)

Cài đặt

Clone repository:git clone https://github.com/kangourouuu/backend_template_personal.git


Di chuyển vào thư mục dự án:cd backend_template_personal


Cài đặt các gói phụ thuộc:go mod tidy

Chạy ứng dụng:go run main.go

Ứng dụng sẽ chạy tại http://localhost:9000.

Cách sử dụng
Sau khi chạy ứng dụng, bạn có thể sử dụng hai API để tạo dự án hoặc mã CRUD. Mã được sinh ra sẽ được lưu tại C:/ProjectName/.
Lưu ý

Đảm bảo ổ C:/ có quyền ghi để lưu trữ mã được sinh ra.
Sử dụng các công cụ như Postman hoặc cURL để gửi yêu cầu đến API.
Kiểm tra định dạng JSON của các tham số đầu vào để tránh lỗi.

API
Dự án cung cấp hai API chính:
1. Tạo dự án

Endpoint: POST http://localhost:9000/api/v1/generate-project
Mô tả: Tạo một dự án Go mới với cấu trúc thư mục tại C:/ProjectName/.
Input (JSON):{
  "ProjectName": "MyNewProject",
  "Port": "8080"
}


Output: Thư mục dự án được tạo tại C:/MyNewProject/.

2. Tạo mã CRUD

Endpoint: POST http://localhost:9000/api/v1/generate-crud
Mô tả: Tạo mã CRUD cho một thực thể với các trường tùy chỉnh.
Input (JSON):{
  "EntityName": "User",
  "EntityNameLower": "user",
  "Fields": [
    {
      "Name": "ID",
      "Type": "int"
    },
    {
      "Name": "Name",
      "Type": "string"
    }
  ]
}


Output: Mã CRUD được sinh ra trong thư mục dự án tại C:/ProjectName/.

Cấu trúc thư mục
Cấu trúc mã nguồn được sinh ra tại C:/ProjectName/ sẽ bao gồm:
C:/ProjectName/
├── main.go             # File chính để chạy ứng dụng
├── api/                # Chứa mã liên quan đến API
│   └── v2/             # Phân chia các phiên bản API
├── build/              # Chứa các file build hoặc script build
├── common/             # Code chung
│   ├── api_response/   # Xử lý phản hồi API
│   ├── err_response/   # Xử lý lỗi trả về
│   ├── log/            # Logic logging
│   └── limiter/        # Cấu hình giới hạn
├── configs/            # Cấu hình ứng dụng
├── constant/           # Hằng số
├── internal/           # Code nội bộ
│   ├── sqlclient/      # Cấu hình kết nối SQL
│   └── redis/          # Logic liên quan đến Redis
├── middleware/         # Middleware
├── model/              # Định nghĩa các model/thực thể
├── repository/         # Logic truy cập dữ liệu
├── server/             # Logic server
│   └── http/           # Repository HTTP
├── service/            # Logic nghiệp vụ
├── tmp/                # Thư mục tạm
├── dto/                # Data Transfer Objects
└── docs/               # Tài liệu dự án
