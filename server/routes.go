package server

import (
	"back_agendamento/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		empresa := main.Group("/empresa")
		{
			empresa.POST("/", controllers.CriarEmpresa)
			empresa.GET("/:id", controllers.BuscarEmpresa)
		}
		cliente := main.Group("/cliente")
		{
			cliente.POST("/", controllers.CriarCliente)
			cliente.GET("/:id", controllers.BuscarCliente)
		}
	}

	return router
}
