package user

import "time"

type User struct {
	ID           string
	Username     string
	PasswordHash string

	MFAEnabled bool
	MFASecret  *string

	FailedAttempts int
	LockedUntil    *time.Time

	CreatedAt time.Time
	LastLogin *time.Time
}
