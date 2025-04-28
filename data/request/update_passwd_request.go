package request

type UpdatePasswordRequest struct {
	OldPassword    string `json:"old_password" validate:"required"`
	Password       string `json:"password" validate:"required,strongpassword"`
	RetypePassword string `json:"retype_password" validate:"required,strongpassword"`
}
