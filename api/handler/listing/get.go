package listing

import (
	"auction-house-service/api/model"
	"auction-house-service/api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET listing
// @Summary      fetches listing
// @Description  get listing based on the user's query params
// @Tags         listing
// @Accept       json
// @Produce      json
// @Param        q    query     model.ListingParams
// @Success      200  {array}   model.Listing
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /listing [get]
func GET(c *gin.Context) {
	log.Printf("FETCH")
	var listingParams model.ListingParams
	if c.ShouldBind(&listingParams) == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	data, err := services.FetchListing(listingParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
