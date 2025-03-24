package request

type UpdateStudentRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required, min=3, max=255" json:"name"`
	Nim  string `validate:"required, min=3, max=255" json:"nim"`
	Age  int    `validate:"required, min=17" json:"age"`
}
