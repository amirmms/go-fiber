package dtos

type CreateUserDto struct {
	Name   string `json:"name" validate:"required,min=4"`
	Family string `json:"family" validate:"required,min=4"`
	Mobile string `json:"mobile" validate:"required,max=11,min=11,alphanum"`
}

type GetUserDto struct {
	Id int `json:"id" validate:"required,number"`
}

type UpdateUserDto struct {
	Id     int    `json:"id" validate:"required,number"`
	Name   string `json:"name" validate:"min=4"`
	Family string `json:"family" validate:"min=4"`
	Mobile string `json:"mobile" validate:"max=11,min=11,alphanum"`
}
