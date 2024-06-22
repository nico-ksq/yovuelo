package user

import "database/sql"

type Registration interface {
	Save(tx *sql.Tx) error
}

// Save guarda un usuario en la base de datos.
func (u DBUser) Save(tx *sql.Tx) error {
	return nil
}
