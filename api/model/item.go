package model

type (
	Item struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		RarityId    int    `json:"rarityId"`
		ItemTypeId  int    `json:"itemTypeId"`
		Description string `json:"description"`
	}

	ItemParams struct {
		Id         *[]int
		Name       *string
		RarityId   *[]int
		ItemTypeId *[]int
	}
)
