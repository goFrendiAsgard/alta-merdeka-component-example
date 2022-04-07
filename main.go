package main

import (
	"merdeka/config"
	"merdeka/controller"
	"merdeka/model"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.NewCfg()
	e := echo.New()
	db, err := gorm.Open(mysql.Open(cfg.DBConnectionString), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Seharusnya tabelnya dibuat dulu di db
	db.AutoMigrate(&model.Author{})

	controller.RegisterRoute(e, db)
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
