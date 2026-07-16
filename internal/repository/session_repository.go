package repository

import (
	"context"

	"koda-b8-ewallet-cli/internal/model"

	"github.com/jackc/pgx/v5"
)

type SessionRepository struct {
	db *pgx.Conn
}

func NewSessionRepository(db *pgx.Conn) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (r *SessionRepository) CreateSession(session *model.UserSession) error {
	_, err := r.db.Exec(
		context.Background(),
		`
		INSERT INTO user_sessions (
			user_id,
			session_token,
			expired_at
		)
		VALUES (
			$1,
			$2,
			$3
		)
		`,
		session.UserID,
		session.SessionToken,
		session.ExpiredAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) DeleteSessionByUserID(userID int64) error {

	_, err := r.db.Exec(
		context.Background(),
		`
		DELETE FROM user_sessions
		WHERE user_id = $1
		`,
		userID,
	)

	return err
}