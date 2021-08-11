package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/GSlon/todoGO/internal/entity"
)

type Authorization interface {
    CreateUser(user entity.User) (int, error)
    GetUser(user entity.SignInUser) (int, error)
    GetUserById(id int) (entity.User, error)
}

type TodoItem interface {
    Create(userId int, item entity.TodoItem) (int, error)
    GetAllItems(userId int) ([]entity.TodoItem, error)
    GetItemById(id int) (entity.TodoItem, error)
    Delete(itemId int) error
    Update(itemId int, input entity.UpdateTodoItem) error
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
