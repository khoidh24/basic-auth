# 🌿 Basic Authentication App

A simple and clean authentication system using Go Fiber and MongoDB.

---

## 📚 Tech Stack

- ⚡ [Fiber](https://gofiber.io/) — Fast HTTP web framework
- 🍃 [MongoDB + MGM](https://github.com/kamva/mgm) — ODM for MongoDB
- 🔐 [JWT](https://jwt.io/) — Secure authentication
- 🔒 [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — Password hashing
- 📘 [Scalar](https://scalar.com/) — API documentation tool

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/your-repo/auth-app.git
cd auth-app
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Set up environment variables

Create a `.env` file in the root directory and add the following:

```env
DOMAIN=your-domain
PORT=3000
MONGODB_URI=your-mongodb-uri
JWT_SECRET=your-jwt-secret
```

### 4. Run the application

```bash
go run cmd/main.go
```

The app will run on `localhost:3000`. You can now register, login, and manage users.

### 5. API Documentation

The API documentation is generated using Scalar. You can access it at `localhost:3000/api/v1/reference`.