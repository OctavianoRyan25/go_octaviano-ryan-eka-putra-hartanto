package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     Name   `json:"name"`
}

type Name struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func fetchUsers(url string, wg *sync.WaitGroup, users chan User) {
	defer wg.Done()

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var userList []User
	err = json.Unmarshal(body, &userList)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, user := range userList {
		users <- user
	}
}

func main() {
	url := "https://fakestoreapi.com/users"
	var wg sync.WaitGroup
	users := make(chan User)

	wg.Add(1)
	go fetchUsers(url, &wg, users)

	go func() {
		wg.Wait()
		close(users)
	}()
	fmt.Println("Users:")
	for user := range users {
		fmt.Println("======")
		fmt.Printf("User ID: %d\n Username: %s\n Email: %s\n First Name: %s\n Last Name: %s\n", user.ID, user.Username, user.Email, user.Name.FirstName, user.Name.LastName)
		fmt.Println("======")
	}
}
