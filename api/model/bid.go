package model

type (
	Bid struct {
		Id        int    `json:"id"`
		ListingId int    `json:"listingId"`
		BidderId  int    `json:"bidderId"`
		BidAmount int    `json:"bidAmount"`
		CreatedAt string `json:"createdAt"` // For simplicity sake
		IsBuyout  bool   `json:"isBuyout"`
	}

	BidSearchParams struct {
		Id        *[]int
		ListingId *int
		BidderId  *int
		BidAmount *int
	}
)
