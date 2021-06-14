package main

import "github.com/adjust/rmq/v3"
import "fmt"

type Task struct {
	order string
	image_id int
	user_id int
}

func main() {
	connection, err := rmq.OpenConnection("message_broker", "tcp", "localhost:6379", 1, errChan)
	taskQueue, err := connection.OpenQueue("tasks")

	// CREATE CONSUMER FUNCTION
	err := taskQueue.StartConsuming(10, time.Second)
	name, err := taskQueue.AddConsumerFunc(func(delivery rmq.Delivery) {
		var task Task
		if err = json.Unmarshal([]byte(delivery.Payload()), &task); err != nil {
			// handle json error
			if err := delivery.Reject(); err != nil {
				// handle reject error
			}
			return
		}

		// perform task
		
		log.Printf("performing task %s", task)
		if err := delivery.Ack(); err != nil {
			// handle ack error
		}
	})}
