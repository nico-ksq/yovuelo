package user

import "database/sql"

type Registration interface {
	Save(tx *sql.Tx) error
}

func (u *DBUser) Save(tx *sql.Tx) error {
	// Preparamos la sentencia SQL para insertar el usuario.
	stmt, err := tx.Prepare(`
		INSERT INTO users (
			username, password, email, country, first_name, last_name,
			date_of_birth, phone_number, address, city, state, postal_code
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Ejecutamos la sentencia SQL con los valores del usuario.
	_, err = stmt.Exec(
		u.Username, u.Password, u.Email, u.Country, u.FirstName, u.LastName,
		u.DateOfBirth, u.PhoneNumber, u.Address, u.City, u.State, u.PostalCode,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}
