# 📝 OTP AUTH with GIN framework

A simple OTP AUTH API with GIN framework and GOLANG 
---

## 🚀 Features

- 🔒 OTP-based login (via phone number)
- 🧾 JWT token generation
- 🗃 SQLite as the database
- 📦 GORM for ORM and migrations
- 🧪 OTP Generator And Rate limiter 

---

## 🗂️ Project Structure

```
.
├── internal/database/         # DB connection & migration logic
├── internal/handlers/         # Route handler functions (user, post)
├── internal/models/           # GORM models (User)
├── internal/routes/           # Route setup
├── internal/utils/            # JWT and OTP and RateLimiting utilities 
├── main.go           # App entry point
├── go.mod            # Go modules
├── .env              # Environment variables
```

---

## 🛠️ Installation & Run

1. **Clone the repository**

```bash
docker-compose up --build
```

2. **Set up environment variables**

Create a `.env` file in the root directory:

```env
JWT_SECRET=your_jwt_secret
```

3. **Run Locally**

```bash
go mod tidy
go run main.go
```

4. **Access the API**

API will run at: `http://localhost:8080`

---

## 📬 API Endpoints

### Authentication
| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| POST   | `/auth/v1/get-otp `      | Get a new OTP   |
| POST   | `/auth/v1/verify-otp`         | verify OTP  |



## 🧑‍💻 Author

Developed by [@dayiamin](https://github.com/dayiamin)  
Feel free to contribute, raise issues, or fork the project.

---
