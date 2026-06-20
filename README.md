# CipherGate

A secure, containerized Command-Line Authentication System built with Go, PostgreSQL, Docker, and TOTP-based Multi-Factor Authentication (MFA).

CipherGate provides user registration, authentication, session management, account lockout protection, and Google/Microsoft Authenticator compatible MFA through an interactive CLI experience.

---

## Features

### Authentication

* User registration
* Secure login with username and password
* Password hashing using bcrypt
* Account lockout after multiple failed login attempts
* Last login tracking
* Current login tracking

### Multi-Factor Authentication

* Enable TOTP-based MFA
* Disable MFA
* Compatible with:

  * Google Authenticator
  * Microsoft Authenticator
  * Other RFC 6238 compatible applications

* QR Code enrollment directly in terminal
* MFA verification during login

### Session Management

* Session creation after successful authentication
* Configurable session timeout
* Session expiration validation
* Session logout support

### Interactive CLI

* Interactive command shell
* Command history
* Tab completion
* Context-aware commands
* Helpful error messages and feedback

### Persistence

* PostgreSQL database
* Dockerized database deployment
* Database migrations
* Persistent Docker volumes

---

## Architecture

CipherGate follows a layered architecture:

```text
┌─────────────┐
│     CLI     │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  Services   │
├─────────────┤
│ User        │
│ Session     │
│ MFA         │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ Repository  │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ PostgreSQL  │
└─────────────┘
```

---

## Project Structure

```text
ciphergate/
│
├── internal/
│   ├── auth/
│   │   └── context.go
│   │
│   ├── cli/
│   │   ├── login.go
│   │   ├── logout.go
│   │   ├── register.go
│   │   ├── whoami.go
│   │   ├── enable_2fa.go
│   │   ├── disable_2fa.go
│   │   ├── shell.go
│   │   ├── session_guard.go
│   │   └── helpers.go
│   │
│   ├── database/
│   │   └── postgres.go
│   │
│   ├── migration/
│   │   └── runner.go
│   │
│   ├── mfa/
│   │   └── service.go
│   │
│   ├── session/
│   │   ├── model.go
│   │   ├── repository.go
│   │   ├── postgres_repository.go
│   │   └── service.go
│   │
│   └── user/
│       ├── model.go
│       ├── repository.go
│       ├── postgres_repository.go
│       └── service.go
│
├── migrations/
│   ├── 000_schema_migrations.sql
│   ├── 001_create_users.sql
│   └── 002_create_sessions.sql
│
├── .env
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

## Technology Stack

| Component          | Technology                                     |
| ------------------ | ---------------------------------------------- |
| Language           | Go                                             |
| Database           | PostgreSQL                                     |
| Containerization   | Docker                                         |
| Password Hashing   | bcrypt                                         |
| MFA                | TOTP                                           |
| CLI                | Readline                                       |
| Database Driver    | lib/pq                                         |

---

## Prerequisites

* Go 1.24+
* Docker
* Docker Compose

---

## Getting Started

### Clone Repository

```bash
git clone https://github.com/shubhamranswal/CipherGate.git
cd ciphergate
```

---

### Configure Environment

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=ciphergate
DB_SSLMODE=disable
```

---

### Start PostgreSQL

```bash
docker compose up -d
```

Verify:

```bash
docker ps
```

---

### Run Application

```bash
go run .
```

Expected startup:

```text
🔐 CipherGate v0.1
✅ Connected to PostgreSQL
```

---

## Quick Demo

### Register User

```text
ciphergate> register

Username: shubham
Password:
Confirm Password:

✅ User registered successfully
```

---

### Login

```text
ciphergate> login

Username: shubham
Password:

✅ Login successful
```

---

### Enable MFA

```text
ciphergate(shubham)> enable-2fa
```

Scan the QR code using Microsoft Authenticator or Google Authenticator and enter the generated code.

---

### Logout

```text
ciphergate(shubham)> logout
```

---

### Login With MFA

```text
ciphergate> login

Username: shubham
Password:

🔐 MFA Verification Required

TOTP Code: 123456

✅ Login successful
```

---

## Available Commands

### Guest Commands

| Command  | Description                |
| -------- | -------------------------- |
| register | Create a new user account  |
| login    | Authenticate user          |
| help     | Display available commands |
| exit     | Exit application           |

---

### Authenticated Commands

| Command     | Description                  |
| ----------- | ---------------------------- |
| whoami      | Display current user details |
| enable-2fa  | Enable MFA                   |
| disable-2fa | Disable MFA                  |
| logout      | Logout current session       |
| help        | Display available commands   |

---

## User Information

After successful login:

```text
👤 User Details
────────────────────────────────────────

Username
Registration Date
MFA Status
Last Login
Session Status
Session Expiration Time
```

---

## Security Features

### Password Security

* Passwords never stored in plaintext
* bcrypt password hashing
* Secure password verification

### Account Lockout

After multiple failed login attempts:

```text
Maximum Failed Attempts: 5
Lockout Duration: 15 Minutes
```

Account access is temporarily blocked until the lockout period expires.

### Multi-Factor Authentication

* RFC 6238 compliant TOTP
* QR code onboarding
* Compatible with standard authenticator applications

### Session Security

* Session IDs generated using UUIDs
* Session expiration support
* Session validation before authenticated operations
* Automatic logout on session expiration

---

## Database Schema

### Users Table

```sql
users
(
    id,
    username,
    password_hash,
    mfa_enabled,
    mfa_secret,
    failed_attempts,
    locked_until,
    current_login,
    last_login,
    created_at
)
```

---

### Sessions Table

```sql
sessions
(
    id,
    user_id,
    created_at,
    expires_at,
    active
)
```

---

## Session Management

Default session timeout: 30 Minutes

Expired sessions are automatically invalidated and require re-authentication.

---

## Docker

Start services:

```bash
docker compose up -d
```

Stop services:

```bash
docker compose down
```

Remove services and volumes:

```bash
docker compose down -v
```

---

## Manual Testing Performed

* User registration
* Duplicate username validation
* Password validation
* Successful login
* Failed login
* Account lockout
* Session creation
* Session expiration
* Logout
* MFA enable
* MFA disable
* MFA login verification
* QR code enrollment
* Database persistence
* Docker startup and shutdown

---

## Future Enhancements

* Password reset workflow
* Role-based access control
* Audit logging
