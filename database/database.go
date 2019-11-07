package database

type Todo struct {
	ID int `json:"ID"`
	Todo string  `json:"Todo"`
}


func (todo *Todo) Validate() []string {
	errs := []string{}

	if len(todo.Todo) < 3 || len(todo.Todo) > 120 {
		errs = append(errs, "The todo field must be between 3-120 chars!")
	}

	return errs
}

type AllTodo []Todo

var Todos = AllTodo{
	{
		ID: 1,
		Todo: "Create a rest api",
	},
}