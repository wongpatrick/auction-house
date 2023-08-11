package model

type Status = string

const (
	Active  Status = "A"
	Sold    Status = "S"
	Expired Status = "E"
	Deleted Status = "D"
)

type (
	Listing struct {
		Id            *int    `json:"id"`
		ItemId        *int    `json:"itemId"`
		SellerId      *int    `json:"sellerId"`
		StartingPrice *int    `json:"startingPrice"`
		BuyoutPrice   *int    `json:"buyoutPrice"`
		CreatedAt     *string `json:"createdAt"` // For simplicity sake
		EndedAt       *string `json:"endedAt"`   // For simplicity sake
		Status        *Status `json:"status"`
	}

	ListingParams struct {
		Id         *[]int
		ItemId     *[]int
		SellerId   *int
		CreatedAt  *string
		EndedAt    *string
		Name       *[]string
		RarityId   *[]int
		ItemTypeId *[]int
	}
)
