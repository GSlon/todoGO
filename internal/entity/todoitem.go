package entity

type TodoItem struct {
    Id          int    `json:"id"`
    Title       string `json:"title" binding: "required"`
    Description string `json:"description"`
}

type UpdateTodoItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
