package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct{}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (c *HelloController) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func (c *HelloController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "foo",
	})
}
