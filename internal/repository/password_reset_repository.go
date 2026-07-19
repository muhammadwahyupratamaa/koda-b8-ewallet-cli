package repository

import (
	"context"

	"koda-b8-ewallet-cli/internal/model"

	"github.com/jackc/pgx/v5"
)

type PasswordResetRepository struct {
	db *pgx.Conn
}

func NewPasswordResetRepository(db *pgx.Conn) *PasswordResetRepository {
	return &PasswordResetRepository{
		db: db,
	}
}

func (r *PasswordResetRepository) CreateResetToken(reset *model.PasswordReset) error {

	query := `
	INSERT INTO password_resets (
		user_id,
		reset_token,
		expired_at
	)
	VALUES ($1,$2,$3)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		reset.UserID,
		reset.ResetToken,
		reset.ExpiredAt,
	)

	return err
}

func (r *PasswordResetRepository) GetResetToken(token string) (*model.PasswordReset, error) {

	query := `
	SELECT
		id,
		user_id,
		reset_token,
		expired_at,
		created_at
	FROM password_resets
	WHERE reset_token = $1
	`

	var reset model.PasswordReset

	err := r.db.QueryRow(
		context.Background(),
		query,
		token,
	).Scan(
		&reset.ID,
		&reset.UserID,
		&reset.ResetToken,
		&reset.ExpiredAt,
		&reset.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &reset, nil
}

func (r *PasswordResetRepository) DeleteResetToken(id int) error {

	query := `
	DELETE FROM password_resets
	WHERE id = $1
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}