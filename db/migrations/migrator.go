package migrations

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrator(db *sql.DB, dsn string) {
	// Inicializar la migración
	m, err := migrate.New(
		"file://db/migrations/sql",
		"mysql://"+dsn,
	)
	if err != nil {
		log.Fatalf("Error al inicializar la migración: %v", err)
	}

	// Aplicar las migraciones y registrar cada una
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error al aplicar la migración: %v", err)
	}

	// Obtener la lista de migraciones aplicadas
	migrations, _, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatalf("Error al obtener la versión de la migración: %v", err)
	}
	print(migrations)
	//// Registrar las migraciones en la tabla gorp_migration
	//for _, migration := range migrations {
	//	migrationName := fmt.Sprintf("Migration %d applied", migration)
	//	err = registerMigration(db, migrationName)
	//	if err != nil {
	//		log.Fatalf("Error al registrar la migración: %v", err)
	//	}
	//}
}

// Función para registrar una migración en la tabla gorp_migration
func registerMigration(db *sql.DB, migrationName string) error {
	insertMigration := `INSERT INTO gorp_migration (migration_name) VALUES (?);`
	_, err := db.Exec(insertMigration, migrationName)
	return err
}
