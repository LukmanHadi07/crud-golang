package router

import (
	"mahasiswa-api-golang/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(studentController *controller.StudentController) *gin.Engine {
	router := gin.Default()
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	studentRouter := baseRouter.Group("/student")
	studentRouter.POST("", studentController.Create)
	studentRouter.GET("", studentController.FindAll)
	studentRouter.GET("/:studentId", studentController.FindById)
	studentRouter.PUT("/:studentId", studentController.Update)
	studentRouter.DELETE("/:studentId", studentController.Delete)
	return router
}
