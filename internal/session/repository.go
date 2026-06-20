package session

import "context"

type Repository interface {
	Create(ctx context.Context, session *Session) error
	GetByID(ctx context.Context, id string) (*Session, error)
	Deactivate(ctx context.Context, id string) error
}
