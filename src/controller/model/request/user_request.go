package request

type UserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=150,containsany=!@#"`
	Name string	`json:"name" binding:"required,min=4,max=50"`
	Age uint8 `json:"age" binding:"required,min=1,max=127"`
}

type UserUpdateRequest struct {
	Name string	`json:"name" binding:"omitempty,min=4,max=50"`
	Age uint8 `json:"age" binding:"omitempty,min=1,max=127"`
}