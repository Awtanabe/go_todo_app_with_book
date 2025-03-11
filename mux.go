package main

import (
	"go_todo_app/handler"
	"go_todo_app/store"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func NewMux(db *gorm.DB) http.Handler {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json: charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	taskStore := store.NewTaskStore(db)
	v := validator.New()
	at := &handler.AddTask{Store: taskStore, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{Store: taskStore}
	mux.Get("/tasks", lt.ServeHTTP)

	return mux
}
