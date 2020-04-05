package groups

import (
	"Tenancity/API/core"
	"Tenancity/API/groups/routes"
)

func Register(I *core.Instance) {
	I.AddRoute("/grouptest", groups.Routes())
}
