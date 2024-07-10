package implement

import (
	"azizdev/helper"
	"azizdev/model/domain"
	"azizdev/repository"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct{}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQl := `insert into user (role_id, name, email, password) values (?, ?, ?)`
	result, err := tx.ExecContext(ctx, SQl, user.RoleId, user.Name, user.Email, user.Password)
	helper.PanicIfError(err)

	userId, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(userId)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `update user set role_id = ?, name = ?, email = ?, password = ? where id = ?`
	_, err := tx.QueryContext(ctx, SQL, user.RoleId, user.Name, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := `delete from user where id = ?`
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
	SQL := `select id, role_id, name, email, password from user where id = ?
            order by id, role_id, name, email, password asc`
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.RoleId, &user.Name, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := `select id, role_id, name, email, password from user
            order by id, role_id, name, email, password asc`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.RoleId, &user.Name, &user.Email, &user.Password)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
