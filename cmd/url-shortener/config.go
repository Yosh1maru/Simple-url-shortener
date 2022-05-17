package main

// Config app configuration file
type Config struct {
	// HttpServerHost http sever host name
	HttpServerHost string `env:"HTTP_SERVER_HOST"`
	// HttpServerPort http sever port
	HttpServerPort string `env:"HTTP_SERVER_PORT"`
	// EncryptorSalt salt for encryption url
	EncryptorSalt string `env:"ENCRYPTOR_SALT"`
	// DbHost db host
	DbHost string `env:"DB_HOST"`
	// DbPort db port
	DbPort string `env:"DB_PORT"`
	// DbName database for connection
	DbName string `env:"DB_DATABASE"`
	// DbUser db user
	DbUser string `env:"DB_USER"`
	// DbPassword db user password
	DbPassword string `env:"DB_PASSWORD"`
}
