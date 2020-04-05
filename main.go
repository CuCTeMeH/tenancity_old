package main

import (
	"Tenancity/API/core"
	"Tenancity/API/groups"
	"Tenancity/API/user"
	"runtime"
)

func main() {
	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	core.Server = core.NewInstance()

	//DB
	core.Server.InitDB()
	for _, connection := range core.Server.DB.Connections {
		defer connection.Close()
	}
	//DB

	core.AddLogrusHooks()
	core.Server.InitRoutes()

	InitModules(core.Server)

	core.Server.InitRouter()
	core.Server.Start()
}

func InitModules(Server *core.Instance) {
	user.Register(Server)
	groups.Register(Server)
}
