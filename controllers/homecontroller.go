package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (this HomeController) ServeHomePage(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)
}

func (this HomeController) ServeAboutPage(c *gin.Context) {

}