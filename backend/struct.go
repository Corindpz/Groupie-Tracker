<<<<<<< HEAD
package backend

type Pricing struct {
	Ref      string `json:"ref"`
	Quantity int    `json:"quantity"`
}

type Content struct {
	Ref      string `json:"ref"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Bundle struct {
	Title           string    `json:"title"`
	Desc            string    `json:"desc"`
	Tag             string    `json:"tag"`
	PurchaseLimit   int       `json:"purchaseLimit"`
	IsAvailable     bool      `json:"isAvailable"`
	ExpireTimestamp int64     `json:"expireTimestamp"`
	ShopType        string    `json:"shopType"`
	Pricing         []Pricing `json:"pricing"`
	Content         []Content `json:"content"`
	OfferID         string    `json:"offerID"`
	Asset           string    `json:"asset"`
}

=======
package backend
>>>>>>> 5009e5137a04e6fbc208e7da539cf3e916a6bf7c
