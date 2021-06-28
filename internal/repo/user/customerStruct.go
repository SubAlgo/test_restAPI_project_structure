package user

type Customer struct {
	ID       int    `json:"id"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
}
