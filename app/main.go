package main

import (
	"rakamin-final/internal/helper"
	"rakamin-final/internal/infrastructure/container"
	"rakamin-final/internal/infrastructure/mysql"

	rest "rakamin-final/internal/server/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	containerConf := container.InitContainer()
	defer mysql.CloseDatabaseConnection(containerConf.MySqlDB)

	app := fiber.New()
	app.Use(logger.New())

	rest.HTTPRouteInit(app, containerConf)

	addr := containerConf.Apps.Host + ":" + containerConf.Apps.HttpPort

	helper.Logger("main.go", helper.LoggerLevelFatal, app.Listen(addr).Error())
}
