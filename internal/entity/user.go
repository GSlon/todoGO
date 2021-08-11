package entity

type User struct {
    Id       int    `json:"-" db:"id"`
    Name     string `json:"name" db:"name" binding:"required"`
    Surname  string `json:"surname" db:"surname" binding:"required"`
    Password string `json:"password" db:"password_hash" binding:"required"`
}

type SignInUser struct {
    Name     string `json:"name" db:"name" binding:"required"`
    Password string `json:"password" db:"password_hash" binding:"required"`
}
