package item

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EquipmentInfo struct {
	Items []ItemInfo
}

type ItemInfo struct {
	Description string
	Type        string
	Rarity      string
	Level       int
}

func (e *EquipmentInfo) QueryInfo(id string) error {
	apiURL := fmt.Sprintf("https://api.guildwars2.com/v2/items?ids=%v", id)

	res, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("apiURL string err: %w", err)
	}
	defer res.Body.Close()

	// ReadAll returns Byte data.
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("res.Body err: %w", err)
	}

	// Needs to be &e.Items, not &e.
	err = json.Unmarshal(data, &e.Items)
	if err != nil {
		return fmt.Errorf("item Unmarshal err:%w ", err)
	}

	if e.Items[0].Description == "" {
		e.Items[0].Description = "No Description"
	}

	fmt.Println("Info Item: ", e.Items)
	if err != nil {
		return fmt.Errorf("item println err: %w", err)
	}
	return nil

}
