package dto

type CreateUserRequest struct {
	Login     string `json:"login" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Sex       string `json:"sex" binding:"required"`
	HairColor string `json:"hair_color" binding:"required"`
}
