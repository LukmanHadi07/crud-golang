package repository

import (
	"errors"
	"mahasiswa-api-golang/data/request"
	"mahasiswa-api-golang/model"

	"gorm.io/gorm"
)

type StudentRepositoryImpl struct {
	Db *gorm.DB
}

func NewStudentRepositoryImp(Db *gorm.DB) StudentRepository {
	return &StudentRepositoryImpl{Db: Db}
}

// func Delete implements MahasiswaRepository interface
func (t *StudentRepositoryImpl) Delete(studentId int) {
	var students model.Students
	result := t.Db.Where("id = ?", studentId).Delete(&students)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (t *StudentRepositoryImpl) FindAll() []model.Students {
	var students []model.Students
	result := t.Db.Find(&students)
	if result.Error != nil {
		panic(result.Error)
	}
	return students
}

func (t *StudentRepositoryImpl) FindById(studentId int) (students model.Students, err error) {
	var student model.Students
	result := t.Db.Find(&student, studentId)
	if result != nil {
		return student, nil
	} else {
		return student, errors.New("Mahasiswa not found")
	}
}

func (t *StudentRepositoryImpl) Save(students model.Students) {
	result := t.Db.Create(&students)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (t *StudentRepositoryImpl) Update(students model.Students) {
	var updateStudent = request.UpdateStudentRequest{
		Name: students.Name,
		Nim:  students.Nim,
		Age:  students.Age,
	}

	result := t.Db.Model(&students).Updates(updateStudent)
	if result.Error != nil {
		panic(result.Error)
	}
}
