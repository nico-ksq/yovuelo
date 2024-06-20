package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"yovuelo/routes/requests"
)

// ValidateRegisterUserRequest valida los campos de la solicitud de registro de usuario.
func ValidateRegisterUserRequest(req requests.RegisterUserRequest) error {
	if req.Email == "" || req.Password == "" {
		return errors.New("email y Password son obligatorios")
	}

	// Validar seguridad de la contraseña
	if !isPasswordStrong(req.Password) {
		return errors.New("la contraseña debe tener al menos 8 caracteres, incluyendo una letra mayúscula, una letra minúscula, un número y un carácter especial")
	}

	return nil
}

// isPasswordStrong verifica si una contraseña cumple con los requisitos de seguridad.
func isPasswordStrong(password string) bool {
	var passwordRegex = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	return passwordRegex.MatchString(password)
}

// RegistrarUsuarioHandler maneja las solicitudes de registro de nuevos usuarios.
func (s *Server) UserHandler(w http.ResponseWriter, r *http.Request) {
	var req requests.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// Validar los campos obligatorios y la seguridad de la contraseña
	if err := ValidateRegisterUserRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Guardar el usuario en la base de datos
	// TODO return error
	if s.user.Register(req) {
		http.Error(w, "Error al guardar el usuario", http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusCreated)
}
