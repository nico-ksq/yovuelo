package driver

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"yovuelo/db"
	ssh2 "yovuelo/db/ssh"
)

func keyString(k ssh.PublicKey) string {
	return k.Type() + " " + base64.StdEncoding.EncodeToString(k.Marshal()) // e.g. "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTY...."
}

func trustedHostKeyCallback(trustedKey string) ssh.HostKeyCallback {
	if trustedKey == "" {
		return func(_ string, _ net.Addr, k ssh.PublicKey) error {
			log.Printf("WARNING: SSH-key verification is *NOT* in effect: to fix, add this trustedKey: %q", keyString(k))
			return nil
		}
	}

	return func(_ string, _ net.Addr, k ssh.PublicKey) error {
		ks := keyString(k)
		if trustedKey != ks {
			return fmt.Errorf("SSH-key verification: expected %q but got %q", trustedKey, ks)
		}

		return nil
	}
}

// Función para iniciar el túnel SSH con autenticación por usuario y contraseña
func sshTunnel(sshRemoteHost, sshRemoteUser, sshRemotePassword, sshRemotePort string) (*ssh.Client, error) {
	// Configurar la conexión SSH
	sshConfig := &ssh.ClientConfig{
		User: sshRemoteUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshRemotePassword),
		},
		HostKeyCallback: trustedHostKeyCallback(""),
	}

	// Establecer conexión SSH al servidor remoto
	sshConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", sshRemoteHost, sshRemotePort), sshConfig)
	if err != nil {
		return nil, fmt.Errorf("error connecting to SSH server: %v", err)
	}

	return sshConn, nil
}

// Función para registrar un usuario en la base de datos
func Register(ssh *ssh2.Config, db *db.Config) (*sql.DB, error) {
	// Iniciar el túnel SSH
	sshClient, err := sshTunnel(ssh.GetSSHRemoteHost(), ssh.GetSSHRemoteUser(), ssh.GetSSHRemotePassword(), ssh.GetSSHRemotePort())
	if err != nil {
		return nil, fmt.Errorf("error starting SSH tunnel: %v", err)
	}
	defer sshClient.Close()

	// Register the custom dial function with the MySQL driver
	mysql.RegisterDialContext("mysql+tcp", MySQLDialContext(sshClient, ""))

	// Configure MySQL connection
	dsn := fmt.Sprintf("%s:%s@mysql+tcp(%s:%s)/%s", db.GetDBUser(), db.GetDBPassword(), db.GetDBHost(), db.GetDBPort(), db.GetDBName())

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

	fmt.Println("Successfully connected to the database over SSH!")
	return conn, nil
}

// MySQLDialContext is a custom dialer for MySQL over SSH
func MySQLDialContext(client *ssh.Client, addr string) func(ctx context.Context, addr string) (net.Conn, error) {
	return func(ctx context.Context, addr string) (net.Conn, error) {
		return client.Dial("tcp", addr)
	}
}
