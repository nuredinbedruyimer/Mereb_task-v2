package routes

import (
	"Mereb-V2/controllers"

	"github.com/gin-gonic/gin"
)

func PersonRoutes(personRoutes *gin.Engine) *gin.Engine {

	personRoutes.POST("/persons", controllers.CreatePersonController)
	personRoutes.GET("/persons", controllers.GetPersonsController)
	personRoutes.GET("/persons/:id", controllers.GetPersonController)
	personRoutes.PUT("/persons/:id", controllers.UpdatePersonController)
	personRoutes.DELETE("/persons/:id", controllers.DeletePersonController)
	personRoutes.NoRoute(controllers.NotFoundHandler)

	return personRoutes

}
