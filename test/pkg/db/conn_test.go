package conn

import (
	"testing"
)

func TestGetConnection(t *testing.T) {
	db, err := GetConnection()

	if db == nil || err != nil {
		t.Fatalf(`Fail to connect the data base for %#v and  %#q`, db, err)
	}
}
