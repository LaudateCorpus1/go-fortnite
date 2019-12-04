package fortnite

// Item is a item data in fortnite
type Item struct {
	ImageURL      string `json:"imageUrl"`
	ManifestID    int    `json:"manifestId"`
	Name          string `json:"name"`
	Rarity        string `json:"rarity"`
	StoreCategory string `json:"storeCategory"`
	VBucks        int    `json:"vBucks"`
}

func GetStore(s *Service) (result []Item, err error) {
	_, err = s.Do("store", result)
	return
}
