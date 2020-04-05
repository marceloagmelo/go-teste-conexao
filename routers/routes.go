package routers

import (
	"github.com/labstack/echo"
	"github.com/marceloagmelo/go-teste-conexao/controllers"
)

//App é uma instância de echo
var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", controllers.Home)

	api := App.Group("/v1")
	api.POST("/conectar", controllers.Conectar)
}
