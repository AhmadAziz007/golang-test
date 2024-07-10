package update

type UserUpdateRequest struct {
	Id       int    `validate:"required"`
	RoleId   int    `validate:"required"`
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
