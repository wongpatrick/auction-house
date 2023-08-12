package bid

import (
	"auction-house-service/api/model"
	"auction-house-service/api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST bid
// @Summary      Create a bid
// @Description  Create bid based on the user's bid request
// @Tags         bid
// @Accept       json
// @Produce      json
// @Param        q    query     model.Bid
// @Success      204
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /bid [get]
func POST(c *gin.Context) {
	log.Printf("POST")
	var bid model.Bid
	if c.ShouldBind(&bid) == nil {
		// LOG
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	data, err := services.BuyerBidding(bid)
	if err != nil {
		// LOG
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
