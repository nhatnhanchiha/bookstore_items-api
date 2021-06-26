package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nhatnhanchiha/bookstore_items-api/clients/elasticsearch"
	"github.com/nhatnhanchiha/bookstore_items-api/domain/queries"
	"github.com/nhatnhanchiha/bookstore_utils-go/rest_errors"
	"strings"
)

const (
	indexItems = "items"
	strType    = "item"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, strType, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", rest_errors.NewError("database error"))
	}

	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, strType, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get index by id = %s", i.Id), rest_errors.NewError("database error"))
	}

	if !result.Found {
		return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
	}

	bytes, _ := result.Source.MarshalJSON()
	if err := json.Unmarshal(bytes, i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", rest_errors.NewError("database error"))
	}
	i.Id = itemId

	return nil
}

func (i Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, strType, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	items := make([]Item, result.TotalHits())

	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("Not found any item")
	}
	//fmt.Println(result)
	return items, nil
}

func (i *Item) Update() rest_errors.RestErr {
	if _, err := elasticsearch.Client.Update(indexItems, strType, i.Id, i); err != nil {
		return rest_errors.NewInternalServerError("error when update item", errors.New("database error"))
	} else {
		return nil
	}
}
