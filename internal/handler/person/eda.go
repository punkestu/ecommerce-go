package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/punkestu/ecommerce-go/lib"
)

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
	listener()
}

func listener() {
	signals := make(chan os.Signal)
	chanMessage := make(chan *sarama.ConsumerMessage, 256)
	chanClaim := make(chan *sarama.ConsumerMessage, 256)
	chanError := make(chan sarama.ConsumerError)
	go lib.SubKafka("listener.main", chanMessage, chanError, sarama.OffsetNewest)
	go lib.SubKafka("listener.claim", chanClaim, chanError, sarama.OffsetNewest)
	id := uuid.NewString()
	taskId := ""
	task := ""
subLoop:
	for {
		select {
		case err := <-chanError:
			log.Println(err.Err.Error())
		case sig := <-signals:
			if sig == os.Interrupt {
				break subLoop
			}
		case msg := <-chanMessage:
			var event lib.Event
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Println("bad request")
				continue
			}
			if event.Event == "task.task" {
				taskId = event.ID
				task = event.Param
				lib.PubKafka("listener.claim", lib.Event{
					Event: "task.claim",
					ID:    event.ID,
					Param: id,
				})
				log.Println("try to claim " + task)
			}
		default:
			break
		}
		if taskId != "" {
			for {
				clm := <-chanClaim
				var event lib.Event
				if err := json.Unmarshal(clm.Value, &event); err != nil {
					log.Println("bad request")
					continue
				}
				if event.Event == "task.claim" {
					if event.ID == taskId {
						if event.Param == id {
							log.Println(task + " start")
							// process with 'task'
							time.Sleep(time.Second * 10)
							log.Println(task + " done")
						}
						taskId = ""
						task = ""
						break
					}
				}
			}
		}
	}
}
