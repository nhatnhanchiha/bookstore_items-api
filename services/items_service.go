package services

import (
	"github.com/nhatnhanchiha/bookstore_items-api/domain/items"
	"github.com/nhatnhanchiha/bookstore_items-api/domain/items/requests"
	"github.com/nhatnhanchiha/bookstore_items-api/domain/queries"
	"github.com/nhatnhanchiha/bookstore_utils-go/rest_errors"
)

var (
	ItemService itemsServiceInterface = &itemService{}
)

const (
	disableItemStatus = "disable"
)

type itemsServiceInterface interface {
	Create(*items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr)
	Patch(patchRequest requests.PatchItemRequest) (*items.Item, rest_errors.RestErr)
	UpdateFull(updateRequest requests.UpdateItemRequest) (*items.Item, rest_errors.RestErr)
	DeleteItem(id string, seller int64) rest_errors.RestErr
}

type itemService struct{}

func (s *itemService) Create(item *items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *itemService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}

func (s *itemService) Patch(patchRequest requests.PatchItemRequest) (*items.Item, rest_errors.RestErr) {
	current := items.Item{Id: patchRequest.Id}

	if err := current.Get(); err != nil {
		return nil, err
	}

	if current.Seller != patchRequest.Seller {
		return nil, rest_errors.NewUnauthorizedError("Not permission")
	}

	if patchRequest.Title != nil {
		current.Title = *patchRequest.Title
	}

	if patchRequest.Description != nil {
		if patchRequest.Description.PlainText != nil {
			current.Description.PlainText = patchRequest.Description.PlainText
		}
		if patchRequest.Description.Html != nil {
			current.Description.Html = patchRequest.Description.Html
		}
	}

	if patchRequest.Pictures != nil {
		current.Pictures = patchRequest.Pictures
	}

	if patchRequest.Video != nil {
		current.Video = *patchRequest.Video
	}

	if patchRequest.Price != nil {
		current.Price = *patchRequest.Price
	}

	if patchRequest.AvailableQuantity != nil {
		current.AvailableQuantity = *patchRequest.AvailableQuantity
	}

	if patchRequest.SoldQuantity != nil {
		current.SoldQuantity = *patchRequest.SoldQuantity
	}

	if patchRequest.Status != nil {
		current.Status = *patchRequest.Status
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return &current, nil
}

func (s *itemService) UpdateFull(updateRequest requests.UpdateItemRequest) (*items.Item, rest_errors.RestErr) {
	current := items.Item{Id: updateRequest.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if current.Seller != updateRequest.Seller {
		return nil, rest_errors.NewUnauthorizedError("Not permission")
	}

	current = items.Item(updateRequest)

	if err := current.Update(); err != nil {
		return nil, err
	}

	return &current, nil
}

func (s itemService) DeleteItem(id string, seller int64) rest_errors.RestErr {
	status := disableItemStatus
	item := requests.PatchItemRequest{Id: id, Seller: seller, Status: &status}
	if _, err := s.Patch(item); err != nil {
		return err
	} else {
		return nil
	}
}
