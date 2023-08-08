package services

import (
	"auction-house-service/api/model"
)

// Fetch bid with params
func FetchBid(listing model.BidSearchParams) ([]model.Bid, error) {
	// TODO: Query for specific bid, using a bid repoistory

	return []model.Bid{}, nil
}
