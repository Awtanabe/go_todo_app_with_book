package main

import (
	"context"
	"go_todo_app/auth"
	"go_todo_app/config"
	"go_todo_app/handler"
	"go_todo_app/service"
	"go_todo_app/store"
	"net/http"

	"github.com/budougumi0617/go_todo_app/clock"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func NewMux(ctx context.Context, db *gorm.DB, cfg *config.Config) (http.Handler, error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json: charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	taskStore := store.NewTaskStore(db)
	v := validator.New()
	clocker := clock.RealClocker{}
	r := store.Repository{Clocker: clocker}
	ru := &handler.RegisterUser{
		// rはインターフェースなのでポインタ！！
		Service:   &service.RegisterUser{DB: db, Repo: &r},
		Validator: v,
	}

	mux.Post("/register", ru.ServeHTTP)

	rcli, err := store.NewKVS(ctx, cfg)
	if err != nil {
		return nil, err
	}

	jwter, err := auth.NewJWTer(rcli)

	if err != nil {
		return nil, err
	}

	at := &handler.AddTask{Store: taskStore, Validator: v}
	// mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{Store: taskStore}
	// mux.Get("/tasks", lt.ServeHTTP)

	mux.Route("/tasks", func(r chi.Router) {
		r.Use(handler.AuthMiddleware(jwter))
		r.Post("/", at.ServeHTTP)
		r.Get("/", lt.ServeHTTP)
	})

	l := &handler.Login{
		Service: &service.Login{
			DB:             db,
			Repo:           &r,
			TokenGenerator: jwter,
		},
		Validator: v,
	}

	mux.Post("/login", l.ServeHTTP)

	return mux, nil
}
