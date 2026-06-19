package user

import (
	"context"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(
	ctx context.Context,
	user *User,
) error {

	query := `
	INSERT INTO users (
		id,
		username,
		password_hash,
		mfa_enabled,
		failed_attempts,
		created_at
	)
	VALUES (
		$1,$2,$3,$4,$5,$6
	)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.ID,
		user.Username,
		user.PasswordHash,
		user.MFAEnabled,
		user.FailedAttempts,
		user.CreatedAt,
	)

	return err
}

func (r *PostgresRepository) GetByUsername(
	ctx context.Context,
	username string,
) (*User, error) {

	query := `
	SELECT
		id,
		username,
		password_hash,
		mfa_enabled,
		mfa_secret,
		failed_attempts,
		locked_until,
		created_at,
		last_login
	FROM users
	WHERE username = $1
	`

	var user User

	err := r.db.QueryRowContext(
		ctx,
		query,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.MFAEnabled,
		&user.MFASecret,
		&user.FailedAttempts,
		&user.LockedUntil,
		&user.CreatedAt,
		&user.LastLogin,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresRepository) Update(
	ctx context.Context,
	user *User,
) error {

	query := `
	UPDATE users
	SET
		password_hash=$1,
		mfa_enabled=$2,
		mfa_secret=$3,
		failed_attempts=$4,
		locked_until=$5,
		last_login=$6
	WHERE id=$7
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.PasswordHash,
		user.MFAEnabled,
		user.MFASecret,
		user.FailedAttempts,
		user.LockedUntil,
		user.LastLogin,
		user.ID,
	)

	return err
}
