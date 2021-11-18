package handler

import (
	"net/http"
)

type IndexTodo struct{
   Todos []Todo
}


func (h *Handler) Home (rw http.ResponseWriter, r *http.Request) {
	todos := []Todo{}
    h.db.Select(&todos, "SELECT * FROM tasks")
	lt := IndexTodo{
		Todos: todos,
	}
	if err:= h.templates.ExecuteTemplate(rw,"index-todo.html", lt); err !=nil{
		http.Error(rw, err.Error(),http.StatusInternalServerError)
		return
	}
}
