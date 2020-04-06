package message

import (
	"Tenancity/API/core"
	"log"
	"path/filepath"
)

func Register(I *core.Instance) {
	autoMigrate()
}

func autoMigrate() {
	dir, err := filepath.Abs("./message/migrations/")
	if err != nil {
		log.Fatal(err)
	}
	core.AutoMigrateModules(dir)
}
