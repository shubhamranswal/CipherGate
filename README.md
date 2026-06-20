# CipherGate

A secure containerized CLI authentication system built with Go, PostgreSQL, Docker, and TOTP-based Multi-Factor Authentication.

## Features

* User registration and authentication
* Secure password hashing using bcrypt
* TOTP-based Multi-Factor Authentication (Google/Microsoft Authenticator compatible)
* Account lockout after repeated failed login attempts
* Session management with configurable timeout
* PostgreSQL persistence
* Interactive CLI with command history and tab completion
* Dockerized deployment
* Database migrations

## Architecture

CipherGate follows a layered architecture:

CLI Layer
→ Service Layer
→ Repository Layer
→ PostgreSQL

### Components

* CLI: Interactive terminal interface
* User Service: Authentication and user management
* MFA Service: TOTP generation and verification
* Session Service: Session lifecycle management
* Repository Layer: Database access abstraction
* PostgreSQL: Persistent storage

## Requirements

* Go 1.24+
* Docker
* Docker Compose

## Setup

Clone the repository:

```bash
git clone <repo-url>
cd ciphergate
```

Start PostgreSQL:

```bash
docker compose up -d
```

Run the application:

```bash
go run .
```

## Available Commands

### Guest Commands

* register
* login
* help
* exit

### Authenticated Commands

* whoami
* enable-2fa
* disable-2fa
* logout
* help

## MFA Setup

1. Login
2. Run enable-2fa
3. Scan the displayed QR code using Microsoft Authenticator or Google Authenticator
4. Enter the generated TOTP code
5. MFA is now enabled

## Security Features

* bcrypt password hashing
* TOTP-based MFA
* Session expiration
* Account lockout protection
* Password verification using constant-time bcrypt comparison

## Database Schema

### users

* id
* username
* password_hash
* mfa_enabled
* mfa_secret
* failed_attempts
* locked_until
* current_login
* last_login
* created_at

### sessions

* id
* user_id
* active
* created_at
* expires_at

## Future Improvements

* Password reset workflow
* Admin roles and permissions
* Audit logging
* Refresh tokens
* Email verification
* Rate limiting
