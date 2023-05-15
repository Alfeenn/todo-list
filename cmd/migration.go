package cmd

import (
	"flag"
	"log"

	"github.com/Alfeenn/online-learning/databases"
	"github.com/Alfeenn/online-learning/middleware"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

var enforcer = middleware.UserPolicy()

func Command(tables []interface{}) (bool, error) {
	// MigrateCmd the schema
	migrate := flag.String("migrate", "", "For migration up or down")
	flag.Parse() // parse the flags from the command line, must be called after all flags are defined and before flags are accessed by the program

	// migrate up or down
	if *migrate == "up" { // migrate up
		return true, Up(databases.MigrationDB(), tables)
	} else if *migrate == "down" { // migrate down
		return true, Down(databases.MigrationDB(), tables)
	}

	return false, nil
}

func MigrateCmd() (bool, *casbin.Enforcer) {
	// Schema table
	tables := databases.Tables()
	check, err := Command(tables)

	// check error
	if err != nil {
		panic(err)
	}
	if err = enforcer.LoadPolicy(); err != nil {
		log.Fatal("Failed to load policy")
	}

	return check, enforcer
}

// Up migrates the schema
func Up(db *gorm.DB, dst []interface{}) error {
	for _, v := range dst {
		err := db.Migrator().AutoMigrate(v)
		if err != nil {
			return err
		}
	}

	return nil
}

// Down drops the schema
func Down(db *gorm.DB, dst []interface{}) error {
	for _, v := range dst {
		err := db.Migrator().DropTable(v)
		if err != nil {
			return err
		}
	}

	return nil
}
