package model

type (
	User struct {
		Id    *int    `json:"id"`
		Name  *string `json:"name"`
		Money *int    `json:"money"`
	}
)
