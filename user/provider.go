package user

import (
	"Tenancity/API/core"
	"Tenancity/API/user/routes"
	"log"
	"path/filepath"
)

func Register(I *core.Instance) {
	autoMigrate()
	I.AddRoute("/test", user.Routes())
}

func autoMigrate() {
	dir, err := filepath.Abs("./user/migrations/")
	if err != nil {
		log.Fatal(err)
	}
	core.AutoMigrateModules(dir)
}
