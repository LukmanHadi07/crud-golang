package service

import (
	"mahasiswa-api-golang/data/request"
	"mahasiswa-api-golang/data/response"
	"mahasiswa-api-golang/model"
	"mahasiswa-api-golang/repository"

	"github.com/go-playground/validator/v10"
)

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepository
	Validate          *validator.Validate
}

func NewStudentServiceImpl(studentRepository repository.StudentRepository, validate *validator.Validate) StudentService {
	return &StudentServiceImpl{StudentRepository: studentRepository, Validate: validate}
}

func (t *StudentServiceImpl) Create(students request.CreateStudentRequest) {
	err := t.Validate.Struct(students)
	if err != nil {
		panic(err)
	}
	studentModel := model.Students{
		Name: students.Name,
		Nim:  students.Nim,
		Age:  students.Age,
	}

	t.StudentRepository.Save(studentModel)
}

func (t *StudentServiceImpl) Delete(studentId int) {
	t.StudentRepository.Delete(studentId)
}

func (t *StudentServiceImpl) FindAll() []response.StudentResponse {
	result := t.StudentRepository.FindAll()
	var students []response.StudentResponse
	for _, value := range result {
		student := response.StudentResponse{
			Id:   value.Id,
			Name: value.Name,
			Nim:  value.Nim,
			Age:  value.Age,
		}
		students = append(students, student)
	}

	return students
}

func (t *StudentServiceImpl) FindById(studentId int) response.StudentResponse {
	result, err := t.StudentRepository.FindById(studentId)
	if err != nil {
		return response.StudentResponse{}
	}
	student := response.StudentResponse{
		Id:   result.Id,
		Name: result.Name,
		Nim:  result.Nim,
		Age:  result.Age,
	}

	return student
}

func (t *StudentServiceImpl) Update(students request.UpdateStudentRequest) {
	studentData, err := t.StudentRepository.FindById(students.Id)
	if err != nil {
		panic(err)
	}
	studentData.Name = students.Name
	studentData.Nim = students.Nim
	studentData.Age = students.Age
	t.StudentRepository.Update(studentData)
}
