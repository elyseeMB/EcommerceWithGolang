package entity

type PromotionType int

const (
	SameSellerPromotion PromotionType = iota + 1
	TaxonomyPromotion
	TotalPricePromotion
)

type SameSellerPromotionDiscount struct {
	DiscountRate float64 `json:"discountRate"`
}

type TaxonomyPromotionDiscount struct {
	DiscountRate float64 `json:"discountRate"`
	TaxonomyID   int     `json:"taxonomyID"`
}

type TotalPricePromotionDiscount struct {
	PriceRangeStart float64 `json:"priceRangeStart"`
	PriceRangeEnd   float64 `json:"priceRangeEnd"`
	DiscountAmount  float64 `json:"discountAmount"`
}

type Promotion struct {
	Id            int
	PromotionType PromotionType                  `json:"promotionType"`
	SameSellerP   *SameSellerPromotionDiscount   `json:"sameSellerPromotion"`
	TaxonomyP     *TaxonomyPromotionDiscount     `json:"taxonomyPromotion"`
	TotalPriceP   []*TotalPricePromotionDiscount `json:"totalPricePromotions"`
}
