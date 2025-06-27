🎉 Code Generator

An automated Go code generator with simple REST APIs — quickly bootstrap projects and generate CRUD code, saved directly to your C:/ drive.

📚 Table of Contents
	•	🔰 Introduction
	•	✨ Features
	•	⚙️ Requirements
	•	📦 Installation
	•	🚀 Usage
	•	🧩 API Reference
	•	1. Generate Project
	•	2. Generate CRUD Code
	•	📁 Generated Folder Structure

🔰 Introduction

Code Generator is a tool designed to automate Go project scaffolding and CRUD code generation.

You can use the provided APIs to:
	•	Generate a standard Go project folder structure
	•	Generate CRUD boilerplate for your custom entities

📍 All generated code is saved to: C:/ProjectName/

This tool is ideal for developers looking to speed up project setup or quickly build basic backend logic.

✨ Features
	•	✅ Automatically creates a complete Go project structure
	•	✅ Generates customizable CRUD code for any entity
	•	✅ Simple and flexible REST APIs
	•	✅ Saves code in C:/ with clean folder organization

⚙️ Requirements
	•	Go 1.16 or higher
	•	Windows OS (required for writing files to C:/)
	•	API tools like Postman, cURL, Insomnia, etc.

📦 Installation

# Clone the repository
git clone https://github.com/kangourouuu/backend_template_personal.git

# Navigate into the project directory
cd backend_template_personal

# Download dependencies
go mod tidy

# Run the application
go run main.go

📍 The app will run at: http://localhost:9000

🚀 Usage

Once the application is running, you can:
	1.	Generate a new Go project
	2.	Generate CRUD code for your defined entity

📌 Notes:
	•	Ensure C:/ has write permissions
	•	Use valid JSON in your API requests
	•	You can send requests using Postman or any HTTP client

🧩 API Reference

1. Generate Project
	•	Endpoint: POST /api/v1/generate-project
	•	Description: Creates a new Go project folder at C:/ProjectName/

📨 Request (JSON)

{
  "project_name": "MyNewProject",
  "port": "8080"
}

📂 Result: A project folder will be created at C:/MyNewProject/

2. Generate CRUD Code
	•	Endpoint: POST /api/v1/generate-crud
	•	Description: Generates CRUD boilerplate for a custom entity with your defined fields

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

📂 Result: CRUD code will be generated under C:/MyNewProject/

📁 Generated Folder Structure

C:/ProjectName/
├── main.go                 # Entry point of the application
├── Dockerfile              # Docker container build file
├── docker-compose.yaml     # Service orchestration
├── api/
│   └── v2/                 # API versioning
├── build/                  # Build scripts or CI/CD pipelines
├── common/                 # Shared logic
│   ├── api_response/       # Standard API response formatting
│   ├── err_response/       # Error handling logic
│   ├── log/                # Logging
│   └── limiter/            # Rate limiter configuration
├── configs/                # Application configuration
├── constant/               # Shared constants
├── internal/
│   ├── sqlclient/          # SQL client setup
│   └── redis/              # Redis integration logic
├── middleware/             # Request/response middleware
├── model/                  # Entity/model definitions
├── repository/             # Data access layer
├── server/
│   └── http/               # HTTP server setup
├── service/                # Business logic layer
├── tmp/                    # Temporary files
├── dto/                    # Data Transfer Objects (DTOs)
└── docs/                   # Project documentation
