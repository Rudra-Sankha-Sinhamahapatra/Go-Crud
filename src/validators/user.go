package validators

import "github.com/go-playground/validator/v10"

type UserInput struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Pincode  int    `json:"pincode" validate:"required,min=100000,max=999999"`
	City     string `json:"city" validate:"required"`
	State    string `json:"state" validate:"required"`
	Country  string `json:"country" validate:"required"`
}

var validate = validator.New()

func (u *UserInput) Validate() error {
	return validate.Struct(u)
}
