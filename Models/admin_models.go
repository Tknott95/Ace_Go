package Models

var DbAdminUsers = map[string]AdminUser{}
var DbSessions = map[string]string{}

type AdminUser struct {
	AdminID  int
	Email    string
	Password []byte
}
