package handler

import (
	"encoding/json"
	"go_todo_app/entity"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

type RegisterUser struct {
	Service   RegisterUserService
	Validator *validator.Validate
}

func (ru *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ユーザーの情報 パラメーター
	var b struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}

	// デコードだから jsonを構造体に変更
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// 構造体のチェック
	if err := ru.Validator.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	log.Print("==通った")
	u, err := ru.Service.RegisterUser(ctx, b.Name, b.Password, b.Role)

	log.Print("==err", err)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	resp := struct {
		ID entity.UserID `json:"id"`
	}{ID: u.ID}
	RespondJSON(ctx, w, resp, http.StatusOK)
}
