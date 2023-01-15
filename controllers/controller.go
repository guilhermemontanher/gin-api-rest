package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermemontanher/api-go-gin/database"
	"github.com/guilhermemontanher/api-go-gin/models"
)

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func GetStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student Not found"})

		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted successfully"})
}

func EditStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, &student)
}

func GetStudentByCpf(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student Not found"})

		return
	}

	c.JSON(http.StatusOK, student)
}
