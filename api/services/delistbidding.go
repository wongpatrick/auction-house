package services

import (
	"auction-house-service/api/model"
	"auction-house-service/api/repository"
	"errors"
)

func DelistListing(listingID int) error {
	bidParams := model.BidSearchParams{
		ListingId: &listingID,
	}
	bids, err := repository.FetchBid(bidParams)
	if err != nil {
		return err
	}

	if len(bids) != 0 {
		return errors.New("There are currently bids for the listing")
	}

	ptrDelete := model.Deleted
	patchListingParams := model.Listing{
		Id:     &listingID,
		Status: &ptrDelete,
	}
	patchErr := repository.PatchListing(patchListingParams)
	if patchErr != nil {
		return err
	}
	return nil
}
