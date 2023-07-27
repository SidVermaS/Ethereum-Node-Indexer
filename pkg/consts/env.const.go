package consts

type EnvE string

const (
	API_HOST EnvE ="API_HOST"
	API_PORT EnvE ="API_PORT")
const (
	PG_USER     EnvE = "PG_USER"
	PG_PASSWORD EnvE = "PG_PASSWORD"
	PG_DB       EnvE = "PG_DB"
	PG_HOST     EnvE = "PG_HOST"
	PG_PORT     EnvE = "PG_PORT"
	PG_SSL_MODE EnvE = "PG_SSL_MODE"	
)
//	Constants to access environment variables needed to access Redis cache
const (	
	REDIS_HOST EnvE="REDIS_HOST"
	REDIS_PORT EnvE="REDIS_PORT"
	REDIS_PASSWORD EnvE="REDIS_PASSWORD"
)


const (	
	CONSENSYS_CLIENT_HOST EnvE="CONSENSYS_CLIENT_HOST"
	CONSENSYS_CLIENT_PORT EnvE="CONSENSYS_CLIENT_PORT"
)