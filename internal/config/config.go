package config

import (
	"os"
	"strconv"
	"time"
)

var (
	SessionTimeout    = 30 * time.Minute
	LockoutDuration   = 15 * time.Minute
	MaxFailedAttempts = 5
)

func Load() {

	if v := os.Getenv("SESSION_TIMEOUT_MINUTES"); v != "" {
		if minutes, err := strconv.Atoi(v); err == nil {
			SessionTimeout = time.Duration(minutes) * time.Minute
		}
	}

	if v := os.Getenv("LOCKOUT_DURATION_MINUTES"); v != "" {
		if minutes, err := strconv.Atoi(v); err == nil {
			LockoutDuration = time.Duration(minutes) * time.Minute
		}
	}

	if v := os.Getenv("MAX_FAILED_ATTEMPTS"); v != "" {
		if attempts, err := strconv.Atoi(v); err == nil {
			MaxFailedAttempts = attempts
		}
	}
}
