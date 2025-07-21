
# Go Auth Invoice API

A simple RESTful API built with Golang to demonstrate basic authentication and invoice listing. This project is built as part of backend developer portfolio with the following features:

- Basic user registration and login (mock logic)
- JWT token-based authentication (mock token)
- Secure access to invoice listing
- Clean architecture using MVC pattern
- Dockerized app ready for local development

## 🔧 Tech Stack

- Language: Go (Golang)
- Framework: net/http
- REST API
- JWT (mock)
- Docker
- Postman Collection (for testing)
- Git (Version Control)

---

## 📂 Folder Structure

```
go-auth-invoice/
├── cmd/
│   └── main.go           # Entry point
├── controller/
│   └── user_controller.go
├── middleware/
│   └── auth_middleware.go
├── model/
│   └── user_model.go
├── routes/
│   └── router.go
├── go.mod / go.sum
├── Dockerfile
└── README.md
```

---

## 🚀 Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/fakh1m/go-auth-invoice.git
cd go-auth-invoice
```

### 2. Run the App (Local)

```bash
go run ./cmd/main.go
```

### 3. Run via Docker

```bash
docker build -t go-auth-invoice .
docker run -p 8080:8080 go-auth-invoice
```

---

## 🧪 API Endpoints

### ✅ Register

**POST** `/api/register`

```json
{
  "username": "fakhim",
  "password": "123456"
}
```

🟢 _Response_:

```json
{
  "message": "User registered successfully"
}
```

---

### ✅ Login

**POST** `/api/login`

```json
{
  "username": "fakhim",
  "password": "123456"
}
```

🟢 _Response_:

```json
{
  "token": "mock-jwt-token"
}
```

---

### 🔐 Get Invoices (Protected)

**GET** `/api/invoices`

🟡 _Header_:

```
Authorization: Bearer mock-jwt-token
```

🟢 _Response_:

```json
[
  {
    "id": 1,
    "description": "Invoice for July",
    "amount": 150000
  },
  {
    "id": 2,
    "description": "Invoice for August",
    "amount": 200000
  }
]
```

---

## 🧪 Testing with Postman

Import the file [`go-auth-invoice.postman_collection.json`](./go-auth-invoice.postman_collection.json) into Postman and try all endpoints easily.

---

## 📌 Contribution Guide

1. Fork this repository
2. Create a new branch (`git checkout -b feature/my-feature`)
3. Commit your changes (`git commit -m 'Add something'`)
4. Push to your branch (`git push origin feature/my-feature`)
5. Open a Pull Request

---

## 📄 License

This project is licensed under MIT License.

---

## 🙋‍♂️ Author

**Fakhim Imam**  
[GitHub](https://github.com/fakh1m)