package handler

import (
	"go_todo_app/entity"
	"go_todo_app/store"
	"net/http"
)

type ListTask struct {
	Store *store.TaskStore
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := lt.Store.All()

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	resp := []task{}
	for _, t := range tasks {
		resp = append(resp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}

	RespondJSON(ctx, w, resp, http.StatusOK)
}
