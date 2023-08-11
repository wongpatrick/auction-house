package services

import (
	"auction-house-service/api/model"
	"auction-house-service/api/repository"
	"errors"
)

// Post buyer bid after validating
func BuyerBidding(bid model.Bid) (bool, error) {
	listingParams := model.ListingParams{
		Id: &[]int{bid.ListingId},
	}
	listing, err := FetchListing(listingParams)

	if err != nil {
		return false, err
	}

	if len(listing) != 1 {
		errMsg := errors.New("There is either no listing or more than 1 was found")
		return false, errMsg
	}

	player, err := repository.FetchPlayer(bid.BidderId)
	if err != nil {
		return false, err
	}

	validateErr := ValidateBidForBidding(player, listing[0], bid)
	if validateErr != nil {
		return false, validateErr
	}

	postErr := repository.PostBid(bid)
	if postErr != nil {
		return false, postErr
	}

	return true, nil
}
