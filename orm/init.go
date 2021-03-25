package orm

import "github.com/MarekWojt/GoMan/db"

// Init migrates all databases
func Init() {
	db.DB.AutoMigrate(User{})
}
