package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	//GET ALL
	resGet, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resGet.Body.Close()
	get1, err := io.ReadAll(resGet.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var resTodo []Todo
	err = json.Unmarshal(get1, &resTodo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response 1:")
	fmt.Println(resTodo)

	//GET BY ID
	resGet1, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resGet1.Body.Close()
	get2, err := io.ReadAll(resGet1.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var resTodoID Todo
	err = json.Unmarshal(get2, &resTodoID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response 1:")
	fmt.Println(resTodoID)

	//POST
	request := Todo{
		UserId:    1,
		Id:        2,
		Title:     "new Todo",
		Completed: true,
	}
	reqBody, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	resPost, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resPost.Body.Close()
	post, err := io.ReadAll(resPost.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var resTodoPost Todo
	err = json.Unmarshal(post, &resTodoPost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response post:")
	fmt.Println(resTodoPost)

	//PUT
	resPut, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/todos", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resPut.Body.Close()
	put, err := io.ReadAll(resPut.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var resTodoPut Todo
	err = json.Unmarshal(put, &resTodoPut)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response put:")
	fmt.Println(resTodoPut)

	//DELETE
	request5, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	client2 := &http.Client{}
	response5, err := client2.Do(request5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response5.Body.Close()
	fmt.Println("Response 5:")
	if response5.StatusCode == 200 {
		fmt.Println("Success delete")
	} else {
		fmt.Println("Failed delete")
	}
}
