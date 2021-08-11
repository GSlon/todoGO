package service

import (
    "github.com/GSlon/todoGO/internal/repository"
    "github.com/GSlon/todoGO/internal/entity"
)

type Authorization interface {
    CreateUser(user entity.User) (int, error)       // возвращает id созданного юзера
    GetUser(user entity.SignInUser) (int, error)    // возвращает id заданного юзера
    GetUserById(id int) (entity.User, error)
}

type TodoItem interface {
    Create(userId int, item entity.TodoItem) (int, error)
    GetAllItems(userId int) ([]entity.TodoItem, error)
    GetItemById(id int) (entity.TodoItem, error)
    Delete(itemId int) error
    Update(itemId int, input entity.UpdateTodoItem) error
}

// бизнес-логика
type Service struct {
    Authorization
    TodoItem
}

func NewService(repo *repository.Repository) *Service {
    return &Service{
        Authorization: NewAuthService(repo.Authorization),
    }
}
