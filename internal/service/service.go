package service

import (
    "github.com/GSlon/todoGO/internal/repository"
    "github.com/GSlon/todoGO/internal/entity"
)

type Authorization interface {
    CreateUser(user entity.User) (int, error)  // возвращает id созданного юзера
    GetUser(user entity.SignInUser) (int, error)     // возвращает id заданного юзера
}

type TodoItem interface {

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
