package response

type UserResponse struct {
	Id       int    `json:"id"`
	RoleId   int    `json:"roleId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
