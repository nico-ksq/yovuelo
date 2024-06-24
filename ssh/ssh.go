package ssh

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
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
func Tunnel(sshRemoteHost, sshRemoteUser, sshRemotePassword, sshRemotePort string) (*ssh.Client, error) {
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
