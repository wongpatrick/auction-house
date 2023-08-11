package services

import (
	"auction-house-service/api/model"
	"auction-house-service/api/repository"
)

// Post buyer bid after validating
func BuyerBidding(player model.User, listing model.Listing, bid model.Bid) (bool, error) {
	validateErr := ValidateBidForBidding(player, listing, bid)
	if validateErr != nil {
		return false, validateErr
	}

	postErr := repository.PostBid(bid)
	if postErr != nil {
		return false, postErr
	}

	return true, nil
}
