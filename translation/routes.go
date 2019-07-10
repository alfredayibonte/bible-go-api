package translation

import (
	"github.com/rkeplin/bible-go-api/core"
)

var handler = Handler{}

func AddRoutes() {
	routes := core.NewRoutes()

	routes.Add(core.Route{
		Name:        "FindAll",
		Method:      "GET",
		Pattern:     "/translations",
		HandlerFunc: handler.FindAll,
	})

	routes.Add(core.Route{
		Name:        "FindOne",
		Method:      "GET",
		Pattern:     "/translations/{id}",
		HandlerFunc: handler.FindOne,
	})
}
