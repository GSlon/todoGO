package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/GSlon/todoGO/internal/entity"
    "fmt"
)

type TodoItemPostgres struct {
    db *sqlx.DB
}

func (p *TodoItemPostgres) Create(userId int, item entity.UpdateTodoItem) (int, error) {
    query := fmt.Sprintf("INSERT INTO %s (title, description, user_id) values ($1, $2, $3) RETURNING id",
                         todoitemsTable)

    var itemId int
    row := p.db.QueryRow(query, item.Title, item.Description, userId)
    err := row.Scan(&itemId)
    if err != nil {
        return 0, err
    }

    return itemId, nil
}

func (p *TodoItemPostgres) GetAllItems(userId int) ([]entity.TodoItem, error) {
    var items []entity.TodoItem
    query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE user_id=$1", todoitemsTable)

    if err := p.db.Select(&items, query, userId); err != nil {
        return items, err
    }

    return items, nil
}

func (p *TodoItemPostgres) GetItemById(id int) (entity.TodoItem, error) {
    var item entity.TodoItem
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id = $1", todoitemsTable)

	if err := p.db.Get(&item, query, id); err != nil {
		return item, err
	}

	return item, nil
}

func (p *TodoItemPostgres) Delete(itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", todoitemsTable)
	_, err := p.db.Exec(query, itemId)
	return err
}

func (p *TodoItemPostgres) Update(itemId int, input entity.UpdateTodoItem) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 WHERE id=$3", todoitemsTable)
	_, err := p.db.Exec(query, input.Title, input.Description, itemId)
	return err
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
    return &TodoItemPostgres {
        db: db,
    }
}

