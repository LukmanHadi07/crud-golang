package repository

import "mahasiswa-api-golang/model"

type StudentRepository interface {
	Save(student model.Students)
	Update(student model.Students)
	Delete(studentId int)
	FindById(studentId int) (students model.Students, err error)
	FindAll() []model.Students
}
