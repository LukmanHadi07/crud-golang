package request

type CreateStudentRequest struct {
	Name string `validate:"required,min=3,max=255" json:"name"`
	Nim  string `validate:"required,min=3,max=255" json:"nim"`
	Age  int    `validate:"required,gte=17" json:"age"`
}
