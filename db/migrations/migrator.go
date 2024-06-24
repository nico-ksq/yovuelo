package migrations

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql" // Importa el driver de MySQL
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrator(db *sql.DB) {
	// Inicializar la migración
	driver, _ := mysql.WithInstance(db, &mysql.Config{
		MigrationsTable:  "gorp_migration",
		DatabaseName:     "u549962429_yovuelo",
		NoLock:           false,
		StatementTimeout: 5 * time.Second,
	})

	// Obtén el directorio de trabajo actual
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error obteniendo el directorio de trabajo actual: %v", err)
	}

	// Construye la ruta absoluta al directorio de migraciones
	migrationsDir := filepath.Join(wd, "db/migrations/sql")
	migrationsPath := "file://" + migrationsDir

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"u549962429_yovuelo",
		driver,
	)
	if err != nil {
		panic("Migration fails")
	}

	// Aplicar las migraciones y registrar cada una
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error al aplicar la migración: %v", err)
	}
}
