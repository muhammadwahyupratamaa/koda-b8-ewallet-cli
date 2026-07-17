package repository

import (
	"context"

	"koda-b8-ewallet-cli/internal/model"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	err := r.db.QueryRow(
	context.Background(),
	`
	INSERT INTO users (
		full_name,
		user_name,
		age,
		address,
		email,
		password
	)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	)
	RETURNING id
	`,
	user.FullName,
	user.UserName,
	user.Age,
	user.Address,
	user.Email,
	user.Password,
	).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User

	err := r.db.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			full_name,
			user_name,
			age,
			address,
			email,
			password,
			created_at,
			updated_at
		FROM users
		WHERE user_name = $1
		`,
		username,
	).Scan(
		&user.ID,
		&user.FullName,
		&user.UserName,
		&user.Age,
		&user.Address,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}