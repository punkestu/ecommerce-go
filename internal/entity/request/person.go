package request

type PersonRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PersonLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
