package conn

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "root"
// 	dbname   = "library"
// )

func GetConnection() (*sql.DB, error) {

	psCredentials := utils.GoDotEnvVariable("DB_CONNECTION_STRING")

	db, err := sql.Open("postgres", psCredentials)

	if err == nil {
		err = db.Ping()
		if err != nil {
			return nil, err
		} else {
			return db, nil
		}
	} else {
		return nil, err
	}
}
