package user_registration

import (
	"context"
	"database/sql"
	"log"

	"yovuelo/db/user"
	"yovuelo/routes/requests"
)

type UserRegistration interface {
	RegisterUser(ctx context.Context, req requests.RegisterUserRequest) error
}

type User struct {
	database *sql.DB
}

func New(db *sql.DB) User {
	return User{
		database: db,
	}
}

func (u User) RegisterUser(ctx context.Context, req requests.RegisterUserRequest) error {
	dbuser := transformToDBModel(req)
	// Begin a transaction
	tx, err := u.database.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = dbuser.Save(tx)
	if err != nil {
		// log error
		return err
	}
	return nil
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
