package router

import (
	"github.com/eXoterr/StrimaNET_Backend/pkg/handlers/accounts"
	"github.com/eXoterr/StrimaNET_Backend/pkg/handlers/adapters"
)

func (http *HTTPRouter) initHandlers() {
	internal := http.instance.Group("/api")
	{
		auth := internal.Group("/auth")
		{
			auth.GET("/", adapters.Adapter(accounts.Login))
			auth.GET("/login", adapters.Adapter(accounts.Login))
			auth.GET("/register", adapters.Adapter(accounts.Register))
		}
	}
}
