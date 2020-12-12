package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/FernandoH-G/gw2-items-server/graph/generated"
	"github.com/FernandoH-G/gw2-items-server/graph/model"
	"github.com/FernandoH-G/gw2-items-server/internal/itemInfo"
	"github.com/FernandoH-G/gw2-items-server/internal/itemTP"
)

func (r *queryResolver) GetItemByID(ctx context.Context, id string) (*model.Item, error) {
	var tpItem = itemTP.QueryTP(id)
	var infoItem = itemInfo.QueryInfo(id)
	sellPrice := itemTP.ParsePrice(tpItem.Results[0].Sell)
	buyPrice := itemTP.ParsePrice(tpItem.Results[0].Buy)

	resultItem := &model.Item{
		ID:     fmt.Sprint(tpItem.Results[0].ID),
		Name:   tpItem.Results[0].Name,
		ImgURL: tpItem.Results[0].ImgURL,
		Sell: &model.Price{
			Gold: sellPrice.Gold,
			Silver: sellPrice.Silver,
			Copper: sellPrice.Copper,
		},
		Buy: &model.Price{
			Gold: buyPrice.Gold,
			Silver: buyPrice.Silver,
			Copper: buyPrice.Copper,
		},
		Description: &infoItem[0].Description ,
		Type: infoItem[0].Type,
		Rarity: infoItem[0].Rarity,
		Level: fmt.Sprint(infoItem[0].Level),
	}

	return resultItem, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
