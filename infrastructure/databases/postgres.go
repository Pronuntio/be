package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(host string, port uint, dbname, user, pass string) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("%s:%s@%s:%d/%s", user, pass, host, port, dbname))
}
