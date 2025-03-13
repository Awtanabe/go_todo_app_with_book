package handler

import (
	"encoding/json"
	"fmt"
	"go_todo_app/store"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

type AddTask struct {
	// db接続
	Store     *store.TaskStore
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("xxx")
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
		log.Print("aaa")
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	task, err := NewTaskService(at.Store).AddTask(ctx, b.Title)
	if err != nil {
		log.Print("bbb")

		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	resp := struct {
		ID int `json:"id"`
	}{ID: int(task.ID)}
	RespondJSON(ctx, w, resp, http.StatusOK)
}
