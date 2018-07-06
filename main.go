package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	df "github.com/leboncoin/dialogflow-go-webhook"
)

type params struct {
	Crypto    string  `json:"crypto"`
	Amount    float32 `json:"amount"`
	AccountTo string  `json:"account-to"`
}

func webhook(c *gin.Context) {
	var err error
	var dfr *df.Request
	var p params

	if err = c.BindJSON(&dfr); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	/*
		// Retrieve the params of the request
		if err = dfr.GetParams(&p); err != nil {
			fmt.Println("dfr.GetParams() error = %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Do something with params
		fmt.Println(p.Crypto)
		fmt.Println(p.Amount)
		fmt.Println(p.AccountTo)
	*/

	if err = dfr.GetParams(&p); err != nil {
		fmt.Println("dfr.GetParams() error = %v", err)
	}

	fmt.Println(p.Crypto)
	fmt.Println(p.Amount)
	fmt.Println(p.AccountTo)

	// Send back a fulfillment
	dff := &df.Fulfillment{
		FulfillmentMessages: df.Messages{
			df.ForGoogle(df.SingleSimpleResponse("hello", "hello")),
			{RichMessage: df.Text{Text: []string{fmt.Sprintf("Sending %.2f %v to %v", p.Amount, p.Crypto, p.AccountTo)}}},
		},
	}
	c.JSON(http.StatusOK, dff)
}

func main() {
	fmt.Println("v3")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK!")
	})

	r.POST("/webhook", webhook)
	if err := r.Run(":9090"); err != nil {
		panic(err)
	}
}
