package link_shorter

import (
	"database/sql"
)

type DbLinkShorter struct {
	db *sql.DB
}
