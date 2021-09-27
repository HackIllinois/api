package models

type RedeemEventRequest struct {
	ID      string `json:"id"`
	EventID string `json:"eventID"`
}

type RedeemEventResponse struct {
	Status string `json:"status"`
}

type AwardPointsRequest struct {
	ID     string `json:"id"`
	Points int    `json:"points"`
}

type AwardShopPointsRequest struct {
	ID         string `json:"id"`
	ShopPoints int    `json:"shopPoints"`
}

type RedeemShopPointsRequest struct {
	ID         string `json:"id"`
	ShopItemID string `json:"itemID"`
}
