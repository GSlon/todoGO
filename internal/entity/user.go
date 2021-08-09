package entity

type User struct {
    Id       int    `json:"-" db:"id"`
    Name     string `json:"name" binding:"required"`
    Surname  string `json:"surname" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type SignInUser struct {
    Name     string `json:"name" binding:"required"`
    Password string `json:"password" binding:"required"`
}
