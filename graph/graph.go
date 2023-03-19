package graph

import (
	"log"
	"net/http"
	"time"

	"github.com/andrewjdunn/speedometer-backend/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/records", func(c *gin.Context) {
		records, err := database.SpeedRecords()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, records)
	})

	r.Run()
}
