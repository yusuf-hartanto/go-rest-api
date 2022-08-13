package response

import (
	"rest-api/models"
)

type LoginResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"user_name"`
	Password  string `json:"password"`
	CreatedBy int64  `json:"created_by"`
	CreatedAt string `json:"created_at"`
	UpdatedBy int64  `json:"updated_by"`
	UpdatedAt string `json:"updated_at"`
	Token     string `json:"token"`
}

func (r LoginResponse) LoginTransform(m models.User, Token string) LoginResponse {
	var res LoginResponse
	res.ID = m.ID
	res.Name = m.Name
	res.Username = m.Username
	res.Password = m.Password
	res.CreatedBy = m.CreatedBy
	res.CreatedAt = m.CreatedAt
	res.UpdatedBy = m.UpdatedBy
	res.UpdatedAt = m.UpdatedAt
	res.Token = Token
	return res
}
