package config

// Configurations models a struct of configurations
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

// ServerConfigurations models a struct of server configurations
type ServerConfigurations struct {
	Port string
}

// DatabaseConfigurations models a struct of database configurations
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}
