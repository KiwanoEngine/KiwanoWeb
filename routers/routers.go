package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KiwanoEngine/KiwanoWeb/modules/templates"
	"github.com/KiwanoEngine/KiwanoWeb/render"
)

func Setup(engine *gin.Engine) {
	engine.Delims("{{", "}}")
	engine.SetFuncMap(templates.FuncMap)
	engine.Static("/static", "./static_web/static")

	htmlRender := render.NewHTMLRender()
	htmlRender.TemplatesDir = "static_web/"
	htmlRender.Ext = ".html"
	engine.HTMLRender = htmlRender.Create()

	// Homepage
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", nil)
	})

	// Kiwano for golang
	engine.GET("/go", func(c *gin.Context) {
		c.HTML(http.StatusOK, "go/index", nil)
	})
	engine.GET("/kiwano", func(c *gin.Context) {
		c.HTML(http.StatusOK, "kiwano/index", nil)
	})
	engine.GET("/demos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "demos/index", nil)
	})
}
