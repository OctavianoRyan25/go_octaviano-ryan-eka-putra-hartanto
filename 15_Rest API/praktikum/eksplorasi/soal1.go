package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Fruit struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Family     string    `json:"family"`
	Order      string    `json:"order"`
	Genus      string    `json:"genus"`
	Nutritions Nutrition `json:"nutritions"`
}

type Nutrition struct {
	Calories      int     `json:"calories"`
	Fat           float32 `json:"fat"`
	Sugar         float32 `json:"sugar"`
	Carbohydrates float32 `json:"carbohydrates"`
	Protein       float32 `json:"protein"`
}

func main() {
	fruitName := flag.String("fruit", "", "Name of the fruit")
	flag.Parse()

	if *fruitName == "" {
		fmt.Println("Please provide a fruit name using the -fruit flag")
		return
	}

	res, err := http.Get("https://www.fruityvice.com/api/fruit/" + *fruitName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var fruit Fruit
	err = json.Unmarshal(resBody, &fruit)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("FRUIT DATA")
	fmt.Println("==========")
	fmt.Println("Fruit Name:" + fruit.Name)
	fmt.Println("Fruit Family:" + fruit.Family)
	fmt.Println("Calories:", fruit.Nutritions.Calories)
	fmt.Println("Fat:", fruit.Nutritions.Fat)
	fmt.Println("Sugar:", fruit.Nutritions.Sugar)
	fmt.Println("Carbohydrate:", fruit.Nutritions.Carbohydrates)
	fmt.Println("Protein:", fruit.Nutritions.Protein)
}
