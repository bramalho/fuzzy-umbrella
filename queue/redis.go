package queue

import (
	"log"
	"os"
	"strconv"

	"github.com/adeven/redismq"
	"github.com/joho/godotenv"
)

var queue *redismq.Queue
var consumer *redismq.Consumer

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PASS")
	redisDB, _ := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 64)
	redisQueue := os.Getenv("REDIS_QUEUE")

	queue := redismq.CreateQueue(redisHost, redisPort, redisPass, redisDB, redisQueue)

	consumer, err = queue.AddConsumer("my_consumer")
	if err != nil {
		panic(err)
	}
}

func GetQueue() *redismq.Queue {
	return queue
}

func GetConsumer() *redismq.Consumer {
	return consumer
}
