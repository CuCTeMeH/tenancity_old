package estate

import (
	"tenancity/api/core"
	"tenancity/api/estate/routes"
)

func Register(I *core.Instance) {
	I.AddRoute("/estate", estate.Routes())
}
