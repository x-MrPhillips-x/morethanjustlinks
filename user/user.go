package user

const (
	UUIDKEY = "uuid"
)

type User struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Psword   string `json:"psword`
	Verified bool   `json:"verified"`
}

type Auth struct {
	Name   string `json:"name"`
	Psword string `json:"psword"`
}
