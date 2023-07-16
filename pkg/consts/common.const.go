package consts

type EnvE string
const (
	POSTGRES_USER EnvE = "POSTGRES_USER"
	POSTGRES_PASSWORD EnvE = "POSTGRES_PASSWORD"
	POSTGRES_DB EnvE = "POSTGRES_DB"
	POSTGRES_HOST EnvE = "POSTGRES_HOST"
	POSTGRES_PORT EnvE = "POSTGRES_PORT"
	POSTGRES_SSL_MODE EnvE = "POSTGRES_SSL_MODE"
)