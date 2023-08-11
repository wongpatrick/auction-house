package services

import (
	"auction-house-service/api/model"
	"auction-house-service/api/repository"
)

// Fetch Listing with params
func FetchListing(listing model.ListingParams) ([]model.Listing, error) {
	// TODO: Query for specific item ID based off if user wants rarity/item type/name
	// Below is an example
	itemParams := model.ItemParams{
		Id:         listing.ItemId,
		RarityId:   listing.RarityId,
		ItemTypeId: listing.ItemTypeId,
	}
	itemIds, err := repository.FetchItem(itemParams)
	if err != nil {
		// LOG
		return []model.Listing{}, err
	}
	listing.ItemId = &itemIds
	// TODO: Query for specific listing after getting a list of specific item ids
	listingData, queryErr := repository.FetchListing(listing)
	if queryErr != nil {
		// LOG
		return []model.Listing{}, queryErr
	}

	return listingData, nil
}
