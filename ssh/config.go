package ssh

// Config struct definition
type Config struct {
	sshRemoteHost     string
	sshRemoteUser     string
	sshRemotePassword string
	sshRemotePort     string
}

// NewConfig is the constructor for Config
func NewConfig(host, user, password, port string) *Config {
	return &Config{
		sshRemoteHost:     host,
		sshRemoteUser:     user,
		sshRemotePassword: password,
		sshRemotePort:     port,
	}
}

// GetSSHRemoteHost returns the sshRemoteHost field
func (c *Config) GetSSHRemoteHost() string {
	return c.sshRemoteHost
}

// GetSSHRemoteUser returns the sshRemoteUser field
func (c *Config) GetSSHRemoteUser() string {
	return c.sshRemoteUser
}

// GetSSHRemotePassword returns the sshRemotePassword field
func (c *Config) GetSSHRemotePassword() string {
	return c.sshRemotePassword
}

// GetSSHRemotePort returns the sshRemotePort field
func (c *Config) GetSSHRemotePort() string {
	return c.sshRemotePort
}
