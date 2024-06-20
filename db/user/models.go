package user

// DBUser representa un usuario en la base de datos.
type DBUser struct {
	Username    string
	Password    string
	Email       string
	Country     string
	FirstName   string
	LastName    string
	DateOfBirth string
	PhoneNumber string
	Address     string
	City        string
	State       string
	PostalCode  string
}
