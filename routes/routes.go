package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermemontanher/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.GetStudentById)
	r.DELETE("/students/:id", controllers.DeleteStudentById)
	r.PATCH("/students/:id", controllers.EditStudentById)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	r.Run()
}
