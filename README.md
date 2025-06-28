# 🎉 Code Generator

**An automated Go source code generator with RESTful APIs — helping you quickly scaffold projects and CRUD logic, saved directly to the `C:/` drive.**

---

## 📚 Table of Contents

* [🔰 Introduction](#-introduction)
* [✨ Features](#-features)
* [⚙️ Requirements](#-requirements)
* [📦 Installation](#-installation)
* [🚀 Usage](#-usage)
* [🧩 API](#-api)

  * [1. Generate Project](#1-generate-project)
  * [2. Generate CRUD Code](#2-generate-crud-code)
* [📁 Generated Folder Structure](#-generated-folder-structure)

---

## 🔰 Introduction

**Code Generator** is a tool designed to automate the creation of Go-based backend projects.

It allows you to:

* ✅ Automatically generate a clean project folder structure
* ✅ Create CRUD code for custom entities

📍 All generated files are saved at: `C:/ProjectName/`

Perfect for developers who want to save time during project initialization or quickly build basic features.

---

## ✨ Features

* ✅ Auto-generate Go project folder structure
* ✅ Generate CRUD logic based on custom entity fields
* ✅ RESTful API interface for flexibility and ease
* ✅ Saves all code under `C:/` with organized layout

---

## ⚙️ Requirements

* Go **1.16** or newer
* OS: **Windows** (required for saving under `C:/`)
* An HTTP client (e.g., Postman, `cURL`, Insomnia, etc.)

---

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/kangourouuu/backend_template_personal.git

# Navigate into the project directory
cd backend_template_personal

# Install dependencies
go mod tidy

# Run the application
go run main.go
```

The server will start at: [http://localhost:9000](http://localhost:9000)

---

## 🚀 Usage

Once the app is running, you can:

1. Generate a Go project
2. Generate CRUD code for a specific entity

📌 **Notes:**

* Ensure the `C:/` drive has write permissions
* Use valid JSON in request bodies
* Use Postman or similar tools to interact with the APIs

---

## 🧩 API

### 1. Generate Project

* **Endpoint:** `POST /api/v1/generate-project`
* **Description:** Creates a new Go project at `C:/ProjectName/`

#### 📨 Sample Request

```json
{
  "project_name": "MyNewProject",
  "port": "8080"
}
```

📂 **Result:** Folder `C:/MyNewProject/` is created with the full project scaffold.

---

### 2. Generate CRUD Code

* **Endpoint:** `POST /api/v1/generate-crud`
* **Description:** Generates CRUD logic for a custom entity inside an existing project.

#### 📨 Sample Request

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

📂 **Result:** CRUD code for the `User` entity is generated in `C:/MyNewProject/`.

---

## 📁 Generated Folder Structure

```
C:/ProjectName/
├── main.go                 # Entry point of the application
├── Dockerfile              # Docker container definition
├── docker-compose.yaml     # Service orchestration
├── api/
│   └── v2/                 # API versioning
├── build/                  # Build scripts / CI configs
├── common/                 # Shared logic
│   ├── api_response/       # API response formatting
│   ├── err_response/       # Error handling logic
│   ├── log/                # Logging setup
│   └── limiter/            # Rate limiter configs
├── configs/                # Application config files
├── constant/               # Constant definitions
├── internal/
│   ├── sqlclient/          # SQL client configuration
│   └── redis/              # Redis setup
├── middleware/             # Middleware handlers
├── model/                  # Data models / entities
├── repository/             # Data access layer
├── server/
│   └── http/               # HTTP server logic
├── service/                # Business logic layer
├── tmp/                    # Temporary files
├── dto/                    # Data Transfer Objects
└── docs/                   # Project documentation
```

