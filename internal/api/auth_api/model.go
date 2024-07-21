package auth_api

import (
	"go-mercury/internal/data/gallery_db"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserData struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

func (d *UserData) SetRoles(roles []gallery_db.UserRole) {
	data := make([]string, 0)
	for _, r := range roles {
		data = append(data, string(r))
	}
	d.Roles = data
}
