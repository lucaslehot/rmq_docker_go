package main

import (
	"github.com/adjust/rmq/v3"
	"github.com/joho/godotenv"
	"log"
	"fmt"
	"net/http"
	"app/database"
	"app/router"
)

type Task struct {
	order string
	user_id int
}

const dwldPath = "./tmp"

func main() {
	port := "8080"
	newRouter := router.NewRouter()

	// Setting up Redis connection
	connection, err := rmq.OpenConnection("message_broker", "tcp", "localhost:6379", 1, errChan)
	taskQueue, err := connection.OpenQueue("tasks")

	// Loading environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.Connect()
	if err != nil {
		log.Fatalf("Impossible de se connecter à la bdd: %v", err)
	}

	log.Print("\nServer started on port " + port)

	newRouter.PathPrefix("/files/").Handler(http.StripPrefix("/files/",
	http.FileServer(http.Dir(dwldPath))))

	http.ListenAndServe(":"+port, newRouter)

	// Publish a task
	// /!\ Move to a controller triggered by http call /!\
	task := Task{"generate_conversions", 1}

	taskBytes, err := json.Marshal(task)
	if err != nil {fmt.Println(err)}

	err = taskQueue.PublishBytes(taskBytes)
	if err != nil {fmt.Println(err)}
}