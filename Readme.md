# 🧾 PayTrack - Payroll & Leave Management System

PayTrack is a multi-tenant backend system built using Go (Fiber) to help organizations manage employee leaves, automate monthly payroll generation, and enforce fine-grained role-based access control.

## 🚀 Features

- 🏢 Multi-organization support
- 👤 Role-based access control (Super Admin, HR, Manager, Employee)
- 🔐 Secure authentication via JWT
- 📝 Leave request and approval workflow
- 📅 Timezone-aware cron jobs for monthly payroll & attendance
- 🧾 Payroll PDF generation & S3 upload
- 📊 API documentation via Swagger / OpenAPI

---

## 🛠️ Tech Stack

| Layer        | Tech                        |
|--------------|-----------------------------|
| Language     | Go                          |
| Framework    | Fiber                       |
| ORM          | GORM                        |
| DB           | PostgreSQL                  |
| Auth         | JWT                         |
| Cloud        | AWS S3 (for payroll PDFs)   |
| Docs         | Swagger / OpenAPI 3         |
| Scheduler    | robfig/cron                 |

---

## 🧱 Architecture

```mermaid
flowchart TD
    A[main.go] --> B[Setup]
    B --> C[Router]
    C --> D[Handlers/Controllers]
    D --> E[Service Layer]
    E --> F[Repository Layer (Generic)]
    F --> G[GORM + PostgreSQL]

    subgraph External
      H[AWS S3] 
    end

    E --> H
```

- **Handlers**: HTTP logic and request validation
- **Service**: Core business logic (leave approval, payroll generation)
- **Repository**: Generic interface for DB operations
- **Cron jobs**: Run on 1st and last day of month (timezone-aware)

---

## 🗃️ Database Schema

```mermaid
erDiagram
    Organization ||--o{ User : has
    Organization ||--o{ Role : has
    Role ||--o{ User : assigned
    User ||--o{ LeaveRequest : submits
    User ||--o{ Payroll : receives
    User }o--o{ Invitation : invited_by

    Organization {
      UUID id PK
      string name
    }

    Role {
      UUID id PK
      string name
      UUID organization_id FK
    }

    User {
      UUID id PK
      string name
      string email
      string password_hash
      UUID role_id FK
      UUID organization_id FK
    }

    LeaveRequest {
      UUID id PK
      UUID user_id FK
      date start_date
      date end_date
      string status
    }

    Payroll {
      UUID id PK
      UUID user_id FK
      date month
      string pdf_url
    }

    Invitation {
      UUID id PK
      string email
      UUID role_id FK
      string first_name
      string last_name
      UUID invited_by_id FK
      UUID organization_id FK
      time expires_at
      string invite_token
      int status
      text message
      string department
      string designation
      decimal base_salary
      decimal bonus
      decimal overtime_rate
      decimal allowances
      decimal health_insurance
      decimal retirement_benefits
      decimal stock_options
      decimal stock_options_vested
      decimal stock_options_unvested
      decimal stock_options_strike_price
      bigint stock_options_quantity
      string stock_options_type
      string stock_options_status
      time joining_date
      string employment_type
      string reporting_to_id
      string phone_number
    }
```

---

## 🕒 Cron Jobs

| Schedule | Task                                  | Timezone      |
|----------|---------------------------------------|---------------|
| 1st of month @ 00:00 | Generate monthly payrolls (PDF + upload to S3) | Configurable (e.g. Asia/Kolkata) |
| Last day of month @ 00:00 | Trigger attendance lock / reporting | Configurable |

---

## 📦 Setup & Run

### Prerequisites

- Go ≥ 1.20
- PostgreSQL
- AWS credentials (for S3)
- (Optional) Docker for DB

### Clone & Install

```bash
git clone https://github.com/yourusername/paytrack.git
cd paytrack
go mod tidy
```

### Env Configuration

Create a `.env` file:

```env
PORT=3000
DB_URL=postgres://user:pass@localhost:5432/paytrack?sslmode=disable
JWT_SECRET=your_secret
S3_BUCKET_NAME=your-bucket
AWS_REGION=ap-south-1
AWS_ACCESS_KEY_ID=...
AWS_SECRET_ACCESS_KEY=...
TIMEZONE=Asia/Kolkata
```

### Run

```bash
go run main.go
```

---

## 📘 API Documentation

After running the app, access:

```
http://localhost:3000/swagger/index.html
```

---

## ✨ Contributing

Contributions are welcome! Please fork the repo and submit a PR.

---

## 🛡️ License

MIT © 2025 YourName
