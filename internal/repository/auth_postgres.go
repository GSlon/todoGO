package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/GSlon/todoGO/internal/entity"
    "fmt"
    "strconv"
)

type AuthPostgres struct {
    db *sqlx.DB
}

func (p *AuthPostgres) CreateUser(user entity.User) (int, error) {
    var id int

    query := fmt.Sprintf(`INSERT INTO %s (name, surname, password_hash)
                        values ($1, $2, $3) RETURNING id`, userTable)
    row := p.db.QueryRow(query,  user.Name, user.Surname, user.Password)
    if err := row.Scan(&id); err != nil {
        return 0, err
    }

    return id, nil
}

func (p *AuthPostgres) GetUser(user entity.SignInUser) (int, error) {
    var id int

    query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1 AND password_hash=$2",
                         userTable)

    //row := p.db.QueryRow(query, user.Name, user.Password)
    if err := p.db.Get(&id, query, user.Name, user.Password); err != nil {
        return 0, err
    }

    return id, nil
}

func (p *AuthPostgres) GetUserById(id int) (entity.User, error) {
    var user entity.User
    query := fmt.Sprintf("SELECT name, surname, password_hash FROM %s WHERE id=$1",
                         userTable)

    //row := p.db.QueryRow(query, id)
    if err := p.db.Get(&user, query, strconv.Itoa(id)); err != nil {
        return user, err
    }

    return user, nil
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
    return &AuthPostgres{db: db}
}
