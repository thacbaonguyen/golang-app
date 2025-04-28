package request

type UserUpdateReq struct {
	FullName string `json:"full_name" validate:"required"`
}
