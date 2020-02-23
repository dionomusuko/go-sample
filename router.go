package main

import (
	"github.com/dionomusuko/sample-app/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "public/assets")

	e.File("/", "public/index.html")
	e.File("/sign-up", "public/sign-up.html")
	e.POST("sign-up", handler.Signup)
	e.File("/sign-in", "public/sign-in")
	e.POST("/sing-in", handler.Login)
	e.File("todos", "public/todos")

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(handler.Config))
	api.GET("/todos", handler.GetTodos)
	api.POST("/todos", handler.AddTodo)
	api.DELETE("todos/:id", handler.DeleteTodo)
	api.PUT("todos/:id/completed", handler.UpdateTodo)

	return e
}
