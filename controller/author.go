package controller

import (
	"fmt"
	"merdeka/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthorController struct {
	Db *gorm.DB
}

func (ac AuthorController) GetAll(c echo.Context) error {
	authors := []model.Author{}
	tx := ac.Db.Find(&authors)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return c.JSON(http.StatusInternalServerError, "cannot fetch authors")
	}
	return c.JSON(http.StatusOK, authors)
}

func NewAuthorController(db *gorm.DB) AuthorController {
	return AuthorController{
		Db: db,
	}
}
