package handler

import (
	"text/template"

	"github.com/jmoiron/sqlx"
)


type Todo struct{
	ID   int    `db:"id"  json:"id"`
	Task string `db:"title" json:"task"`
	IsComplete bool `db:"is_completed" json:"Is_completed"`
	
}

type Handler struct{
	templates *template.Template
	db *sqlx.DB
	
}


func New(db *sqlx.DB) *Handler {
	h:= &Handler{
		 db: db,
	}
	h.parseTemplate()
	return h
}

func (h *Handler) parseTemplate(){
   h.templates = template.Must(template.ParseFiles(
	   "templates/create-todo.html",
	   "templates/index-todo.html",
	   "templates/edit-todo.html",
	   ))
}




