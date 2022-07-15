package dtos

type RegisterIn struct {
	Email    string `json:"email" binding:"required,email,max=100"`
	Username string `json:"username" binding:"required,max=50"`
	Password string `json:"password" binding:"required"`
}
