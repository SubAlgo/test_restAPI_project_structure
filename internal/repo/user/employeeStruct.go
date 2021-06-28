package user

type EmployeeStruct struct {
	ID       int     `json:"id"`
	Role     string  `json:"role"`
	Name     string  `json:"name"`
	Lastname string  `json:"lastname"`
	Email    string  `json:"email"`
	Salary   float64 `json:"salary"`
}
