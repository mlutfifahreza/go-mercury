package gallery_db

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
)

func (r UserRole) IsValid() bool {
	switch r {
	case RoleAdmin, RoleCustomer:
		return true
	default:
		return false
	}
}
