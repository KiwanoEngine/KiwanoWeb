package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GoHome(c *gin.Context) {
	c.HTML(http.StatusOK, "go-index.html", nil)
}

func GoPackageService(c *gin.Context) {
	c.HTML(http.StatusOK, "go-package.html", nil)
}

func GoDemosPackageService(c *gin.Context) {
	c.HTML(http.StatusOK, "go-demos-package.html", nil)
}
