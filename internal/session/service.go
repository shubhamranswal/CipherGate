package session

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shubhamranswal/ciphergate/internal/config"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, userID string) (*Session, error) {
	now := time.Now().UTC()

	session := &Session{
		ID:        uuid.New().String(),
		UserID:    userID,
		CreatedAt: now,
		ExpiresAt: now.Add(config.SessionTimeout),
		Active:    true,
	}

	err := s.repo.Create(ctx, session)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Service) Validate(ctx context.Context, sessionID string) (*Session, error) {
	session, err := s.repo.GetByID(ctx, sessionID)

	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, errors.New("session not found")
	}

	if !session.Active {
		return nil, errors.New("session inactive")
	}

	if time.Now().UTC().After(session.ExpiresAt) {
		return nil, errors.New("session expired")
	}

	return session, nil

}

func (s *Service) Deactivate(ctx context.Context, sessionID string) error {
	return s.repo.Deactivate(ctx, sessionID)
}
