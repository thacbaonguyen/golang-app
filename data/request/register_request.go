package request

type RegisterRequest struct {
	Username       string `json:"username" validate:"required,username,min=3,max=50"`
	Email          string `json:"email" validate:"email,required"`
	Password       string `json:"password" validate:"required,strongpassword"`
	RetypePassword string `json:"retype_password" validate:"required,strongpassword"`
	FullName       string `json:"full_name" validate:"required"`
}
