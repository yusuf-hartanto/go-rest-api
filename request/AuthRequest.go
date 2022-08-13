package request

type LoginReq struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
