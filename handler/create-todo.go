package handler

import (
	"net/http"
)

type formData struct{
		Todo Todo
		IsComplete bool
		Errors map[string]string

}

func (h *Handler) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	Errors := map[string]string{
	}
	todo := Todo{}
	h.loadCreatedTodoForm(rw,todo,Errors)
	
}

func (h *Handler) StoreTodo(rw http.ResponseWriter, r *http.Request){
if err :=r.ParseForm(); err !=nil{
	http.Error(rw, err.Error(),http.StatusInternalServerError)
	return
     }

	 task := r.FormValue("task")
	 todo :=Todo{
		 Task:task,
	 }
	 if task == ""{

		Errors := map[string]string{
			"Task":"This filed cannot be null",
		}
		h.loadCreatedTodoForm(rw,todo,Errors)
		   return
	   }

	if len(task) <3 {

		Errors := map[string]string{
			"Task":"This filed must be greater than or equals 3",
		}
		h.loadCreatedTodoForm(rw,todo,Errors)	
	} 

	const insertTodo = `INSERT INTO tasks(title, is_completed) VALUES($1,$2);`


	res := h.db.MustExec(insertTodo, task, false)

	if ok, err:= res.RowsAffected(); err != nil || ok == 0 {
		http.Error(rw, err.Error(),http.StatusInternalServerError)

		return
	}
	
	 http.Redirect(rw,r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) CompleteTodo(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/todos/complete/"):]
	

	if id == "" {
		http.Redirect(rw,r,"/todos/create/", http.StatusTemporaryRedirect)
		return
	}

		const completedTodo = `UPDATE tasks SET is_completed = true WHERE id = $1`
		res:= h.db.MustExec( completedTodo, id )

		if ok, err:= res.RowsAffected(); err != nil || ok == 0 {
			http.Error(rw, err.Error(),http.StatusInternalServerError)
	
			return
		}

	http.Redirect(rw,r, "/", http.StatusTemporaryRedirect)
}



func (h *Handler) EditTodo(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/todos/edit/"):]

	if id == "" {
		http.Error(rw, "invalid ", http.StatusTemporaryRedirect)
		return
	}


	const getTodo = `SELECT * FROM tasks WHERE id = $1`
	var todo Todo
	h.db.Get(&todo, getTodo, id )

	if todo.ID == 0 {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}

	h.loadUpdateTodoForm(rw,todo,map[string]string{})
}



func (h *Handler) UpdateTodo(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/todos/update/"):]

	if id == "" {
		http.Error(rw, "invalid update", http.StatusTemporaryRedirect)
		return
	}

	const getTodo = `SELECT * FROM tasks WHERE id = $1`
	var todo Todo
	h.db.Get(&todo, getTodo, id )

	if todo.ID == 0 {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}

	if err :=r.ParseForm(); err !=nil{
		http.Error(rw, err.Error(),http.StatusInternalServerError)
		 }
		 task := r.FormValue("task") 
		// todo.Task = task

		 if task == ""{
			Errors := map[string]string{
				"Task":"This filed cannot be null",
			}
			h.loadUpdateTodoForm(rw,todo,Errors)
			   return
		   }
	
		if len(task) <3 {
	
			Errors := map[string]string{
				"Task":"This filed must be greater than or equals 3",
			}
			h.loadUpdateTodoForm(rw,todo,Errors)	
		} 
	
		const completedTodo = `UPDATE tasks SET title = $2 WHERE id = $1`
		res:= h.db.MustExec( completedTodo, id, task)

		if ok, err:= res.RowsAffected(); err != nil || ok == 0 {
			http.Error(rw, err.Error(),http.StatusInternalServerError)
	
			return
		}

	http.Redirect(rw,r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) DeleteTodo(rw http.ResponseWriter, r *http.Request) {
	
	id := r.URL.Path[len("/todos/delete/"):]

	if id == "" {
		http.Error(rw, "invalid update", http.StatusTemporaryRedirect)
		return
	}

	const getTodo = `SELECT * FROM tasks WHERE id = $1`
	var todo Todo
	h.db.Get(&todo, getTodo, id )

	if todo.ID == 0 {
		http.Error(rw, "invalid URL", http.StatusInternalServerError)
		return
	}

	const deleteTodo =`DELETE FROM tasks WHERE id =$1`

	res:= h.db.MustExec( deleteTodo, id)

	if ok, err:= res.RowsAffected(); err != nil || ok == 0 {
		http.Error(rw, err.Error(),http.StatusInternalServerError)

		return
	}


	http.Redirect(rw,r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) loadCreatedTodoForm(rw http.ResponseWriter, todo Todo, errs map[string]string){

	form:=formData{
			Todo: todo,
			Errors: errs,
		}

		if err:= h.templates.ExecuteTemplate(rw,"create-todo.html", form); err !=nil{
			http.Error(rw, err.Error(),http.StatusInternalServerError)
			return
		}

}

func (h *Handler) loadUpdateTodoForm(rw http.ResponseWriter, todo Todo, errs map[string]string){

	form:=formData{
			Todo: todo,
			Errors: errs,
		}

		if err:= h.templates.ExecuteTemplate(rw,"edit-todo.html", form); err !=nil{
			http.Error(rw, err.Error(),http.StatusInternalServerError)
			return
		}

}


