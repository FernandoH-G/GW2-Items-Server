package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/FernandoH-G/gw2-items-server/graph/generated"
	"github.com/FernandoH-G/gw2-items-server/graph/model"
	"github.com/FernandoH-G/gw2-items-server/internal/item"
)

func (r *queryResolver) GetItemByID(ctx context.Context, id string) (*model.Item, error) {
	var infoItem item.EquipmentInfo
	err := infoItem.QueryInfo(id)
	if err != nil {
		return nil, fmt.Errorf("infoItem err: %w", err)
	}

	var tpItem item.EquipmentTP
	err = tpItem.QueryTP(id)
	if err != nil {
		return nil, fmt.Errorf("tpItem err: %w", err)
	}

	var sellPrice item.ItemPrice
	err = sellPrice.FormatPrice(tpItem.Results[0].Sell)
	if err != nil {
		return nil, fmt.Errorf("sellPrice err: %w", err)
	}

	var buyPrice item.ItemPrice
	err = buyPrice.FormatPrice(tpItem.Results[0].Buy)
	if err != nil {
		return nil, fmt.Errorf("buyPrice err: %w", err)
	}

	resultItem := &model.Item{
		ID:     strconv.Itoa(tpItem.Results[0].ID),
		Name:   tpItem.Results[0].Name,
		ImgURL: tpItem.Results[0].ImgURL,
		Sell: &model.Price{
			Gold:   sellPrice.Gold,
			Silver: sellPrice.Silver,
			Copper: sellPrice.Copper,
		},
		Buy: &model.Price{
			Gold:   buyPrice.Gold,
			Silver: buyPrice.Silver,
			Copper: buyPrice.Copper,
		},
		Description: &infoItem.Items[0].Description,
		Type:        infoItem.Items[0].Type,
		Rarity:      infoItem.Items[0].Rarity,
		Level:       strconv.Itoa(infoItem.Items[0].Level),
	}

	return resultItem, nil
}

func (r *queryResolver) GetItemNames(ctx context.Context) ([]*model.ItemNamePair, error) {
	var resultItemNames []*model.ItemNamePair

	inp := item.ItemNamePair{}
	err := inp.QueryNameID()
	if err != nil {
		return nil, fmt.Errorf("inp err: %w", err)
	}

	for _, in := range inp.Items {
		resultItemNames = append(resultItemNames, &model.ItemNamePair{
			ID:   strconv.Itoa(in.ID),
			Name: in.Name})
	}

	return resultItemNames, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
