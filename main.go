package main

import (
	"net/http"

	"github.com/nomango/KiwanoWeb/modules/settings"
	"github.com/nomango/KiwanoWeb/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	settings.Init()
	routers.Setup(engine)

	server := http.Server{Handler: engine, Addr: ":" + settings.Port}
	server.ListenAndServe()
}
