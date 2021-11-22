package repositories

import (
	"context"
	"os"
)

var ctx = context.TODO()

func getDBName() string {
	return os.Getenv("MONGODB_DBNAME")
}
