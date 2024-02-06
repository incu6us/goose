package gomigrations

import (
	"database/sql"

	"github.com/incu6us/goose/v3"
)

func init() {
	goose.AddMigration(up002, nil)
}

func up002(tx *sql.Tx) error {
	return createTable(tx, "bravo")
}
