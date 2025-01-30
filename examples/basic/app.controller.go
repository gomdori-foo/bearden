package main

type AppController struct {
	appService AppService
}

// @Controller()
func NewAppController(appService AppService) *AppController {
	return &AppController{appService: appService}
}

// @Get()
func (c *AppController) GetHello() string {
	return c.appService.GetHello()
}
