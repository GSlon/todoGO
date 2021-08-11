package service

import (
    "github.com/GSlon/todoGO/internal/repository"
    "github.com/GSlon/todoGO/internal/entity"
)

type TodoItemService struct {
    repo repository.TodoItem
    userRepo repository.Authorization
}

func (s *TodoItemService) Create(userId int, item entity.TodoItem) (int, error) {
    _, err := s.userRepo.GetUserById(userId)
    if err != nil {
        // user not found
        return 0, err
    }

    return s.repo.Create(userId, item)
}

func (s *TodoItemService) GetAllItems(userId int) ([]entity.TodoItem, error) {
    _, err := s.userRepo.GetUserById(userId)
    if err != nil {
        // user not found
        return []entity.TodoItem{}, err
    }

    return s.repo.GetAllItems(userId)
}

func (s *TodoItemService) Delete(itemId int) error {
    return s.repo.Delete(itemId)
}

func (s *TodoItemService) Update(itemId int, input entity.UpdateTodoItem) error {
    return s.repo.Update(itemId, input)
}

func (s *TodoItemService) GetItemById(itemId int) (entity.TodoItem, error) {
    return s.repo.GetItemById(itemId)
}

func NewTodoItemService(repo repository.TodoItem, userRepo repository.Authorization) *TodoItemService {
    return &TodoItemService{
        repo: repo,
        userRepo: userRepo,
    }
}

