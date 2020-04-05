package core

import (
	"Tenancity/API/core/structs"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

func (i *Instance) InitConfig() *Instance {
	viper.SetConfigName("env")  // name of config file (without extension)
	viper.AddConfigPath(".")    //look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	//fmt.Println("INIT CONFIG")
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	i.Debug = viper.GetBool("debug")
	i.AppName = viper.GetString("name")
	i.Addr = viper.GetString("address")
	i.Version = viper.GetString("version")

	if i.DB == nil {
		i.DB = &core.DB{}
		if i.DB.Credentials == nil {
			i.DB.Credentials = make(map[string]core.DBConnection)
		}
	}

	viper.UnmarshalKey("databases", &i.DB.Credentials)
	return i
}

func (i *Instance) ConfigWatcher() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		i.InitConfig()
		//For the future we can change config here and then reboot tha server instance for silent reboot :) without rebooting the whole app.
	})
}

func (i *Instance) InitDB() {
	if i.DB.Connections == nil {
		i.DB.Connections = make(map[string]*gorm.DB)
	}

	for name, connection := range i.DB.Credentials {
		GORMConnection := i.InitDBConnection(connection)
		i.DB.Connections[name] = GORMConnection
	}
}

func (i *Instance) InitDBConnection(connection core.DBConnection) *gorm.DB {
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", connection.Username, connection.Password, connection.Host, connection.Port, connection.Name)
	db, err = gorm.Open("mysql", dbSource)
	if i.Debug == true {
		db.LogMode(true)
	}

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
