package main

import "fmt"

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTask []task

var tasks = allTask{
	{
		ID: 1,
		Name: "Task One",
		Content: "Some Content"
	}
}


func main() {

	fmt.Println("Hola mundo")
}
