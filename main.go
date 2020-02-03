package main

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/KiwanoEngine/KiwanoWeb/modules/settings"
	"github.com/KiwanoEngine/KiwanoWeb/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	settings.Init()
	routers.Setup(engine)

	server := http.Server{Handler: engine, Addr: ":" + settings.Port}
	if err := server.ListenAndServe(); err != nil {
		logrus.Panicln("Start Gin server failed", err)
	}
}
