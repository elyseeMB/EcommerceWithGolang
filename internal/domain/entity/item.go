package entity

type ItemType int

const (
	DigitalItem ItemType = iota + 1
	DefaultItem
)

type Item struct {
	Id       int      `json:"id"`
	Taxonomy int      `json:"taxonomy_id"`
	SellerID int      `json:"seller_id"`
	CardID   int      `json:"card_id"`
	Price    float64  `json:"price"`
	Quantity int      `json:"quantity"`
	ItemType ItemType `json:"itemType"`
}
