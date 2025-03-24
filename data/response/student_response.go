package response

type StudentResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Nim  string `json:"nim"`
	Age  int    `json:"age"`
}
