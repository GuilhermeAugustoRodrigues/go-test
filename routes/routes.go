package routes

import (
	"gin-api-rest/controllers"
	"gin-api-rest/database"
	"gin-api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := GetRoutesSetup()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.Run(":8080")
}

func GetRoutesSetup() *gin.Engine {
	r := gin.Default()
	r.GET("/alunos", controllers.GetStudents)
	r.GET("/alunos/:id", controllers.GetStudentById)
	r.GET("/alunos/search/:filter", controllers.GetStudentByFilter)
	r.POST("/alunos", controllers.CreateStudent)
	r.PATCH("/alunos/:id", controllers.EditStudent)
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	r.GET("/:name", hello)
	r.GET("/index", showIndexPage)
	r.NoRoute(showNotFoundPage)

	return r
}

func hello(ctx *gin.Context) {
	name := ctx.Params.ByName("name")
	ctx.JSON(http.StatusOK, gin.H{
		"API says": "Hello, " + name,
	})
}

func showIndexPage(ctx *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func showNotFoundPage(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", nil)
}
