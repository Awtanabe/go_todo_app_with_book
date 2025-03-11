package handler

import (
	"encoding/json"
	"fmt"
	"go_todo_app/entity"
	"go_todo_app/store"
	"net/http"
	"time"

	"github.com/go-playground/validator"
)

type AddTask struct {
	// db接続
	Store     *store.TaskStore
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title string `json:"title" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	fmt.Printf("%+v", b)
	// バリデーションエラーs
	err := at.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	t := &entity.Task{
		Title:   b.Title,
		Status:  entity.TaskStatusTodo,
		Created: time.Now(),
	}

	id, err := store.Tasks.Add(t)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	resp := struct {
		ID int `json:"id"`
	}{ID: int(id)}
	RespondJSON(ctx, w, resp, http.StatusOK)
}
