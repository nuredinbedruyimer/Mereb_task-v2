package config

import (
	"time"

	"github.com/gin-contrib/cors"
)

var CorsConfig = cors.New(cors.Config{
	AllowOrigins:     []string{"*"}, // You can specify specific domains like https://mereb.com
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
})
