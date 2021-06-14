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

	// PUBLISH A TASK -- Move to a controller triggered by http call
	task := Task{"generate_conversions", 1, 1}

	taskBytes, err := json.Marshal(task)
	if err != nil {fmt.Println(err)}

	err = taskQueue.PublishBytes(taskBytes)
	if err != nil {fmt.Println(err)}
}
