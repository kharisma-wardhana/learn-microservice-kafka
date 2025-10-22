package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kharisma-wardhana/learn-microservice-kafka/user-service/internal/model"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(ctx context.Context, u *model.User, password string) error {
	q := `
		INSERT INTO users (email, password, fullname, role, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, now(), now())
	`
	_, err := r.db.Exec(ctx, q, u.Email, password, u.Fullname, u.Role)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	q := `
		SELECT id, email, fullname, role, created_at, updated_at 
		FROM users 
		WHERE email=$1
	`
	row := r.db.QueryRow(ctx, q, email)
	u := &model.User{}
	err := row.Scan(&u.ID, &u.Email, &u.Fullname, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (*model.User, error) {
	q := `
		SELECT id, email, fullname, role, created_at, updated_at 
		FROM users 
		WHERE id=$1
	`
	row := r.db.QueryRow(ctx, q, id)
	u := &model.User{}
	err := row.Scan(&u.ID, &u.Email, &u.Fullname, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) GetAllUser(ctx context.Context, limit, offset int) ([]*model.User, error) {
	q := `
		SELECT id, email, fullname, role, created_at, updated_at 
		FROM users 
		ORDER BY created_at DESC 
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		u := &model.User{}
		if err := rows.Scan(&u.ID, &u.Email, &u.Fullname, &u.Role, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepo) UpdateRole(ctx context.Context, id string, newRole model.Role) error {
	q := `UPDATE users SET role=$1, updated_at=now() WHERE id=$2`
	cmd, err := r.db.Exec(ctx, q, newRole, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, user model.User) error {
	q := `UPDATE users SET fullname=$1, updated_at=now() WHERE id=$2`
	cmd, err := r.db.Exec(ctx, q, user.Fullname, user.ID)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, id string) error {
	q := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}
