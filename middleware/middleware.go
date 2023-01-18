package middleware

import (
	"myapp/data"

	"github.com/smartelink/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
