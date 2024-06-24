package driver

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	gossh "golang.org/x/crypto/ssh"
)

// Register Función para registrar drivers
func Register(dsn string) (*sql.DB, error) {
	// Conectar a la base de datos MySQL
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MySQL via SSH: %v", err)
	}

	conn.SetConnMaxLifetime(5 * time.Minute) // Tiempo máximo de vida de la conexión
	conn.SetMaxIdleConns(5)                  // Número máximo de conexiones inactivas en el pool
	conn.SetMaxOpenConns(20)                 // Número máximo de conexiones abiertas en el pool

	// Verify connection to the database
	err = conn.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return conn, nil
}

// MySQLDialContext is a custom dialer for MySQL over SSH
func MySQLDialContext(client *gossh.Client, addr string) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, addr string) (net.Conn, error) {
		return client.Dial("tcp", addr)
	}
}
