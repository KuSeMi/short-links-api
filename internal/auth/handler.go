package auth

import (
	"demo/configs"
	"demo/pkg/req"
	"demo/pkg/res"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	Config *configs.Config
}

type AuthHandler struct {
	Config *configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		resp := LoginResponse{
			Token: "123",
		}
		res.Json(w, resp, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
	}
}
