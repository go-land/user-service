package dao

type UserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Job       string `json:"job"`
}

var UserInfo = map[string]UserDto{
	"maksym": UserDto{
		FirstName: "Maksym",
		LastName:  "Stepanenko",
		Job:       "EM",
	},
	"olesia": UserDto{
		FirstName: "Olesia",
		LastName:  "Stepanenko",
		Job:       "undefined",
	},
}
