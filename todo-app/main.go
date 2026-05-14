package main

import (
	"fmt"
	"net/http"
)

var shortGoLang = "Watch Go crash course"
var fullGoLang = "Watch Yash's Go GitHub repository"
var rewardDessert = "Reward myself with a doughnut"
var taskItems = []string {shortGoLang, fullGoLang, rewardDessert}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8000", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our Todo List App!"
	fmt.Fprintln(writer, greeting)
}