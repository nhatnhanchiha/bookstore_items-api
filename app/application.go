package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nhatnhanchiha/bookstore_items-api/clients/elasticsearch"
)

var (
	router = gin.New()
)

func StartApplication() {
	elasticsearch.Init()
	mapUrl()
	err := router.Run(":8888")
	if err != nil {
		panic(err.Error())
	}
}
