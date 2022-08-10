package models

type GetPrizeRequest struct {
	ID string `json:"id"`
}

type DeletePrizeRequest struct {
	ID string `json:"id"`
}

type AwardPointsRequest struct {
	ID         string `json:"id"`
	ShopPoints int    `json:"shopPoints"`
}

type RedeemPointsRequest struct {
	ID         string `json:"id"`
	ShopItemID string `json:"itemID"`
}
