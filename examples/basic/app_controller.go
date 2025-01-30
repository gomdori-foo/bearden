package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	appService AppService
}

// @Controller()
func NewAppController(appService AppService) *AppController {
	return &AppController{appService: appService}
}

// @Get()
func (c *AppController) GetHello(ctx *gin.Context) {
	message := c.appService.GetHello()

	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

