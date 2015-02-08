package models

// UserRole represents an user role
type UserRole int

// Gender represents an user gender
type Gender int

const (
	// Roles
	RoleOrganizator = 0
	RoleSuperAdmin  = 1
	RoleReport      = 2
	RoleDesigner    = 3

	// Genders
	Male   = 1
	Female = 2
	Other  = 3
)

var (
	AvailableLanguages = []string{"en", "es"}
)

type User struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `json:"name"`
	LastName  string        `json:"lastname"`
	UserName  string        `json:"username"`
	Password  string        `json:"password"`
	Age       int           `json:"age"`
	Email     string        `json:"email"`
	Role      UserRole      `json:"type"`
	DNI       int           `json:"dni"`
	Gender    Gender        `json:"gender"`
	Address   string        `json:"address"`
	Phone     string        `json:"phone"`
	CellPhone string        `json:"cellphone"`
	Avatar    string        `json:"avatar"`
	Active    bool          `json:"active"`
}
