package elasticsearch

import (
	"context"
	"fmt"
	"github.com/nhatnhanchiha/bookstore_utils-go/logger"
	"github.com/olivere/elastic"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	Index(index string, doc interface{}) (*elastic.IndexResponse, error)
	setClient(client *elastic.Client)
	Get(index string, id string) (*elastic.GetResult, error)
	Search(index string, query elastic.Query) (*elastic.SearchResult, error)
	Update(index string, id string, doc interface{}) (*elastic.UpdateResponse, error)
}

type esClient struct {
	Client *elastic.Client
}

func (e *esClient) setClient(client *elastic.Client) {
	e.Client = client
}

func (e *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := e.Client.Index().
		Index(index).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil
}

func Init() {
	//log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetBasicAuth("elastic", "AL7CIV3vFaF27qXstgqy"),
		//elastic.SetErrorLog(log),
		//elastic.SetInfoLog(log),
	)
	//todo: fix log
	if err != nil {
		panic(err.Error())
	}

	Client.setClient(client)

	// Create the index if it does not exists
}

func (e esClient) Get(index string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := e.Client.Get().Index(index).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get index by id = %s", id), err)
		return nil, err
	}

	if !result.Found {
		return nil, nil
	}

	return result, nil
}

func (e esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := e.Client.Search(index).Query(query).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search document index %s", index), err)
		return nil, err
	}

	return result, nil
}

func (e esClient) Update(index string, id string, doc interface{}) (*elastic.UpdateResponse, error) {
	ctx := context.Background()
	if updateResponse, err := e.Client.Update().Index(index).Id(id).Doc(doc).Do(ctx); err != nil {
		return nil, err
	} else {
		return updateResponse, nil
	}
}
