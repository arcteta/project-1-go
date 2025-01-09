package routes

import (
	"go-gomanager/handler"

	"github.com/labstack/echo/v4"
)

type RouterConfig struct {
	App *echo.Echo
	DepartmentController *handler
}