package main

import (
	"mahasiswa-api-golang/config"
	"mahasiswa-api-golang/controller"
	"mahasiswa-api-golang/model"
	"mahasiswa-api-golang/repository"
	"mahasiswa-api-golang/router"
	"mahasiswa-api-golang/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting Application")
	db := config.DatabaseConnection()
	validate := validator.New()
	db.Table("student").AutoMigrate(&model.Students{})

	studentRepository := repository.NewStudentRepositoryImp(db)
	studentService := service.NewStudentServiceImpl(studentRepository, validate)
	studentController := controller.NewStudentController(studentService)

	routes := router.NewRouter(studentController)
	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
