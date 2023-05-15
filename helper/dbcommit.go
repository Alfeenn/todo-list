package helper

import "database/sql"

func CommitorRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {

		errorRollBack := tx.Rollback()
		PanicIfErr(errorRollBack)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfErr(errorCommit)
	}
}
