package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/otz1/pr/entity"
	"github.com/otz1/pr/resource"
	"github.com/otz1/pr/util"
)

var (
	fetchResource = resource.NewFetchResource()
)

func Fetch(c *gin.Context) {
	var pageRankReq entity.PageRankRequest
	if err := c.MustBindWith(&pageRankReq, binding.JSON); err != nil {
		sentry.CaptureException(err)
		c.AbortWithError(http.StatusBadRequest, errors.New("bad request"))
		return
	}

	resp := fetchResource.Fetch(pageRankReq.Query)
	c.JSON(http.StatusOK, resp)
}

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         "https://5642cac4a6b14dec9815aafe9c87dfff@o372401.ingest.sentry.io/5197869",
		Environment: util.GetEnv("ENVIRONMENT", "local"),
	}); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	router := gin.Default()
	router.POST("/fetch", Fetch)
	router.Run(":8001")
}
