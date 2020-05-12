package estate

import (
	"Tenancity/API/core"
	"Tenancity/API/estate/routes"
)

func Register(I *core.Instance) {
	I.AddRoute("/estate", estate.Routes())
}
