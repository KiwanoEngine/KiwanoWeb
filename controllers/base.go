package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"now": time.Now(),
	})
}

func GoPackageService(c *gin.Context) {
	c.HTML(http.StatusOK, "go-package.html", nil)
}

func GoDemosPackageService(c *gin.Context) {
	c.HTML(http.StatusOK, "go-demos-package.html", nil)
}
