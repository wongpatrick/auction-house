package repository

import (
	"auction-house-service/api/config"
	"auction-house-service/api/model"
	"log"

	sq "github.com/Masterminds/squirrel"
)

// This in theory should be in a different service
// TODO: Double check if i need to filter for other items as well
func FetchItem(itemParams model.ItemParams) ([]int, error) {
	var db, errdb = config.Connectdb()
	if errdb != nil {
		return nil, errdb
	}
	defer db.Close()

	query := buildItemQuery(itemParams)
	log.Printf(query.ToSql())
	rows, err := query.RunWith(db).Query()

	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	items := []int{}

	for rows.Next() {
		var r model.Item
		err = rows.Scan(
			&r.Id,
		)
		if err != nil {
			log.Printf("Scan: %v", err)
			return nil, err
		}
		items = append(items, r.Id)

	}

	return items, nil
}

func buildItemQuery(itemParams model.ItemParams) sq.SelectBuilder {
	query := sq.Select("*").From("item")

	if itemParams.Id != nil && len(*itemParams.Id) > 0 {
		query.Where(sq.Eq{"item_id": itemParams.Id})
	}

	if itemParams.RarityId != nil {
		query.Where(sq.Eq{"rarity_id": itemParams.RarityId})
	}

	if itemParams.ItemTypeId != nil {
		query.Where(sq.Eq{"item_type_id": itemParams.ItemTypeId})
	}

	return query
}
