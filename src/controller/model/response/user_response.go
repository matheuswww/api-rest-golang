package response

type UserResponse struct {
	Email string `json:"email"`
	Name string `json:"name"`
	Age uint8	`json:"age"`
	Id uint `json:"id"`
}