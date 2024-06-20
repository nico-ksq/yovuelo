package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"yovuelo/db/driver"
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
	db, err := driver.Register(dbHost, sshUser, sshPassword, sshPort)
	if err != nil {
		panic(err)
	}

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
