package models

// ISSUE: formattedName no longer available in API v2
type LinkedinUserInfo struct {
	ID        string `json:"id"`
	Email     string `json:"emailAddress"`
	FirstName struct {
		Localized       map[string]string `json:"localized"`
		PreferredLocale struct {
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"preferredLocale"`
	} `json:"firstName"`
	LastName struct {
		Localized       map[string]string `json:"localized"`
		PreferredLocale struct {
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"preferredLocale"`
	} `json:"lastName"`
}

// Response format for FirstName and LastName is in MultiLocaleString
// "localized":{
// 	"en_US":"2029 Stierlin Ct, Mountain View, CA 94043"
// },
// "preferredLocale":{
// 	"country":"US",
// 	"language":"en"
// }
