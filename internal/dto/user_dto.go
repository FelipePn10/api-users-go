package dto

type CreateUserDto struct {
	BasicUserDto
	BasicUserInfoDto
}

type UpdateUserDto struct {
	BasicUserDto
	BasicUserInfoDto
}

type UpdateUserPasswordDto struct {
	BasicUserDto
	OldPassword string `json:"old_password" validate:"required,min=8,max=30,containsany=!@#$%*"`
}

type BasicUserDto struct {
	Name     string `json:"name" validate:"required,min=3,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30,containsany=!@#$%*"`
}

type BasicUserInfoDto struct {
	Age     int    `json:"age" validate:"required,min=14"`
	Gender  string `json:"gender" validate:"required,oneof=male female"`
	Address string `json:"address" validate:"required,min=10,max=100"`
	Phone   string `json:"phone" validate:"required,min=10,max=15"`
}
