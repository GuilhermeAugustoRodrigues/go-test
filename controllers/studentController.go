package controllers

import (
	"gin-api-rest/database"
	"gin-api-rest/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudents(ctx *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	ctx.JSON(http.StatusOK, students)
}

func GetStudentById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	var student models.Student
	database.DB.First(&student, id)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"not found": "student not found",
		})

		return
	}

	ctx.JSON(http.StatusOK, student)
}

func GetStudentByFilter(ctx *gin.Context) {
	filter := ctx.Params.ByName("filter")

	var students []models.Student
	database.DB.Where("students.name ilike ?", "%"+filter+"%").Or("students.cpf ilike ?", "%"+filter+"%").Or("students.rg ilike ?", "%"+filter+"%").Find(&students)

	ctx.JSON(http.StatusOK, students)
}

func CreateStudent(ctx *gin.Context) {
	var student models.Student
	err := ctx.ShouldBindJSON(&student)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		log.Panic(err)

		return
	}

	err = models.ValidateStudent(&student)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		log.Panic(err)

		return
	}

	database.DB.Create(&student)

	ctx.JSON(http.StatusOK, student)
}

func EditStudent(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")

	database.DB.First(&student, id)

	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		log.Panic(err)

		return
	}

	err = models.ValidateStudent(&student)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		log.Panic(err)

		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	ctx.JSON(http.StatusOK, student)
}

func DeleteStudent(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")

	database.DB.Delete(&student, id)

	ctx.JSON(http.StatusOK, student)
}
