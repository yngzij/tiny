package service

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RegistryRouters interface {
	RegistryRouters(e *echo.Group, middlewareFunc ...echo.MiddlewareFunc)
}

type EchoServer struct {
	e *echo.Echo
}

func (s EchoServer) Run() {
	log.Panic(s.e.Start("3000"))
}

func (s EchoServer) loadMiddlewares() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.RemoveTrailingSlash())
}

func (s EchoServer) registerRouter() {

}

func NewEchoServer() EchoServer {
	s := EchoServer{
		e: echo.New(),
	}
	s.loadMiddlewares()
	s.registerRouter()
	return s
}

func BindAndVerify(ctx echo.Context, data interface{}) error {
	if err := ctx.Bind(data); err != nil {
		return err
	}
	return nil
}
