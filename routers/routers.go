package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/nomango/KiwanoWeb/controllers"
	"github.com/nomango/KiwanoWeb/modules/templates"
)

func Setup(engine *gin.Engine) {
	engine.Delims("{{", "}}")
	engine.SetFuncMap(templates.FuncMap)
	engine.LoadHTMLFiles("./views/index.html")

	engine.GET("/", controllers.Home)
}
