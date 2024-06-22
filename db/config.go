package db

// Config struct definition
type Config struct {
	dbHost     string
	dbUser     string
	dbPassword string
	dbName     string
	dbPort     string
}

// NewConfig is the constructor for Config
func NewConfig(host, user, password, name, port string) *Config {
	return &Config{
		dbHost:     host,
		dbUser:     user,
		dbPassword: password,
		dbName:     name,
		dbPort:     port,
	}
}

// GetDBHost returns the dbHost field
func (c *Config) GetDBHost() string {
	return c.dbHost
}

// GetDBUser returns the dbUser field
func (c *Config) GetDBUser() string {
	return c.dbUser
}

// GetDBPassword returns the dbPassword field
func (c *Config) GetDBPassword() string {
	return c.dbPassword
}

// GetDBName returns the dbName field
func (c *Config) GetDBName() string {
	return c.dbName
}

// GetDBPort returns the dbPort field
func (c *Config) GetDBPort() string {
	return c.dbPort
}
