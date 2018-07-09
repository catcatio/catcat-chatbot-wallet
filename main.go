package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	df "github.com/leboncoin/dialogflow-go-webhook"

	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"

	"github.com/catcatio/wallet"
)

type params struct {
	Crypto    string  `json:"crypto"`
	Amount    float32 `json:"amount"`
	AccountTo string  `json:"account-to"`
}

func foo() {
	// address: GB6S3XHQVL6ZBAF6FIK62OCK3XTUI4L5Z5YUVYNBZUXZ4AZMVBQZNSAU
	from := "SCRUYGFG76UPX3EIUWGPIQPQDPD24XPR3RII5BD53DYPKZJGG43FL5HI"

	// seed: SDLJZXOSOMKPWAK4OCWNNVOYUEYEESPGCWK53PT7QMG4J4KGDAUIL5LG
	to := "GA3A7AD7ZR4PIYW6A52SP6IK7UISESICPMMZVJGNUTVIZ5OUYOPBTK6X"

	tx, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: from},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.Payment(
			b.Destination{AddressOrSeed: to},
			b.NativeAmount{Amount: "0.1"},
		),
	)
	if err != nil {
		panic(err)
	}

	txe, err := tx.Sign(from)
	if err != nil {
		panic(err)
	}

	txeB64, err := txe.Base64()

	if err != nil {
		panic(err)
	}

	fmt.Printf("tx base64: %s", txeB64)
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
	fmt.Println("🚀 v3")

	foo()
	wallet.Bar()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK!")
	})

	r.POST("/webhook", webhook)
	if err := r.Run(":9090"); err != nil {
		panic(err)
	}
}
