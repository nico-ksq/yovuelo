package user_registration

import (
	"database/sql"

	"yovuelo/db/user"
	"yovuelo/routes/requests"
)

type UserRegistration interface {
	Register(req requests.RegisterUserRequest) bool
}
type User struct {
	database *sql.DB
}

func New(db *sql.DB) User {
	return User{
		database: db,
	}
}

func (u User) Register(req requests.RegisterUserRequest) bool {
	return false
}

// transformToDBModel transforma una solicitud de registro en un modelo de base de datos.
func transformToDBModel(req requests.RegisterUserRequest) user.DBUser {
	return user.DBUser{
		Username:    req.Username,
		Password:    req.Password,
		Email:       req.Email,
		Country:     req.Country,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		City:        req.City,
		State:       req.State,
		PostalCode:  req.PostalCode,
	}
}
