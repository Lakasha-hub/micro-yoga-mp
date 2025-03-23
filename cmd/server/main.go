package main

import (
	"fmt"
	"log"

	"github.com/Lakasha-hub/micro-yoga-mp/internal/microyoga"
	"github.com/Lakasha-hub/micro-yoga-mp/internal/microyoga/handlers"
	"github.com/Lakasha-hub/micro-yoga-mp/internal/microyoga/mp"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	cfg := LoadConfig()

	// _, err := db.NewDB(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	// if err != nil {
	// 	panic("Error connecting to the database")
	// }

	mpClient, err := mp.NewClient(cfg.MPAccessToken)
	if err != nil {
		panic("Error creating MercadoPago client")
	}

	membershipService := microyoga.NewMembershipService(mpClient)
	//paymentRepo := payment.NewPaymentRepo(db)

	router := gin.Default()

	router.POST("/payment", handlers.NewProccessPayment(*membershipService))
	router.POST("/webhook", handlers.HandleWebhookPayment())

	router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
	return nil
}
