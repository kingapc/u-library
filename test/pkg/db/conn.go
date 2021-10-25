package conn

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

var (
	host     = "localhost"
	port     = 5432
	user     = utils.GoDotEnvVariable("DBUSER")
	password = utils.GoDotEnvVariable("DBPASS")
	dbname   = "library"
)

func GetConnection() (*sql.DB, error) {

	psCredentials := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
