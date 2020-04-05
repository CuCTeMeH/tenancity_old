package core

import (
	"github.com/LyricTian/logrus-mysql-hook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zbindenren/logrus_mail"
)

func AddLogrusHooks() {
	log := logrus.New()
	AddMailHook(log)
	AddMysqlHook(log)
}

func AddMailHook(log *logrus.Logger) {
	// if you do not need authentication for your smtp host FETCH FROM VIPER config
	name := viper.GetString("name")
	host := viper.GetString("mail.host")
	port := viper.GetInt("mail.port")
	user := viper.GetString("mail.user")
	password := viper.GetString("mail.pass")

	from := viper.GetString("mail.from")
	to := viper.GetString("mail.to")

	hook, err := logrus_mail.NewMailAuthHook(name, host, port, from, to, user, password)
	if err == nil {
		log.Hooks.Add(hook)
	}
}

func AddMysqlHook(log *logrus.Logger) {
	db := Server.DB.Connections["logs"].DB()
	//Fetch db from config and db init.
	mysqlHook := mysqlhook.Default(db, "logs")
	log.AddHook(mysqlHook)
}
