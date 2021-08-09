package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/GSlon/todoGO/internal/entity"
)

type Authorization interface {
    CreateUser(user entity.User) (int, error)
    GetUser(user entity.SignInUser) (int, error)
}

type TodoItem interface {

}

type Repository struct {
    Authorization
    TodoItem
}

// dependency inversion (абстрагируемся от конкретной БД)
func NewRepository(db *sqlx.DB) *Repository {
    return &Repository{
            Authorization: NewAuthPostgres(db)}
}
