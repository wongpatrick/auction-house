package repository

import (
	"auction-house-service/api/dbconfig"
	"auction-house-service/api/model"
)

func PostBid(bid model.Bid) error {
	// TODO: connect to db and create bid
	_ = dbconfig.DB
	return nil
}
