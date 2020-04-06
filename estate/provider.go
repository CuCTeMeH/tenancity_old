package estate

import (
	"Tenancity/API/core"
	"Tenancity/API/estate/routes"
	"log"
	"path/filepath"
)

func Register(I *core.Instance) {
	autoMigrate()
	I.AddRoute("/test", estate.Routes())
}

func autoMigrate() {
	dir, err := filepath.Abs("./estate/migrations/")
	if err != nil {
		log.Fatal(err)
	}
	core.AutoMigrateModules(dir)
}
