package services

import (
	"auction-house-service/api/model"
	"errors"
)

// validate the bid that is being made
func ValidateBidForBidding(player model.User, listing model.Listing, bid model.Bid) error {
	if *listing.SellerId == bid.Id {
		return errors.New("A seller cannot bid on their own listing")
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
