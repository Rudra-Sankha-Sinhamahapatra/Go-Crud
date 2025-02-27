package validators

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (l *LoginInput) Validate() error {
	return validate.Struct(l)
}
