package helper

import (
	"database/sql"
	"log"
)

func CommitorRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		log.Print("error rollback")
		errorRollBack := tx.Rollback()
		log.Print("doing rollback")
		PanicIfErr(errorRollBack)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfErr(errorCommit)
	}
}
