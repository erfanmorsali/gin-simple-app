package dtos

type PostIn struct {
	Title       string `json:"title" binding:"required,max=50"`
	Description string `json:"description" binding:"required,max=100"`
	UserID      uint   `json:"user_id" binding:"required"`
	CategoryIDS []uint `json:"category_ids"`
}
