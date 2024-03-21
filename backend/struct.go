
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

type News struct {
	Title      string `json:"title"`
	Link       string `json:"link"`
	ImageURL   string `json:"img"`
	ShortDesc  string `json:"short_desc"`
}

type NewsData struct {
	News []News `json:"news"`
	PageNext string `json:"next"`
	PagePrev string `json:"prev"`
}

type RankData struct {
    FoundRank             int    `json:"foundRank"`
    Val                   int    `json:"val"`
    UID                   string `json:"uid"`
    UpdateTimestamp       int    `json:"updateTimestamp"`
    TotalMastersAndPreds int    `json:"totalMastersAndPreds"`
}

type PlatformData struct {
    PC     RankData `json:"PC"`
    PS4    RankData `json:"PS4"`
    X1     RankData `json:"X1"`
    SWITCH RankData `json:"SWITCH"`
}

type RotaData struct {
    RP PlatformData `json:"RP"`
    AP PlatformData `json:"AP"`
}

type RotationInfo struct {
	Start             int    `json:"start"`
	End               int    `json:"end"`
	ReadableDateStart string `json:"readableDate_start"`
	ReadableDateEnd   string `json:"readableDate_end"`
	Map               string `json:"map"`
	Code              string `json:"code"`
	DurationInSecs    int    `json:"DurationInSecs"`
	DurationInMinutes int    `json:"DurationInMinutes"`
	Asset             string `json:"asset,omitempty"`
	RemainingSecs     int    `json:"remainingSecs,omitempty"`
	RemainingMins     int    `json:"remainingMins,omitempty"`
	RemainingTimer    string `json:"remainingTimer,omitempty"`
}

type MapRotation struct {
	Current RotationInfo `json:"current"`
	Next    RotationInfo `json:"next"`
}

type Tracker struct {
    Global        GlobalData      `json:"global"`
    Realtime      RealtimeData    `json:"realtime"`
    Internal      InternalData    `json:"mozambiquehere_internal"`
    Total         TotalData       `json:"total"`
    ProcessingTime float64        `json:"processingTime"`
}

type GlobalData struct {
    Name    string    `json:"name"`
    Avatar  string    `json:"avatar"`
    Rank    RankData2 `json:"rank"`
}

type RankData2 struct {
    RankName  string  `json:"rankName"`
}

type RealtimeData struct {
    SelectedLegend string  `json:"selectedLegend"`
    CurrentState   string  `json:"currentState"`
}

type InternalData struct {
    IsNewToDB   bool    `json:"isNewToDB"`
}

type TotalData struct {
    Kills             TotalValue  `json:"kills"`
    WinsSeason20      TotalValue  `json:"wins_season_20"`
    Damage            TotalValue  `json:"damage"`
    EnemiesScanned    TotalValue  `json:"enemies_scanned"`
    SpecialEventKills TotalValue  `json:"specialEvent_kills"`
    KillsSeason20     TotalValue  `json:"kills_season_20"`
    KD                TotalValueKD  `json:"kd"`
}

type TotalValue struct {
    Name string  `json:"name"`
	Value int `json:"value"`
}

type TotalValueKD struct {
    Name string  `json:"name"`
	Value string `json:"value"`
}