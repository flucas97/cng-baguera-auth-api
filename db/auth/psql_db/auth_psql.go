package psql_db

import (
	"database/sql"
	"os"

	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
)

func init() {
	var err error
	psqlInfo := os.Getenv("BAGUERA_DSN")

	Client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error("cannot open database connection", err)
		panic(err)
	}

	err = Client.Ping()
	if err != nil {
		logger.Error("cannot ping database", err)
		panic(err)
	}

	logger.Info("accounts_db successfuly connected")
}
