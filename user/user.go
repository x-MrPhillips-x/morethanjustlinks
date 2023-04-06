package user

const (
	USERKEY = "user"
)

type User struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Verified bool   `json:"verified"`
}

type Auth struct {
	Username string `json:"username"`
	Psword   string `json:"psword"`
}
