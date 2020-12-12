package itemInfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ItemInfo struct {
	Description string
	Type        string
	Rarity      string
	Level       int
}

func QueryInfo(id string) []ItemInfo {
	apiURL := fmt.Sprintf("https://api.guildwars2.com/v2/items?ids=%v", id)

	res, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// ReadAll returns Byte data.
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print("ReadAll Error: ", err, "\n")
	}

	// slice of ItemInfo due to the way the json is structured from api call.
	var item []ItemInfo
	err = json.Unmarshal(data, &item)
	if err != nil {
		fmt.Print("Unmarshal error: ", err, "\n")
	}
	fmt.Println("info item: ", item)
	return item

}
