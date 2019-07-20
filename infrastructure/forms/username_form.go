package forms

// UsernameForm used to write a username as json
type UsernameForm struct {
	Username string `json:"username"`
}

// Valid form
func (form *UsernameForm) Valid() bool {
	return len(form.Username) > 8
}
