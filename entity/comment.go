package entity

import (
	"database/sql"
)

//Like model in Java

type Comment struct {
	ID      uint64
	Email   sql.NullString
	Comment sql.NullString
}
