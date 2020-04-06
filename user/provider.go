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
	dir, err := filepath.Rel("", "user/migrations/")
	if err != nil {
		log.Fatal(err)
	}
	//println(dir)
	//return
	core.AutoMigrateModules(dir)
}
