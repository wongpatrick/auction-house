package repository

import (
	"auction-house-service/api/config"
	"auction-house-service/api/model"
	"log"

	sq "github.com/Masterminds/squirrel"
)

// This in theory should be in a different service
func FetchItem(listingParams model.ListingParams) ([]int, error) {
	var db, errdb = config.Connectdb()
	if errdb != nil {
		return nil, errdb
	}
	defer db.Close()

	query := buildItemQuery(listingParams)
	log.Printf(query.ToSql())
	rows, err := query.RunWith(db).Query()

	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	bid := []model.Bid{}

	for rows.Next() {
		var r model.Bid
		err = rows.Scan(
			&r.Id,
			&r.ListingId,
			&r.BidderId,
			&r.BidAmount,
			&r.CreatedAt,
			&r.IsBuyout,
		)
		if err != nil {
			log.Printf("Scan: %v", err)
			return nil, err
		}
		bid = append(bid, r)

	}

	return bid, nil
}

func buildItemQuery(itemParams model.ListingParams) sq.SelectBuilder {
	query := sq.Select("*").From("item")

	if bidParams.Id != nil && len(*bidParams.Id) > 0 {
		query.Where(sq.Eq{"bid_id": bidParams.Id})
	}

	if bidParams.ListingId != nil {
		query.Where(sq.Eq{"listing_id": bidParams.ListingId})
	}

	if bidParams.BidderId != nil {
		query.Where(sq.Eq{"bidder_id": bidParams.BidderId})
	}

	return query
}
