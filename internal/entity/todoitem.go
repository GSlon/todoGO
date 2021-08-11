package entity

type TodoItem struct {
    Id          int    `json:"id" db:"id"`
    Title       string `json:"title" db:"title" binding: "required"`
    Description string `json:"description" db:"description"`
}

type UpdateTodoItem struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}
