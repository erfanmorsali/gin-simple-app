package dtos

type CategoryIn struct {
	Name string `json:"name" binding:"required,max=50"`
}
