package services

import (
	"auction-house-service/api/model"
	"auction-house-service/api/repository"
)

// Post buyer bid after validating
func buyerbidding(listing model.Listing, bid model.Bid) (bool, error) {
	validateErr := ValidateBidForBidding(listing, bid)
	if validateErr != nil {
		return false, validateErr
	}

	postErr := repository.PostBid(bid)
	if postErr != nil {
		return false, postErr
	}

	return true, nil
}
