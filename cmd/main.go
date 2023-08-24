package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/punkestu/ecommerce-go/lib"
)

type Event struct {
	Event string `json:"event"`
	ID    string `json:"ID"`
	Param string `json:"param"`
}

func monitor(chanError chan sarama.ConsumerError, chanMessage chan *sarama.ConsumerMessage) {
	signals := make(chan os.Signal)
subLoop:
	for {
		select {
		case err := <-chanError:
			log.Println("on runtime ", err.Err.Error())
		case msg := <-chanMessage:
			var e lib.Event
			if err := json.Unmarshal(msg.Value, &e); err != nil {
				log.Println("bad event")
				continue
			}
			log.Println(e.Param)
			time.Sleep(time.Second)
		case sig := <-signals:
			if sig == os.Interrupt {
				break subLoop
			}
		}
	}
}

func main() {
	pubConn, err := lib.KafkaPubConn()
	if err != nil {
		log.Fatalln(err)
	}
	defer pubConn.Close()
	subConn, err := lib.KafkaSubConn()
	if err != nil {
		log.Fatalln(err)
	}
	defer subConn.Close()

	chanMessage := make(chan *sarama.ConsumerMessage, 256)
	chanError := make(chan sarama.ConsumerError)
	go lib.SubKafka("listener.main", chanMessage, chanError, sarama.OffsetNewest)
	go monitor(chanError, chanMessage)

	app := fiber.New()
	ctr := 0
	app.Post("/login", func(c *fiber.Ctx) error {
		ctr++
		if err := lib.PubKafka("listener.main", lib.Event{
			Event: "task.task",
			ID:    uuid.NewString(),
			Param: strconv.Itoa(ctr),
		}); err != nil {
			log.Fatalln(err)
		}
		return c.SendStatus(fiber.StatusOK)
	})
	app.Listen(":8080")
}
