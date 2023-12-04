package repository

import (
	"auction-house-service/api/dbconfig"
	"auction-house-service/api/model"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func FetchBid(bidParams model.BidSearchParams) ([]model.Bid, error) {
	db := dbconfig.DB

	query := buildBidQuery(bidParams)
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

func buildBidQuery(bidParams model.BidSearchParams) sq.SelectBuilder {
	query := sq.Select("*").From("bid")

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
