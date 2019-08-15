package configuration

type PostgresConfiguration struct {
	Host     string
	Port     uint32
	Username string
	Password string
	DBName   string
}
