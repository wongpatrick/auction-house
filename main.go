package main

// @title           Auction House Service
// @version         1.0
// @description     This is an auction house service.
// @termsOfService  https://www.inflexion.io/

// @contact.name   Auction House API Support
// @contact.url    https://www.inflexion.io/
// @contact.email  support@inflection.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  Auction House Service OpenAPI
// @externalDocs.url          https://inflection.io/resources/open-api/

import (
	"auction-house-service/api/handler/bid"
	"auction-house-service/api/handler/listing"
	"auction-house-service/api/handler/ping"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// A very basic Logger for the sake of the assignment
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.Default()
	r.Use(Logger())
	// r.use(AuthenticationMiddleware())
	r.GET("/ping", ping.GET)

	r.GET("/listing", listing.GET)
	// r.POST("/listing", listing.POST)
	// r.PATCH("/listing", listing.PATCH)
	// r.DELETE("/listing", listing.DELETE)

	r.GET("/bid", bid.GET)
	// r.POST("/bid", bid.POST)

	r.Run() // listen and serve on localhost:8080
}
