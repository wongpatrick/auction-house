package services

import (
	"auction-house-service/api/model"
	"errors"
	"time"
)

// validate the bid that is being made
func ValidateBidForBidding(player model.User, listing model.Listing, bid model.Bid) error {
	if *listing.SellerId == bid.Id {
		return errors.New("A seller cannot bid on their own listing")
	}

	if *listing.Status == model.Expired || *listing.Status == model.Deleted {
		return errors.New("A bid cannot be expired or deleted")
	}

	// Need to fix time date, in theory the listing should be expired if it's past the end time
	timeParse, _ := time.Parse("2006-01-02", *listing.EndedAt)
	if timeParse.Before(time.Now()) {
		return errors.New("the listing has expired")
	}

	if bid.BidAmount <= 0 {
		return errors.New("A bid cannot be less than or equal to 0")
	}

	if *player.Money < bid.BidAmount {
		return errors.New("Buyer does not have sufficient funds")
	}

	if bid.IsBuyout == true && &bid.BidAmount != listing.BuyoutPrice {
		return errors.New("A buyer's bid is not equal to buyout price")
	}

	return nil
}
