package Validations

type UserCreate struct {
	FirstName 	string `json:"FirstName" binding:"required"`
	LastName 	string `json:"LastName" binding:"required"`
	Email   	string `json:"Email" binding:"required"`
	Password   	string `json:"Password" binding:"required"`
	Phone   	string `json:"Phone" binding:"required"`
	Address 	*string `json:"Address"`
}

type UserUpdate struct {
	FirstName 	string `json:"FirstName" binding:"required"`
	LastName 	string `json:"LastName" binding:"required"`
	Phone   	string `json:"Phone" binding:"required"`
	Address 	*string `json:"Address"`
}

