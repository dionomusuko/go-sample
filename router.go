package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "public/assets")

	e.File("/", "public/index.html")
	e.File("/sign-up", "public/sign-up.html")
	e.POST("sign-up", handler.Sign-up)
	e.File("/sign-in", "public/sign-in")
	e.POST("/sing-in", handler.Sing-in)
	e.File("todos", "public/todos")

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(handler.Config))
	api.GET("/todos", handler.GetTodos)
	api.POST("/todos", handler.PostTodo)
	api.DELETE("todos/:id", handler.DeleteTodo)
	api.PUT("todos/:id/completed", handler.UpdateTodo)

	return e
}
