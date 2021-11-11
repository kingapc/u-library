package conn

import (
	"testing"

	conn "github.com/rpinedafocus/u-library/internal/app/u-library/db"
)

func TestGetConnection(t *testing.T) {
	db, err := conn.GetConnection()

	if db == nil || err != nil {
		t.Fatalf(`Fail to connect the data base for %#v and  %#q`, db, err)
	}
}
