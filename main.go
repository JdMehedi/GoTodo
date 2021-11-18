package main

import (
	"Todo/handler"
	"log"
	// "database/sql"

	"net/http"

	_ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

func main() {

	var schema = `
	CREATE TABLE IF NOT EXISTS tasks (
		id serial,
		title text,
		is_completed boolean,

		primary key(id)
	);`

	db, err := sqlx.Connect("postgres", "user=postgres password=Passw0rd dbname=todo sslmode=disable")
    if err != nil {
        log.Fatalln(err)
    }

	db.MustExec(schema)



	
	h:= handler.New(db)

	http.HandleFunc("/",h.Home) //Route

	http.HandleFunc("/todos/create",h.CreateTodo)
	http.HandleFunc("/todos/store",h.StoreTodo)
	http.HandleFunc("/todos/complete/",h.CompleteTodo)
	http.HandleFunc("/todos/edit/",h.EditTodo)
	http.HandleFunc("/todos/update/",h.UpdateTodo)
	http.HandleFunc("/todos/delete/",h.DeleteTodo)
	log.Println("Server starting ...........")

 if err := http.ListenAndServe("127.0.0.1:3000",nil); err !=nil{
	log.Fatal(err)
  }

}

