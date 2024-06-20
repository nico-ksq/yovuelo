package driver

import (
	"crypto/tls"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

// create human-readable SSH-key strings
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
func Register(sshRemoteHost, sshRemoteUser, sshRemotePassword, sshRemotePort string) (*sql.DB, error) {
	// Iniciar el túnel SSH
	sshConn, err := sshTunnel(sshRemoteHost, sshRemoteUser, sshRemotePassword, sshRemotePort)
	if err != nil {
		return nil, fmt.Errorf("error starting SSH tunnel: %v", err)
	}
	defer sshConn.Close()

	//// Configurar el DSN (Data Source Name) para la conexión a MySQL a través del túnel SSH
	mysql.RegisterTLSConfig("custom", &tls.Config{}) // Opcional: configurar TLS si es necesario
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		mySqlUsername, mySqlPassword, sshLocalHost, sshLocalPort, mySqlDatabase)

	// Conectar a la base de datos MySQL
	//db, err := sql.Open("mysql", dsn)
	//if err != nil {
	//	return nil, fmt.Errorf("error connecting to MySQL via SSH: %v", err)
	//}

	//// Iniciar transacción
	//tx, err := db.Begin()
	//if err != nil {
	//	db.Close()
	//	return nil, fmt.Errorf("error starting transaction: %v", err)
	//}
	//
	//// Ejemplo de operación dentro de la transacción (no olvides hacer rollback o commit al finalizar)
	//_, err = tx.Exec("INSERT INTO usuarios (nombre) VALUES (?)", name)
	//if err != nil {
	//	tx.Rollback()
	//	db.Close()
	//	return nil, fmt.Errorf("error inserting data: %v", err)
	//}
	//
	//// Commit de la transacción
	//if err := tx.Commit(); err != nil {
	//	db.Close()
	//	return nil, fmt.Errorf("error committing transaction: %v", err)
	//}

	return nil, nil
}
