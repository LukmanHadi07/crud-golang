package service

import (
	"mahasiswa-api-golang/data/request"
	"mahasiswa-api-golang/data/response"
)

type StudentService interface {
	Create(students request.CreateStudentRequest)
	Delete(students int)
	Update(students request.UpdateStudentRequest)
	FindById(studentId int) response.StudentResponse
	FindAll() []response.StudentResponse
}
