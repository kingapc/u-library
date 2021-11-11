package conn

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/rpinedafocus/u-library/internal/utils"
)

var (
	host     = utils.GoDotEnvVariable("DBHOST")               //"localhost"
	port, _  = strconv.Atoi(utils.GoDotEnvVariable("DBPORT")) //5432
	user     = utils.GoDotEnvVariable("DBUSER")
	password = utils.GoDotEnvVariable("DBPASS")
	dbname   = utils.GoDotEnvVariable("DATABASE") //"library"
)

func GetConnection() (*sql.DB, error) {

	psCredentials := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

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
