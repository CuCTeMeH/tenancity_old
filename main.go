package main

import (
	"log"
	"path/filepath"
	"tenancity/api/bill"
	"tenancity/api/core"
	"tenancity/api/estate"
	"tenancity/api/file"
	"tenancity/api/message"
	"tenancity/api/user"
)

func main() {
	// Use all CPU cores
	core.Server = core.NewInstance()

	//DB
	core.Server.InitDB()
	for _, connection := range core.Server.DB.Connections {
		defer connection.Close()
	}
	//DB
	AutoMigrate()
	core.AddLogrusHooks()
	core.Server.InitRoutes()

	InitModules(core.Server)

	core.Server.InitRouter()
	core.Server.Start()
}

func InitModules(Server *core.Instance) {
	user.Register(Server)
	estate.Register(Server)
	bill.Register(Server)
	file.Register(Server)
	message.Register(Server)
}

func AutoMigrate() {
	dir, err := filepath.Rel("", "migrations")
	if err != nil {
		log.Fatal(err)
	}
	core.AutoMigrateModules(dir)
}
