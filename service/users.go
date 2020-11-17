package service

import (
	"tidy/resp"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
}

func (u UsersController) RegistryRouters(e *echo.Group, middlewareFunc ...echo.MiddlewareFunc) {
	users := e.Group("/users")
	users.POST("/login", login)
}

func login(ctx echo.Context) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := BindAndVerify(ctx, req); err != nil {
		return resp.NewParseError(ctx)
	}
	return resp.Success(ctx)
}
