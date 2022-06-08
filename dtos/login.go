package dtos

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenValidator struct {
	Token string `json:"token"`
}
