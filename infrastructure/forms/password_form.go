package forms

// PasswordForm used to write a password as json
type PasswordForm struct {
	Password string `json:"password"`
}

// Valid form
func (form *PasswordForm) Valid() bool {
	return len(form.Password) > 8
}
