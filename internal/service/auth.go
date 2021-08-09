package service

import (
    _ "github.com/GSlon/todoGO/internal/repository"
    "github.com/GSlon/todoGO/internal/entity"
    _ "github.com/sirupsen/logrus"
    "crypto/sha1"
    "encoding/hex"
)

type AuthService struct {
    repoauth Authorization
}

// передаем дальше, к уровню логики работы с БД (repository)
func (s *AuthService) CreateUser(user entity.User) (int, error) {
    user.Password = s.generatePasswordHash(user.Password)
    return s.repoauth.CreateUser(user)
}

func (s *AuthService) GetUser(user entity.SignInUser) (int, error) {
    user.Password = s.generatePasswordHash(user.Password)
    //logrus.Info(user.Password)
    return s.repoauth.GetUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
    hasher := sha1.New()
    hasher.Write([]byte(password))
    sha1_hash := hex.EncodeToString(hasher.Sum(nil))
    return sha1_hash
}

func NewAuthService(repoauth Authorization) *AuthService {
    return &AuthService{repoauth: repoauth}
}
