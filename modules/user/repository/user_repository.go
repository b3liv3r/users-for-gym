package repository

import (
	"context"
	"github.com/b3liv3r/users-for-gym/modules/user/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type UserRepositoryDB struct {
	db *sqlx.DB
}

func NewUserRepositoryDB(db *sqlx.DB) UsererRepository {
	return &UserRepositoryDB{db: db}
}

func (r *UserRepositoryDB) Create(ctx context.Context, userId int, username, phone, email string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id, username, phone, email, subscription_lvl, current_gym_id) VALUES ($1, $2, $3, $4)", userId, username, phone, email, 0, 0)
	return err
}

func (r *UserRepositoryDB) Profile(ctx context.Context, userId int) (models.User, error) {
	var user models.User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", userId)
	return user, err
}

func (r *UserRepositoryDB) Update(ctx context.Context, userId int, username, phone, email string) error {
	query := "UPDATE users SET "
	args := []interface{}{userId}

	if username != "" {
		query += "username = $2, "
		args = append(args, username)
	}

	if phone != "" {
		query += "phone = $3, "
		args = append(args, phone)
	}

	if email != "" {
		query += "email = $4, "
		args = append(args, email)
	}

	// Удаляем последнюю запятую и пробел из запроса
	query = strings.TrimSuffix(query, ", ")

	// Добавляем условие WHERE
	query += " WHERE id = $1"

	// Выполняем запрос
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *UserRepositoryDB) UpdateSubLvl(ctx context.Context, userId int, level int) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET subscription_lvl = $2 WHERE id = $1", userId, level)
	return err
}

func (r *UserRepositoryDB) UpdateGymID(ctx context.Context, userId int, gymId int) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET current_gym_id = $2 WHERE id = $1", userId, gymId)
	return err
}
