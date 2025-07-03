# 🎉 Code Generator

A backend code scaffolding tool built in Go — enabling developers to instantly generate full-featured RESTful projects and CRUD logic using flexible API inputs.

📍 All generated code is saved to: `C:/ProjectName/`

Ideal for bootstrapping backend services in seconds — with Docker, PostgreSQL, Redis, and built-in unit tests ready from the start.

---

## 📚 Table of Contents
- 🔰 Introduction
- ✨ Features
- ⚙️ Requirements
- 📦 Installation
- 🧰 CLI Tool
- 🐳 Docker Usage
- 🚀 Usage
- 🧩 API Reference
  - Generate Project
  - Generate CRUD Code
- 📁 Generated Folder Structure

---

## 🔰 Introduction

**Code Generator** is a developer tool that automates the process of creating scalable Go-based backend projects. It leverages Go’s `text/template` package to render boilerplate code, saving development time and ensuring a consistent code structure.

### What it does:
- ✅ Scaffolds project structure and boilerplate files
- ✅ Auto-generates CRUD endpoints based on user-defined entity fields
- ✅ Uses RESTful APIs or CLI commands to trigger code generation dynamically

---

## ✨ Features
- ✅ Generate Go project structure with API layers and clean layout
- ✅ Auto-create CRUD logic from custom fields and entity names
- ✅ Integrated PostgreSQL and Redis setup in the generated output
- ✅ Generates Docker/Docker Compose configuration for containerized development
- ✅ Prewritten unit tests using Go’s `testing` and `testify`
- ✅ Projects saved to `C:/ProjectName/`
- ✅ Use both HTTP APIs and CLI for flexible code generation workflows

---

## ⚙️ Requirements
- Go 1.16+
- OS: Windows (for C:/ path support)
- HTTP client (Postman, Insomnia, curl, etc.) or terminal for CLI

---

## 📦 Installation

```bash
git clone https://github.com/kangourouuu/backend_template_personal.git
cd backend_template_personal
go mod tidy
go run main.go
```

> The server runs at: `http://localhost:9000`

---

## 🧰 CLI Tool
Besides HTTP API support, Code Generator also comes with a CLI tool (powered by Cobra).

# Run CLI:
```bash
go run cmd/main.go --help
```
# Available CLI Commands:
```bash
- Generate Project
go run cmd/main.go generate-project -n ProjectName -p Port ( int )

Generate CRUD Code
go run cmd/main.go generate-crud -E entityName -n projectName -f "name:string,price:float,available:bool"

Note: The quantity of fields can adjust with the format "name:type" and is seperated with ","
```


## 🐳 Docker Usage

You can also run this project in a containerized environment using Docker or Docker Compose.

```bash
🔨 Build Docker Image
docker build -t code-generator .
▶️ Run with Docker
docker run -p 9000:9000 code-generator
🧱 Run with Docker Compose (Recommended)
docker-compose up --build

⚙️ Additional Docker Commands
| Command                      | Description                                |
| ---------------------------- | ------------------------------------------ |
| `docker-compose up`          | Start services                             |
| `docker-compose up --build`  | Build and start services                   |
| `docker-compose stop`        | Stop containers                            |
| `docker-compose down -v`     | Remove containers, network and volumes     |
| `docker ps`                  | List running containers                    |
| `docker logs <container-id>` | View logs from a specific container        |

📌 After startup, the service will be accessible at: http://localhost:9000

This will automatically build and run the service as defined in docker-compose.yml.
```
---

## 🚀 Usage

You can use the provided APIs to:

1. ✅ Generate a new Go backend project
2. ✅ Generate CRUD logic for custom entities

📌 Make sure:
- The `C:/` drive is writable
- JSON body is valid
- Use Postman or similar tools for testing

---

## 🧩 API Reference

### 1. Generate Project

**Endpoint**: `POST /api/v1/generate-project`  
**Description**: Creates a new Go backend project at `C:/ProjectName/`

#### 📨 Sample Request:
```json
{
  "project_name": "MyNewProject",
  "port": "8080"
}
```

#### 📁 Output:
Project scaffold created in `C:/MyNewProject/`

---

### 2. Generate CRUD Code

**Endpoint**: `POST /api/v1/generate-crud`  
**Description**: Adds CRUD logic for a specific entity to an existing generated project

#### 📨 Sample Request:
```json
{
  "entity_name": "User",
  "entity_name_lower": "user",
  "fields": [
    { "name": "ID", "type": "int" },
    { "name": "Name", "type": "string" }
  ]
}
```

#### 📁 Output:
CRUD files added to: `C:/MyNewProject/`

---

## 📁 Generated Folder Structure

```
C:/ProjectName/
├── main.go                 # Entry point
├── Dockerfile              # Container setup
├── docker-compose.yaml     # Service orchestration
├── api/
│   └── v2/                 # Route handling
├── build/                  # Build/CI scripts
├── common/                 # Shared utils: errors, responses
│   ├── api_response/
│   ├── err_response/
│   ├── log/
│   └── limiter/
├── configs/                # Config files
├── constant/               # Constant values
├── internal/
│   ├── sqlclient/          # DB setup
│   └── redis/              # Redis setup
├── middleware/             # Middlewares
├── model/                  # Data models
├── repository/             # Data access
├── server/
│   └── http/               # HTTP server bootstrap
├── service/                # Business logic
├── tmp/                    # Temp files
├── dto/                    # DTO structs
└── docs/                   # Project documentation
```