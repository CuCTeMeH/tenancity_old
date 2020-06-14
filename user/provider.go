package user

import (
	"tenancity/api/core"
	"tenancity/api/user/routes"
)

func Register(I *core.Instance) {
	I.AddRoute("/user", user.Routes())
}
