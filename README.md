🎉 Code Generator

An automated Go source code generator with handy APIs — helping you quickly create projects and CRUD code, saved directly to the C:/ drive.

📚 Table of Contents
	•	🔰 Introduction
	•	✨ Features
	•	⚙️ Requirements
	•	📦 Installation
	•	🚀 Usage
	•	🧩 API
	•	1. Generate Project
	•	2. Generate CRUD Code
	•	📁 Generated Folder Structure

🔰 Introduction

Code Generator is a tool that automates the creation of Go source code projects.
You only need to call APIs to generate:
	•	A standard Go project structure
	•	CRUD code for your entities

📍 Generated code will be saved at: C:/ProjectName/

This tool is ideal for developers looking to save time when setting up new projects or implementing basic CRUD logic.

✨ Features
	•	✅ Automatically generate a Go project folder structure
	•	✅ Generate CRUD code with customizable fields
	•	✅ Simple, flexible API
	•	✅ Source code saved in C:/ with clear structure

⚙️ Requirements
	•	Go 1.16 or later
	•	Operating system: Windows (code is generated at C:/)
	•	HTTP client tools: Postman, cURL, Insomnia, etc.

📦 Installation

# Clone the repository
git clone https://github.com/kangourouuu/backend_template_personal.git

# Navigate into the project directory
cd backend_template_personal

# Install dependencies
go mod tidy

# Run the application
go run main.go

The application will run at: http://localhost:9000

🚀 Usage

After starting the application, you can:
	1.	Generate a new Go project
	2.	Generate CRUD code for an entity

📌 Note:
	•	Make sure C:/ has write permission
	•	Use valid JSON format in requests
	•	Send requests via Postman or any HTTP client

🧩 API

1. Generate Project
	•	Endpoint: POST /api/v1/generate-project
	•	Description: Create a new Go project with full folder structure at C:/ProjectName/

📨 Request (JSON)

{
  "project_name": "MyNewProject",
  "port": "8080"
}

📂 Result: A project folder will be created at C:/MyNewProject/

2. Generate CRUD Code
	•	Endpoint: POST /api/v1/generate-crud
	•	Description: Generate CRUD code for an entity with custom fields

📨 Request (JSON)

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

📂 Result: CRUD code will be generated inside C:/MyNewProject/

📁 Generated Folder Structure

C:/ProjectName/
├── main.go                 # Entry point of the application
├── Dockerfile              # Docker container build file
├── docker-compose.yaml     # Service orchestration config
├── api/
│   └── v2/                 # API versioning
├── build/                  # Build scripts or CI/CD pipelines
├── common/                 # Shared utilities
│   ├── api_response/       # Standardized API responses
│   ├── err_response/       # Error handling
│   ├── log/                # Logging logic
│   └── limiter/            # Rate limiter settings
├── configs/                # Application configuration files
├── constant/               # Shared constants
├── internal/
│   ├── sqlclient/          # SQL database client config
│   └── redis/              # Redis logic and config
├── middleware/             # Middleware handlers
├── model/                  # Entity/model definitions
├── repository/             # Data access layer
├── server/
│   └── http/               # HTTP server setup
├── service/                # Business logic layer
├── tmp/                    # Temporary files
├── dto/                    # Data Transfer Objects
└── docs/                   # Project documentation
