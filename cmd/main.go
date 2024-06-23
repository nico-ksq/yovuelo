package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"yovuelo/db/driver"
	"yovuelo/db/migrations"
	"yovuelo/db/ssh"
	"yovuelo/server"
)

var (
	dbHost      string
	dbPort      string
	dbUser      string
	dbPassword  string
	dbName      string
	sshUser     string
	sshPassword string
	sshPort     string
)

func main() {
	setEnvs()

	//Conecta la DB
	sshConfig := ssh.NewConfig(dbHost, sshUser, sshPassword, sshPort)
	dbConfig := driver.NewConfig(dbHost, dbUser, dbPassword, dbName, dbPort)

	// Iniciar el túnel SSH
	sshClient, err := ssh.Tunnel(sshConfig.GetSSHRemoteHost(), sshConfig.GetSSHRemoteUser(), sshConfig.GetSSHRemotePassword(), sshConfig.GetSSHRemotePort())
	if err != nil {
		panic("error starting SSH tunnel: " + err.Error())
	}
	// Registrar dial function con MySQL driver
	mysql.RegisterDialContext("mysql+tcp", driver.MySQLDialContext(sshClient, ""))

	// Configure MySQL connection
	dsn := fmt.Sprintf("%s:%s@mysql+tcp(%s:%s)/%s", dbConfig.GetDBUser(), dbConfig.GetDBPassword(), dbConfig.GetDBHost(), dbConfig.GetDBPort(), dbConfig.GetDBName())

	db, err := driver.Register(sshConfig, dbConfig, dsn)
	if err != nil {
		panic(err)
	}

	migrations.Migrator(db, dsn)

	// Inicializa el servidor
	srv := server.NewServer(db)

	// Configura las rutas
	srv.SetupRoutes()

	// Inicia el servidor
	log.Println("Servidor escuchando en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", srv.Router))
}

func setEnvs() {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")

	sshUser = os.Getenv("SSH_USER")
	sshPassword = os.Getenv("SSH_PASSWORD")
	sshPort = os.Getenv("SSH_PORT")

	// Verificar que ninguna variable esté vacía
	missingEnv := false
	var missingVars []string

	if dbHost == "" {
		missingVars = append(missingVars, "DB_HOST")
		missingEnv = true
	}
	if dbPort == "" {
		missingVars = append(missingVars, "DB_PORT")
		missingEnv = true
	}
	if dbUser == "" {
		missingVars = append(missingVars, "DB_USER")
		missingEnv = true
	}
	if dbPassword == "" {
		missingVars = append(missingVars, "DB_PASSWORD")
		missingEnv = true
	}
	if dbName == "" {
		missingVars = append(missingVars, "DB_NAME")
		missingEnv = true
	}
	if sshUser == "" {
		missingVars = append(missingVars, "SSH_USER")
		missingEnv = true
	}
	if sshPassword == "" {
		missingVars = append(missingVars, "SSH_PASSWORD")
		missingEnv = true
	}
	if sshPort == "" {
		missingVars = append(missingVars, "SSH_PORT")
		missingEnv = true
	}

	if missingEnv {
		errMsg := fmt.Sprintf("Falta al menos una variable de entorno requerida: %v", missingVars)
		panic(errMsg)
	}
}
