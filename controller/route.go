package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB) {
	e.GET("/", HelloHandler)

	authorController := NewAuthorController(db)
	e.GET("/authors", authorController.GetAll)
}
