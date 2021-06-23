package api

import (
	"encoding/json"
	"net/http"
)

type TodoApp struct{}

var _ ServerInterface = (*TodoApp)(nil)

func (t *TodoApp) FindTodos(w http.ResponseWriter, r *http.Request, params FindTodosParams) {
	var result []TodoItem

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (t *TodoApp) AddTodo(w http.ResponseWriter, r *http.Request) {
	var result TodoItem

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (t *TodoApp) DeleteTodoByID(w http.ResponseWriter, r *http.Request, id string) {
	var result TodoItem

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(result)
}

func (t *TodoApp) FindTodoByID(w http.ResponseWriter, r *http.Request, id string) {
	var result TodoItem

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (t *TodoApp) UpdateTodoByID(w http.ResponseWriter, r *http.Request, id string) {
	var result TodoItem

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
