package repository

import (
	"auction-house-service/api/config"
	"auction-house-service/api/model"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func FetchListing(listingParams model.ListingParams) ([]model.Listing, error) {
	var db, errdb = config.Connectdb()
	if errdb != nil {
		return nil, errdb
	}
	defer db.Close()

	query := buildListingQuery(listingParams)
	log.Printf(query.ToSql())
	rows, err := query.RunWith(db).Query()

	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	listing := []model.Listing{}

	for rows.Next() {
		var r model.Listing
		err = rows.Scan(
			&r.Id,
			&r.ItemId,
			&r.SellerId,
			&r.StartingPrice,
			&r.BuyoutPrice,
			&r.CreatedAt,
			&r.EndedAt,
			&r.Status,
		)
		if err != nil {
			log.Printf("Scan: %v", err)
			return nil, err
		}
		listing = append(listing, r)

	}

	return listing, nil
}

func buildListingQuery(listingParams model.ListingParams) sq.SelectBuilder {
	query := sq.Select("*").From("bid")

	if listingParams.Id != nil && len(*listingParams.Id) > 0 {
		query.Where(sq.Eq{"bid_id": listingParams.Id})
	}

	if listingParams.ItemId != nil && len(*listingParams.Id) > 0 {
		query.Where(sq.Eq{"item_id": listingParams.ItemId})
	}

	if listingParams.SellerId != nil && len(*listingParams.Id) > 0 {
		query.Where(sq.Eq{"seller_id": listingParams.SellerId})
	}

	// TODO: Add orderby

	return query
}
