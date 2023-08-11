package repository

import (
	"auction-house-service/api/config"
	"auction-house-service/api/model"
	"log"

	sq "github.com/Masterminds/squirrel"
)

// This in theory should be in a different service
func FetchPlayer(userId int) (model.User, error) {
	var db, errdb = config.Connectdb()
	if errdb != nil {
		return model.User{}, errdb
	}
	defer db.Close()

	query := buildPlayerQuery(userId)
	log.Printf(query.ToSql())
	rows, err := query.RunWith(db).Query()

	if err != nil {
		log.Printf(err.Error())
		return model.User{}, err
	}

	player := model.User{}

	for rows.Next() {
		err = rows.Scan(
			&player.Id,
		)
		if err != nil {
			log.Printf("Scan: %v", err)
			return model.User{}, err
		}

	}

	return player, nil
}

func buildPlayerQuery(userId int) sq.SelectBuilder {
	query := sq.Select("*").From("user").Where(sq.Eq{"user_id": userId})

	return query
}
