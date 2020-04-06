# Tenancity

Tenancity Golang API

Requirements:
   - Go Chi Router: go get -u github.com/go-chi/chi
   - Migrations: https://github.com/golang-migrate/migrate
   - Auth: https://github.com/volatiletech/authboss
   - GORM DB ORM: go get -u github.com/jinzhu/gorm (always in .env define the key name lowercase e.g not Tenancity but Tenancity
   - Viper Config Management: go get github.com/spf13/viper
   - Logrus for logs: go get github.com/sirupsen/logrus
   - Logrus mail hook for mailing errors: go get github.com/zbindenren/logrus_mail
   - Logrus slack notification hook: go get github.com/johntdyer/slackrus
   - Logrus MySQL hook - https://github.com/LyricTian/logrus-mysql-hook
   - Dep for dependencies - https://github.com/golang/dep
