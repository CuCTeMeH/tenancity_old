package main

import (
	"Tenancity/API/bill"
	"Tenancity/API/core"
	"Tenancity/API/estate"
	"Tenancity/API/file"
	"Tenancity/API/message"
	"Tenancity/API/user"
	"log"
	"path/filepath"
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
