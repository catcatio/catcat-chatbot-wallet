package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	df "github.com/leboncoin/dialogflow-go-webhook"
)

type params struct {
	Bar string `json:"bar"`
}

func webhook(c *gin.Context) {
	var err error
	var dfr *df.Request

	if err = c.BindJSON(&dfr); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Send back a fulfillment
	dff := &df.Fulfillment{
		FulfillmentMessages: df.Messages{
			df.ForGoogle(df.SingleSimpleResponse("hello", "hello")),
			{RichMessage: df.Text{Text: []string{"hello foo"}}},
		},
	}
	c.JSON(http.StatusOK, dff)
}

func main() {
	fmt.Println("v2")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK!")
	})

	r.POST("/webhook", webhook)
	if err := r.Run(":9090"); err != nil {
		panic(err)
	}
}
