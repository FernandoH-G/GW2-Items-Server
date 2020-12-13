package item

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ItemNamePair struct {
	Items []pair
}

type pair struct {
	ID   int
	Name string
}

func (i *pair) UnmarshalJSON(b []byte) error {
	arr := []interface{}{&i.ID, &i.Name}
	err := json.Unmarshal(b, &arr)

	// Error checking
	if err != nil {
		return fmt.Errorf("Unmarshal item underlying array: %w", err)
	}
	if len(arr) != 2 {
		return fmt.Errorf("Items underlying array should be 2 elements, got %d", len(arr))
	}

	return nil
}

func (i *ItemNamePair) QueryNameID() error {
	apiURL := fmt.Sprintf("http://api.gw2tp.com/1/bulk/items-names.json")

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

	// Loads all of the json objects at once.
	err = json.Unmarshal(data, &i)
	if err != nil {
		return fmt.Errorf("Receiver i Unmarshal error: %w", err)
	}

	return nil
}
