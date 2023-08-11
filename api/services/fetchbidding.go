package services

import (
	"auction-house-service/api/model"
	"auction-house-service/api/repository"
)

// Fetch bid with params
func FetchBid(bidParams model.BidSearchParams) ([]model.Bid, error) {
	bid, err := repository.FetchBid(bidParams)
	if err != nil {
		// LOG
		return []model.Bid{}, err
	}
	return bid, nil
}
