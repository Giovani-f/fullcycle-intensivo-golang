package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/giovani-f/gointensivo-fullcycle/internal/infra/database"
	"github.com/giovani-f/gointensivo-fullcycle/internal/usecase"
	"github.com/giovani-f/gointensivo-fullcycle/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery, 100)
	// forever := make(chan bool)
	go rabbitmq.Consume(ch, out)

	qtdWorkers := 200
	for i := 1; i <= qtdWorkers; i++ {
		go worker(out, &uc, i)
	}
	// <-forever
}

func worker(deliverrMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliverrMessage {
		var inputDTO usecase.OrderInpuDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}
		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Printf("Worker: %d has processed order %s \n", workerId, outputDTO.Id)
		time.Sleep(1 * time.Second)
	}
}
