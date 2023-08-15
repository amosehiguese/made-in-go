package models

type Role struct {
	RoleName    string
	Permissions []string
}

// roles
var (
	AdminRole Role = Role{
		RoleName: "admin",
		Permissions: []string{
			PermCreateStock,
			PermUpdateStock,
			PermDeleteStock,

			PermReadUser,
			PermDeleteUser,
		},
	}
	UserRole Role = Role{
		RoleName:    "user",
		Permissions: []string{},
	}
)

func AssignDefaultRole(u *User) *User {
	u.Role = UserRole
	return u
}

func AssignAdminRole(u *User) *User {
	u.Role = AdminRole
	return u
}

// permissions
const (
	// stocks permissions
	PermCreateStock = "create:stock"
	PermUpdateStock = "update:stock"
	PermDeleteStock = "delete:stock"

	// user permissions
	PermReadUser   = "read:user"
	PermDeleteUser = "delete:user"
)

func IsAdmin(u *User) bool {
	return u.Role.RoleName == AdminRole.RoleName
}
