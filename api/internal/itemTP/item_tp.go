package itemTP

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type EquipmentTP struct {
	Results []ItemTP
}

type ItemTP struct {
	ID     int
	Name   string
	ImgURL string
	Buy    int
	Sell   int
}

type ItemPrice struct {
	Gold   string
	Silver string
	Copper string
}

// UnmarshalJSON is my custom unmarshal for the GW2TP API.
func (i *ItemTP) UnmarshalJSON(b []byte) error {
	arr := []interface{}{&i.ID, &i.Name, &i.ImgURL, &i.Buy, &i.Sell}
	err := json.Unmarshal(b, &arr)

	// Error checking
	if err != nil {
		return fmt.Errorf("Unmarshal item underlying array: %w", err)
	}
	if len(arr) != 5 {
		return fmt.Errorf("Items underlying array should be 3 elements, got %d", len(arr))
	}

	return nil
}

func QueryTP(id string) EquipmentTP {
	apiURL := fmt.Sprintf("http://api.gw2tp.com/1/items?ids=%v&fields=name,img,buy,sell", id)

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

	var item EquipmentTP
	err = json.Unmarshal(data, &item)
	if err != nil {
		fmt.Print("Unmarshal error: ", err, "\n")
	}
	fmt.Println("TP Item: ", item)
	return item
}

func ParsePrice(price int) ItemPrice {
	priceStr := strconv.Itoa(price)
	fmt.Println("priceStr: ", priceStr)
	priceLen := len(priceStr)
	itemPrice := ItemPrice{
		Gold:   "0",
		Silver: "0",
		Copper: "0",
	}

	switch {
	case priceLen > 4:
		itemPrice.Gold = priceStr[0 : priceLen-4]
		itemPrice.Silver = priceStr[priceLen-4 : priceLen-2]
		itemPrice.Copper = priceStr[priceLen-2 : priceLen]
		break
	case 4 >= priceLen && priceLen > 2:
		itemPrice.Silver = priceStr[0 : priceLen-2]
		itemPrice.Copper = priceStr[priceLen-2 : priceLen]
		break
	case 2 >= priceLen && priceLen > 0:
		itemPrice.Copper = priceStr[0:priceLen]
		break
	}
	return itemPrice
}
