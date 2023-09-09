package cms

import (
	"github.com/debuqer/kuti/internal/config"
	"github.com/debuqer/kuti/pkg/dispatch"
)

type Blog struct {
	Config config.Config
}

func (b *Blog) ParseRoutes() {
	dispatch.GET(&dispatch.Route{
		Name:    "root",
		Pattern: "/",
	})

	for rName, r := range b.Config.Routes {
		dispatch.GET(&dispatch.Route{
			Name:    rName,
			Pattern: r.Url,
		})
	}
}

func (b *Blog) Query(address string) *dispatch.RouteOptions {
	r, _ := dispatch.Parse(address)

	return r
}
