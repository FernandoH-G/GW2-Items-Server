package item

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type EquipmentTP struct {
	Results []itemTP
}

type itemTP struct {
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
func (i *itemTP) UnmarshalJSON(b []byte) error {
	arr := []interface{}{&i.ID, &i.Name, &i.ImgURL, &i.Buy, &i.Sell}
	err := json.Unmarshal(b, &arr)

	// Error checking
	if err != nil {
		return fmt.Errorf("Unmarshal item underlying array: %w", err)
	}
	if len(arr) != 5 {
		return fmt.Errorf("Items underlying array should be 5 elements, got %d", len(arr))
	}

	return nil
}

func (e *EquipmentTP) QueryTP(id string) error {
	apiURL := fmt.Sprintf("http://api.gw2tp.com/1/items?ids=%v&fields=name,img,buy,sell", id)

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

	err = json.Unmarshal(data, &e)
	if err != nil {
		return fmt.Errorf("Receiver e Unmarshal err: %w", err)
	}
	fmt.Println("TP Item: ", e)
	if err != nil {
		return fmt.Errorf("e println error: %w", err)
	}
	return nil
}

func (i *ItemPrice) FormatPrice(price int) error {
	priceStr := strconv.Itoa(price)
	_, err := fmt.Println("priceStr: ", priceStr)
	if err != nil {
		return fmt.Errorf("priceStr println err: %w", err)
	}
	priceLen := len(priceStr)
	i.Gold = "00"
	i.Silver = "00"
	i.Copper = "00"

	switch {
	case priceLen > 4:
		i.Gold = priceStr[0 : priceLen-4]
		i.Silver = priceStr[priceLen-4 : priceLen-2]
		i.Copper = priceStr[priceLen-2 : priceLen]
		break
	case 4 >= priceLen && priceLen > 2:
		i.Silver = priceStr[0 : priceLen-2]
		i.Copper = priceStr[priceLen-2 : priceLen]
		break
	case 2 >= priceLen && priceLen > 0:
		i.Copper = priceStr[0:priceLen]
		break
	}
	return nil
}
