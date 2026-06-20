package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shubhamranswal/ciphergate/internal/session"
	"golang.org/x/crypto/bcrypt"
)

const (
	MaxFailedAttempts = 5
	LockoutDuration   = 15 * time.Minute
)

var ErrMFARequired = errors.New(
	"mfa required",
)

type Service struct {
	repo       Repository
	sessionSvc *session.Service
}

func NewService(
	repo Repository,
	sessionSvc *session.Service,
) *Service {

	return &Service{
		repo:       repo,
		sessionSvc: sessionSvc,
	}
}

func (s *Service) Register(
	ctx context.Context,
	username string,
	password string,
) error {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := &User{
		ID:             uuid.New().String(),
		Username:       username,
		PasswordHash:   string(hash),
		MFAEnabled:     false,
		FailedAttempts: 0,
		CreatedAt:      time.Now().UTC(),
	}

	return s.repo.Create(ctx, user)
}

func (s *Service) UsernameAvailable(
	ctx context.Context,
	username string,
) (bool, error) {

	existing, err := s.repo.GetByUsername(
		ctx,
		username,
	)

	if err != nil {
		return false, err
	}

	return existing == nil, nil
}

func (s *Service) ValidateUsername(
	username string,
) error {

	if len(username) < 3 {
		return errors.New(
			"username must be at least 3 characters",
		)
	}

	if len(username) > 50 {
		return errors.New(
			"username cannot exceed 50 characters",
		)
	}

	return nil
}

func (s *Service) ValidatePassword(
	password string,
) error {

	if len(password) < 8 {
		return errors.New(
			"password must be at least 8 characters",
		)
	}

	return nil
}

func (s *Service) Login(
	ctx context.Context,
	username string,
	password string,
) (*User, *session.Session, error) {

	user, err := s.repo.GetByUsername(
		ctx,
		username,
	)

	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		return nil, nil,
			errors.New(
				"invalid username or password",
			)
	}

	if user.LockedUntil != nil &&
		time.Now().Before(*user.LockedUntil) {

		return nil, nil,
			fmt.Errorf(
				"account locked until %s",
				user.LockedUntil.Format(
					"2006-01-02 15:04:05 UTC",
				),
			)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {

		user.FailedAttempts++

		if user.FailedAttempts >= MaxFailedAttempts {

			lockUntil := time.Now().UTC().
				Add(LockoutDuration)

			user.LockedUntil =
				&lockUntil

			user.FailedAttempts = 0
		}

		if updateErr := s.repo.Update(
			ctx,
			user,
		); updateErr != nil {

			return nil, nil,
				updateErr
		}

		return nil, nil,
			errors.New(
				"invalid username or password",
			)
	}

	user.FailedAttempts = 0
	user.LockedUntil = nil

	now := time.Now().UTC()

	user.LastLogin = user.CurrentLogin
	user.CurrentLogin = &now

	err = s.repo.Update(
		ctx,
		user,
	)

	if err != nil {
		return nil, nil, err
	}

	if user.MFAEnabled {
		return user, nil, ErrMFARequired
	}

	sessionObj, err := s.sessionSvc.Create(
		ctx,
		user.ID,
	)

	if err != nil {
		return nil, nil, err
	}

	return user, sessionObj, nil
}

func (s *Service) Update(
	ctx context.Context,
	user *User,
) error {

	return s.repo.Update(
		ctx,
		user,
	)
}

func (s *Service) UpdateLastLogin(
	ctx context.Context,
	user *User,
) error {

	now := time.Now().UTC()

	user.LastLogin = &now

	return s.repo.Update(
		ctx,
		user,
	)
}
