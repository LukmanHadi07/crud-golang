package model

type Students struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(255)"`
	Nim  string `gorm:"type:varchar(255)"`
	Age  int    `gorm:"type:int"`
}
