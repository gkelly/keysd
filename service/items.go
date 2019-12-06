package service

import (
	"context"
	"sort"
	strings "strings"

	"github.com/keys-pub/keys"
	"github.com/keys-pub/keys/keyring"
	"github.com/pkg/errors"
)

// Item (RPC) returns an item for an ID.
func (s *service) Item(ctx context.Context, req *ItemRequest) (*ItemResponse, error) {
	id, err := keys.ParseID(req.ID)
	if err != nil {
		return nil, err
	}
	item, err := s.ks.Get(id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, keys.NewErrNotFound(id, "")
	}
	return &ItemResponse{
		Item: itemToRPC(item),
	}, nil
}

// Items (RPC) returns list of keyring items.
func (s *service) Items(ctx context.Context, req *ItemsRequest) (*ItemsResponse, error) {
	if req.Query != "" {
		return nil, errors.Errorf("query not implemented")
	}

	items, lierr := s.ks.List(&keyring.ListOpts{Type: req.Type})
	if lierr != nil {
		return nil, lierr
	}

	itemsOut := make([]*Item, 0, len(items))
	for _, item := range items {
		itemsOut = append(itemsOut, itemToRPC(item))
	}

	sort.Slice(itemsOut, func(i, j int) bool {
		if itemsOut[i].Type == itemsOut[j].Type {
			return strings.ToLower(itemsOut[i].ID) < strings.ToLower(itemsOut[j].ID)
		}
		return itemsOut[i].Type < itemsOut[j].Type
	})

	return &ItemsResponse{
		Items: itemsOut,
	}, nil
}

func itemToRPC(i *keyring.Item) *Item {
	item := &Item{
		ID:          i.ID,
		Type:        i.Type,
		Description: keys.TypeDescription(i.Type),
	}
	return item
}