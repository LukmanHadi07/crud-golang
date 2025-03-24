package controller

import (
	"mahasiswa-api-golang/data/request"
	"mahasiswa-api-golang/data/response"
	"mahasiswa-api-golang/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type StudentController struct {
	studentService service.StudentService
}

func NewStudentController(service service.StudentService) *StudentController {
	return &StudentController{
		studentService: service,
	}
}

func (controller *StudentController) Create(ctx *gin.Context) {
	log.Info().Msg("Create Student")
	createStudentRequest := request.CreateStudentRequest{}
	err := ctx.ShouldBindJSON(&createStudentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.studentService.Create(createStudentRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *StudentController) Delete(ctx *gin.Context) {
	log.Info().Msg("Delete Student")
	studentId := ctx.Param("studentId")
	id, err := strconv.Atoi(studentId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.studentService.Delete(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *StudentController) Update(ctx *gin.Context) {
	log.Info().Msg("Update Student")
	updateStudentRequest := request.UpdateStudentRequest{}
	err := ctx.ShouldBindJSON(&updateStudentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.studentService.Update(updateStudentRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *StudentController) FindById(ctx *gin.Context) {
	log.Info().Msg("Find By Id Student")
	studentId := ctx.Param("studentId")
	id, err := strconv.Atoi(studentId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	studentResponse := controller.studentService.FindById(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   studentResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *StudentController) FindAll(ctx *gin.Context) {
	log.Info().Msg("Find All Student")
	studentResponse := controller.studentService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   studentResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
