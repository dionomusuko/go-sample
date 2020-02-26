package infrastructure

import (
	"github.com/dionomusuko/go-sample/infrastructure/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "public/assets")

	e.File("/", "public/index.html")
	e.File("/sign-up", "public/signup.html")
	e.POST("sign-up", handler.Signup)
	e.File("/sign-in", "public/login.html")
	e.POST("/sing-in", handler.Login)
	e.File("/tasks", "public/tasks.html")

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(handler.Config))
	api.GET("/tasks", handler.GetTasks)
	api.POST("/tasks", handler.AddTask)
	api.DELETE("tasks/:id", handler.DeleteTask)
	api.PUT("tasks/:id/completed", handler.UpdateTask)

	e.Logger.Fatal(e.Start(":3000"))
}
