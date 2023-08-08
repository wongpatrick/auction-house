package bid

import (
	"auction-house-service/api/model"
	"auction-house-service/api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET bid
// @Summary      fetches bid
// @Description  get bid based on the user's query params
// @Tags         bid
// @Accept       json
// @Produce      json
// @Param        q    query     model.BidParams
// @Success      200  {array}   model.Bid
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /bid [get]
func GET(c *gin.Context) {
	log.Printf("FETCH")
	var bidParams model.BidSearchParams
	if c.ShouldBind(&bidParams) == nil {
		// LOG
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	data, err := services.FetchBid(bidParams)
	if err != nil {
		// LOG
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
