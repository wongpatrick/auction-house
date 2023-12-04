package ping

import (
	"auction-house-service/api/dbconfig"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET ping
// @Summary      ping database
// @Description  ping database to verify connection
// @Tags         ping
// @Accept       json
// @Produce      json
// @Success      200  {array}   string
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /ping [get]
func GET(c *gin.Context) {
	_ = dbconfig.DB
	// TODO: Select something to see if you can query as a ping
	// if errdb != nil {
	// 	log.Printf(errdb.Error())
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Missing Connection",
	// 	})
	// 	return
	// }
	// defer db.Close()
	log.Printf("PING")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success Connection",
	})
}
