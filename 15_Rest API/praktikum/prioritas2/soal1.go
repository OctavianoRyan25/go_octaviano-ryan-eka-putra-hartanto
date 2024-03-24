package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	ID           int    `json:"client_id"`
	ClientKey    string `json:"client_key"`
	ClientSecret string `json:"client_secret"`
	Status       bool   `json:"status"`
}

func main() {
	//GET ALL
	resGet, err := http.Get("https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/client")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resGet.Body.Close()

	body, err := io.ReadAll(resGet.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("GET Response:")
	var resClients []Client
	err = json.Unmarshal(body, &resClients)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resGet.StatusCode)
	fmt.Println(resClients)

	//GET BY ID
	resGetById, err := http.Get("https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/client/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resGetById.Body.Close()

	getById, err := io.ReadAll(resGetById.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("GET Response:")
	var resClient Client
	err = json.Unmarshal(getById, &resClient)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resGetById.StatusCode)
	fmt.Println(resClient)

	//POST
	request := Client{
		ID:           5,
		ClientKey:    "CLIENT03",
		ClientSecret: "SECRET03",
		Status:       true,
	}
	reqBody, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	resPost, err := http.Post("https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/client", "application/json", bytes.NewBuffer(reqBody))
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
	var resClientPost Client
	err = json.Unmarshal(post, &resClientPost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response post:")
	fmt.Println(resPost.StatusCode)
	fmt.Println(resClientPost)

	//PUT
	resPut, err := http.NewRequest("PUT", "https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/client/1", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	client := &http.Client{}
	responsePut, err := client.Do(resPut)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer responsePut.Body.Close()
	put, err := io.ReadAll(responsePut.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var resClientPut Client
	err = json.Unmarshal(put, &resClientPut)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response put:")
	fmt.Println(responsePut.StatusCode)
	fmt.Println(resClientPut)

	//DELETE
	resDelete, err := http.NewRequest("PUT", "https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/client/1", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	client2 := &http.Client{}
	responseDelete, err := client2.Do(resDelete)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer responseDelete.Body.Close()
	delete, err := io.ReadAll(responseDelete.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var resClientDelete Client
	err = json.Unmarshal(delete, &resClientDelete)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response delete:")
	fmt.Println(responseDelete.StatusCode)
	fmt.Println(resClientDelete)
}
