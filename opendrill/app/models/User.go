package models

import "gopkg.in/mgo.v2/bson"
import "time"

// UserRole represents an user role
type UserRole int

// Gender represents an user gender
type Gender int

const (
	// Roles
	ROLE_ORGANIZATOR = 0
	ROLE_SUPER_ADMIN  = 1
	ROLE_REPORT      = 2
	ROLE_DESIGNER    = 3

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
	Name      string   		`json:"name"`
	LastName  string   		`json:"lastname"`
	UserName  string   		`json:"username"`
	Password  string   		`json:"password"`
	Age       int      		`json:"age"`
	Email     string   		`json:"email"`
	Role      UserRole 		`json:"role"`
	DNI       int      		`json:"dni"`
	Gender    Gender   		`json:"gender"`
	Address   string   		`json:"address"`
	Phone     string   		`json:"phone"`
	CellPhone string   		`json:"cellphone"`
	Avatar    string   		`json:"avatar"`
	Active    bool     		`json:"active"`
}

func RegisterUser(user User) (err error, organization2 Organization) {
	var new_organization Organization
	user.Id = bson.NewObjectId()
	_user := []User{user}
	_template := []Template{}
	_list_conatct := []ListContact{}

	organization2 = new_organization
	organization2.Users = _user
	organization2.ListContacts = _list_conatct
	organization2.Templates = _template
	organization2.Name = user.UserName
	organization2.CreatedAt = time.Now()
	organization2.ModifiedAt = time.Now()
	organization2.Id = bson.NewObjectId()
	if err := organizations.Insert(organization2); err != nil {
		return err, new_organization
	}
	return nil, organization2
}