package repository

import (
	"auction-house-service/api/model"

	sq "github.com/Masterminds/squirrel"
)

func PatchListing(listing model.Listing) error {
	// TODO: UPDATE DB with listing values
	_ = buildListingUpdate(listing)
	return nil
}

func buildListingUpdate(listParams model.Listing) sq.UpdateBuilder {
	update := sq.Update("listing").Where("listing_id", listParams.Id)

	if listParams.Status != nil {
		update.Set("status", *listParams.Status)
	}

	return update
}
