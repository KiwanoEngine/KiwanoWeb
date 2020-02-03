package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/KiwanoEngine/KiwanoWeb/controllers"
	"github.com/KiwanoEngine/KiwanoWeb/modules/templates"
)

func Setup(engine *gin.Engine) {
	engine.Delims("{{", "}}")
	engine.SetFuncMap(templates.FuncMap)
	engine.LoadHTMLFiles(
		"./views/index.html",
		"./views/go-index.html",
		"./views/go-package.html",
		"./views/go-demos-package.html",
	)
	engine.Static("/static", "./static")

	engine.GET("/", controllers.Home)
	engine.GET("/go", controllers.GoHome)
	engine.GET("/kiwano", controllers.GoPackageService)
	engine.GET("/demos", controllers.GoDemosPackageService)
}
