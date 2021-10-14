package models

// swagger:parameters auth signIn
// swagger:parameters auth signUp
type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
