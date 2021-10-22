package middlewares

import (
	"github.com/gin-gonic/gin"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	return cors.New(cors.Config{
// 		AllowOrigins:     []string{"http://localhost:3000"},
// 		AllowMethods:     []string{"GET", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 		MaxAge:           12 * time.Hour,
// 	})
// }

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
