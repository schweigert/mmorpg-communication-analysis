package forms

// CredentialForm is used to mock user and pass over network
type CredentialForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Valid form
func (form *CredentialForm) Valid() bool {
	usernameForm := &UsernameForm{Username: form.Username}
	passwordForm := &PasswordForm{Password: form.Password}
	return usernameForm.Valid() && passwordForm.Valid()
}
