package app

import (
	"github.com/ChenXie-sci/goauth-react-demo/backend/controller/users"
	"github.com/gin-contrib/cors"
)

func malUrl() {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Context-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.POST("/api/register", users.Register)
	router.POST("/api/login", user.Login)
	router.GET("/api/user", user.Get)
	router.GET("/api/logout", users.logout)
}
