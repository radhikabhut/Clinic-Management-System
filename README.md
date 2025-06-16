# Clinic-Management-System



## 📌 Overview

The **Clinic Management System** is a web application built in **Go (Golang)** using the **Gin web framework**, connected to a **PostgreSQL** database. It supports user authentication using **JWT** and allows different roles (receptionist and doctor) to perform role-specific actions and **Postman** is used for API testing.

---

## ⚙️ Tech Stack

* **Backend:** Go (Gin Framework)
* **Database:** PostgreSQL
* **Authentication:** JWT (JSON Web Tokens)
* **Testing Tool:** Postman

---

## 🔂️ Project Structure

```bash
clinic-management/
├── config/          # Database & environment configuration
├── handler/         # business logic
├── model/           # Data models and request/response structs
├── router/          # API route definitions
├── view/            # core logic
├── utils/           # Utility functions (e.g., JWT generation)
├── main.go          # Application entry point
├── go.mod
└── .env             # Environment variables (DB creds, secrets)
```

---

## 👥 User Roles

* **Receptionist:**

  * Create patient
  * List all patients
  * Get patient by ID
  * Update patient
  * Delete patient

* **Doctor:**

  * List all patients
  * Get patient by ID
  * Update patient (e.g., treatment details)

---

## 🔐 Authentication Flow (JWT)

1. **Registration:**

   * Users (receptionist/doctor) are created in the database with `username`, `password`, and `role`.
2. **Login:**

   * Users log in using their credentials.
   * JWT token is generated and returned.
3. **Authorization:**

   * Protected routes validate the token and allow access based on the user's role.

---

## 📟 Features

| Feature           | Description                            |
| ----------------- | -------------------------------------- |
| User Registration | Admin-level or pre-inserted users      |
| JWT Login         | Secure token-based login               |
| Role-Based Access | Restricts endpoints by user role       |
| Patient CRUD      | Add, view, update, and delete patients |
| Postman APIs      | API endpoints tested via Postman       |

---

## 🧹 Modules Explained

### 1. **Model Layer (`model/`)**

Defines the structure of:

* Database entities (`User`, `Patient`)
* API request/response structs (`UserRequest`, `PatientRequest`)

### 2. **Database Layer (`database/`)**

Responsible for:

* Connecting to PostgreSQL
* Query logic: create, read, update, delete
* User and patient operations

### 3. **Handler Layer (`handler/`)**

* Maps HTTP request bodies to model structs
* Calls database  functions
* business logic is here

### 4. *View Layer (`view/`)**

* Core business logic
* Calls handler functions

### 5. **Router Layer (`router/`)**

* Defines all API routes with appropriate middleware (e.g., JWT auth)
* Groups routes based on user roles





## 🔀 API Endpoints

| Method | Endpoint       | Role         | Description                              |
| ------ | -------------- | ------------ | -----------------------------------------|
| POST   | `/login`       | All          | Login and receive JWT                    |
| POST   | `/patient`     | Receptionist | Create a new patient                     |
| GET    | `/patients`    | All          | List all patients                        |
| GET    | `/patient/:id` | All          | Get patient by ID                        |
| PUT    | `/patient/:id` | All          | Update patient info                      |
| DELETE | `/patient/:id` | Receptionist | Delete patient by ID                     |
| POST   | `/users/:id`   |              | register user as doctor or receptionist  |

---

## 🥺 Postman Testing

* Import all endpoints into Postman.
* Use the `/login` endpoint to retrieve a JWT.
* Add the token to `Authorization > Bearer Token` in headers.
* Test each endpoint with appropriate roles.

---

## 📏 Environment Variables (`.env`)

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=clinic_db
JWT_SECRET=secret
```

Use `godotenv` to load these into your Go app.

---

## 🚀 How to Run

1. Clone the repository.
2. Create a PostgreSQL DB and run schema SQL if required.
3. Update `.env` file with your DB details.
4. Run the app:

```bash
go run main.go
```

5. Open HTML form in browser or test APIs in Postman.

---
