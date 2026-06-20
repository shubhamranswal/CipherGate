package session

import (
	"context"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, session *Session) error {

	query := `
	INSERT INTO sessions (
		id,
		user_id,
		created_at,
		expires_at,
		active
	)
	VALUES (
		$1,$2,$3,$4,$5
	)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		session.ID,
		session.UserID,
		session.CreatedAt,
		session.ExpiresAt,
		session.Active,
	)

	return err
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*Session, error) {

	query := `
	SELECT
		id,
		user_id,
		created_at,
		expires_at,
		active
	FROM sessions
	WHERE id = $1
	`

	var session Session

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&session.ID,
		&session.UserID,
		&session.CreatedAt,
		&session.ExpiresAt,
		&session.Active,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *PostgresRepository) Deactivate(ctx context.Context, id string) error {

	_, err := r.db.ExecContext(
		ctx,
		`
		UPDATE sessions
		SET active = FALSE
		WHERE id = $1
		`,
		id,
	)

	return err
}
