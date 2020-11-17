package resp

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomError int

var (
	OperateSuccess CustomError = 0

	DatabaseError CustomError = 1
	NetworkError  CustomError = 2
	ParseError    CustomError = 3
	UserNameError CustomError = 4
	PasswordError CustomError = 5
)

type resp struct {
	Code CustomError `json:"code"`
	Msg  string      `json:"msg"`
}

func NewParseError(ctx echo.Context) error {
	return ctx.JSON(http.StatusBadRequest, resp{
		Code: ParseError,
		Msg:  "参数解析错误",
	})
}

func (r resp) Error() string {
	return r.Msg
}

func Success(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, resp{
		Code: OperateSuccess,
		Msg:  "success",
	})
}
