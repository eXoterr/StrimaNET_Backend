package router

import "github.com/eXoterr/StrimaNET_Backend/internal/router/handlers/accounts"

func (http *HTTPRouter) initHandlers() {
	internal := http.instance.Group("/api")
	{
		auth := internal.Group("/auth")
		{
			auth.GET("/", accounts.Login)
			auth.GET("/login", accounts.Login)
			auth.GET("/register", accounts.Register)
		}
	}
}
