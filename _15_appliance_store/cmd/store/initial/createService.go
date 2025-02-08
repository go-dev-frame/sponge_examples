package initial

import (
	"strconv"

	"store/internal/config"
	"store/internal/server"

	"github.com/go-dev-frame/sponge/pkg/app"
)

// CreateServices create http service
func CreateServices() []app.IServer {
	var cfg = config.Get()
	var servers []app.IServer

	// create a http service
	httpAddr := ":" + strconv.Itoa(cfg.HTTP.Port)
	httpServer := server.NewHTTPServer(httpAddr,
		server.WithHTTPIsProd(cfg.App.Env == "prod"),
	)
	servers = append(servers, httpServer)

	return servers
}
