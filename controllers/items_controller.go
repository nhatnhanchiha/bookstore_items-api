package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nhatnhanchiha/bookstore_items-api/domain/items"
	"github.com/nhatnhanchiha/bookstore_items-api/domain/items/requests"
	"github.com/nhatnhanchiha/bookstore_items-api/domain/queries"
	"github.com/nhatnhanchiha/bookstore_items-api/logger"
	"github.com/nhatnhanchiha/bookstore_items-api/services"
	"github.com/nhatnhanchiha/bookstore_oauth-go/oauth"
	"github.com/nhatnhanchiha/bookstore_utils-go/rest_errors"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(*gin.Context)
	Get(*gin.Context)
	Search(*gin.Context)
	Patch(*gin.Context)
	Put(*gin.Context)
	Delete(c *gin.Context)
}

type itemsController struct {
}

func (i *itemsController) Create(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err.Message)
		return
	}

	if oauth.GetCallerId(c.Request) == 0 {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedError("not permission"))
		return
	}

	var item items.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		restErr := rest_errors.NewBadRequestError("Invalid json body for creating item")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	item.Seller = oauth.GetCallerId(c.Request)

	result, err := services.ItemService.Create(&item)
	if err != nil {
		c.JSON(err.Status(), err.Message())
	}

	c.JSON(http.StatusCreated, result)
}

func (i *itemsController) Get(c *gin.Context) {
	itemId := c.Param("id")
	logger.Info(itemId)
	result, err := services.ItemService.Get(itemId)
	if err != nil {
		c.JSON(err.Status(), err.Message())
	}
	c.JSON(http.StatusCreated, result)
}

func (i *itemsController) Search(c *gin.Context) {
	var query queries.EsQuery

	if err := c.Bind(&query); err != nil {
		logger.Info(err.Error())
	}

	foundItem, err := services.ItemService.Search(query)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, foundItem)
}

// Patch
//
// @Summary Patch a book to the book store
// @Description patch a book
// @Accept  json
// @Produce  json
// @Param   item   body    requests.PatchItemRequest     true        "Bla bla"
// @Success 200 {string} string	"ok"
// @Router /testapi/get-string-by-int/{some_id} [get]
func (i *itemsController) Patch(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var seller int64
	seller = oauth.GetCallerId(c.Request)
	if seller == 0 {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedError("not permission"))
		return
	}

	itemId := c.Param("id")
	var req requests.PatchItemRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError(err.Error()))
		return
	}

	req.Id = itemId
	req.Seller = seller

	if item, err := services.ItemService.Patch(req); err != nil {
		c.JSON(err.Status(), err)
	} else {
		c.JSON(http.StatusOK, item)
	}
}

func (i *itemsController) Put(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var seller int64
	seller = oauth.GetCallerId(c.Request)
	if seller == 0 {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedError("not permission"))
		return
	}

	itemId := c.Param("id")
	var req requests.UpdateItemRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError(err.Error()))
		return
	}

	req.Id = itemId
	req.Seller = seller

	if item, err := services.ItemService.UpdateFull(req); err != nil {
		c.JSON(err.Status(), err)
	} else {
		c.JSON(http.StatusOK, item)
	}
}

func (i *itemsController) Delete(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var seller int64
	seller = oauth.GetCallerId(c.Request)
	if seller == 0 {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedError("not permission"))
		return
	}

	itemId := c.Param("id")

	if err := services.ItemService.DeleteItem(itemId, seller); err != nil {
		c.JSON(err.Status(), err)
	} else {
		c.JSON(http.StatusNoContent, "")
	}
}
