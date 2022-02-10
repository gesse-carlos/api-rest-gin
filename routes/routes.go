package routes

import (
	"github.com/gesse-carlos/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.GetAll)
	r.GET("/student/:name", controllers.HandleParams)
	r.POST("/students", controllers.Create)
	r.GET("/students/:id", controllers.GetByID)
	r.DELETE("/students/:id", controllers.Remove)
	r.PATCH("/students/:id", controllers.Update)
	r.GET("/students/cpf/:cpf", controllers.GetByCPF)

	r.Run()
}