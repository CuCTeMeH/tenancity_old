package user

import (
	"Tenancity/API/core"
	"Tenancity/API/user/routes"
)

func Register(I *core.Instance) {
	I.AddRoute("/user", user.Routes())
}
