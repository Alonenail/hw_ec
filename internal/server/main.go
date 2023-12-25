package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Server struct {
	*echo.Echo
}

func NewServer() *Server {
	server := Server{echo.New()}

	server.AddMiddleware(middleware.Logger())
	server.AddMiddleware(middleware.Recover())

	server.Logger.SetLevel(log.DEBUG)

	return &server
}

func (s *Server) Serve() {
	s.Logger.Fatal(s.Start(":1323"))
}

func (s *Server) AddMiddleware(mw echo.MiddlewareFunc) {
	s.Use(mw)
}
